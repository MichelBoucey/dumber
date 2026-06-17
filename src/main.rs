use regex::Regex;
use std::fs::{exists, File};
use std::io::{self, prelude::*, BufReader, BufWriter, Write};
mod args;
mod internal;
use crate::args::cli;
use crate::internal::to_toc_entry;
use std::process::{exit, Command};

fn main() -> io::Result<()> {
    let version: String = "v4.0.0".to_string();

    let mut header_counters: [i8; 7] = [0; 7];
    let mut header_lines: Vec<String> = Vec::new();
    let mut section: String = String::from("");
    let mut rewritten_line: String = String::from("");
    let mut rewritten_lines: Vec<String> = Vec::new();
    let mut first_h1_done: bool = false;
    let mut is_toc_insertion_line: bool = false;
    let mut upper_header_level: usize = 0;

    let toc_insertion_line = Regex::new(r"^<!--\s+\bToC\b\s+-->\s*$").unwrap();
    let toc_line = Regex::new(r"^\s*-\s\[[\d\.]*\]\(#\d*").unwrap();
    let header_line = Regex::new(r"^(#{1,6})\s+([\d\.]*)\s*(.*)$").unwrap();

    let matches = cli().get_matches();
    let flag_write = matches.get_one::<bool>("write").expect("required");
    let flag_remove = matches.get_one::<bool>("remove").expect("required");
    let no_title_skip = matches.get_one::<bool>("all").expect("required");
    let flag_version = matches.get_one::<bool>("version").expect("required");

    if *flag_version {
        let githash = Command::new("git")
            .args(["rev-parse", "--short", "HEAD"])
            .output()
            .unwrap();
        println!(
            "{}",
            "dumber ".to_owned()
                + &version
                + " ("
                + String::from_utf8_lossy(&githash.stdout).trim()
                + ") released under 3-Clause BSD License"
        );
        println!("Copyright © 2021-2026 Michel Boucey (michel.boucey@gmail.com)");
        exit(0)
    }

    if let Some(md_filepath) = matches.get_one::<String>("FILE") {
        if !exists(md_filepath).unwrap() {
            println!("{}", md_filepath.to_owned() + " is not a file");
            exit(1)
        }

        let file = File::open(md_filepath)?;

        let reader = BufReader::new(file);

        for result in reader.lines() {
            let line = result.unwrap();
            let caps = header_line.captures(&line);

            if let Some(cs) = caps {
                if header_line.is_match(&line) && cs.len() == 4 {
                    let header = &cs[1];
                    let title = &cs[3];
                    let current_header_type = &cs[1].len();

                    if first_h1_done || *no_title_skip {
                        header_counters[*current_header_type] += 1
                    }

                    if !first_h1_done && *current_header_type == 1 {
                        first_h1_done = true;
                    }

                    if *flag_remove {
                        rewritten_line = header.to_owned() + " " + title
                    } else {
                        for (header_type, _) in header_counters.iter().enumerate().skip(1) {
                            internal::add_section_chunk(
                                &mut section,
                                &header_counters[header_type],
                                current_header_type,
                                &header_type,
                            );
                        }

                        if !section.is_empty() {
                            section += " "
                        }

                        rewritten_line = header.to_owned() + " " + &*section + title;

                        header_lines.push(rewritten_line.clone());

                        for v in header_counters.iter_mut().skip(current_header_type + 1) {
                            *v = 0;
                        }

                        section = "".to_string();
                    }
                }

                rewritten_lines.push(rewritten_line.clone());
            } else if !toc_line.is_match(&line) {
                if toc_insertion_line.is_match(&line) {
                    is_toc_insertion_line = true
                }
                rewritten_lines.push(line);
            }
        }

        if !flag_remove && is_toc_insertion_line {
            let first_header_line = header_line.captures(&header_lines[0]).unwrap();
            upper_header_level = first_header_line[1].len()
        }

        if *flag_write {
            let file = File::create(md_filepath)?;
            let mut writer = BufWriter::new(file);
            for rewritten_line in rewritten_lines {
                writeln!(writer, "{}", rewritten_line)?;
                if !flag_remove
                    && is_toc_insertion_line
                    && toc_insertion_line.is_match(&rewritten_line)
                {
                    for hline in header_lines.clone().into_iter().skip(1) {
                        writeln!(
                            writer,
                            "{}",
                            to_toc_entry(
                                upper_header_level,
                                header_line.clone(),
                                hline.to_string()
                            )
                        )?
                    }
                }
            }
        } else {
            for rewritten_line in rewritten_lines {
                println!("{}", rewritten_line);
                if !flag_remove
                    && is_toc_insertion_line
                    && toc_insertion_line.is_match(&rewritten_line)
                {
                    for hline in header_lines.clone().into_iter().skip(1) {
                        println!(
                            "{}",
                            to_toc_entry(
                                upper_header_level,
                                header_line.clone(),
                                hline.to_string()
                            )
                        )
                    }
                }
            }
        }
    } else {
        println!("Usage: dumber [OPTIONS] [FILE]");
        println!("Specify the --help flag to see full usage")
    }

    Ok(())
}
