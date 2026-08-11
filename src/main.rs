use regex::Regex;
use std::env;
use std::fs::{File, exists};
use std::io::{self, BufReader, BufWriter, Write, prelude::*};
mod args;
mod internal;
use crate::args::cli;
use crate::internal::to_toc_entry;
use std::process::exit;

fn main() -> io::Result<()> {
    let mut header_counters: [i8; 7] = [0; 7];
    let mut header_lines: Vec<String> = Vec::new();
    let mut section: String = String::new();
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
        println!(
            "{}",
            "dumber ".to_owned()
                + env!("CARGO_PKG_VERSION")
                + " ("
                + env!("GIT_COMMIT_SHORT_HASH")
                + ") released under 3-Clause BSD License"
        );
        println!("Copyright © 2021-2026 Michel Boucey (michel.boucey@gmail.com)");
        exit(0)
    }

    let use_stdin = matches.get_one::<String>("FILE").is_none_or(|f| f == "-");

    if *flag_write && use_stdin {
        eprintln!("Error: --write cannot be used with stdin");
        exit(1);
    }

    let reader: Box<dyn BufRead> = if use_stdin {
        Box::new(BufReader::new(io::stdin()))
    } else {
        let md_filepath = matches.get_one::<String>("FILE").unwrap();
        if !exists(md_filepath).unwrap() {
            println!("{}", md_filepath.to_owned() + " is not a file");
            exit(1)
        }
        Box::new(BufReader::new(File::open(md_filepath)?))
    };

    for result in reader.lines() {
        let line = result?;
        match line.as_bytes().first() {
            Some(b'#') => {
                if let Some(cs) = header_line.captures(&line) {
                    let header = &cs[1];
                    let title = &cs[3];
                    let current_header_type = cs[1].len();

                    if first_h1_done || *no_title_skip {
                        header_counters[current_header_type] += 1
                    }

                    if !first_h1_done && current_header_type == 1 {
                        first_h1_done = true;
                    }

                    let mut rewritten_line =
                        String::with_capacity(header.len() + title.len() + section.len() + 2);

                    if *flag_remove {
                        rewritten_line.push_str(header);
                        rewritten_line.push(' ');
                        rewritten_line.push_str(title);
                    } else {
                        for (header_type, _) in header_counters.iter().enumerate().skip(1) {
                            internal::add_section_chunk(
                                &mut section,
                                &header_counters[header_type],
                                &current_header_type,
                                &header_type,
                            );
                        }

                        if !section.is_empty() {
                            section += " "
                        }

                        rewritten_line.push_str(header);
                        rewritten_line.push(' ');
                        rewritten_line.push_str(&section);
                        rewritten_line.push_str(title);

                        header_lines.push(rewritten_line.clone());

                        for v in header_counters.iter_mut().skip(current_header_type + 1) {
                            *v = 0;
                        }

                        section.clear();
                    }

                    rewritten_lines.push(rewritten_line);
                }
            }
            _ => {
                if line.starts_with("<!--") || line.trim_start().starts_with('-') {
                    if !toc_line.is_match(&line) {
                        if toc_insertion_line.is_match(&line) {
                            is_toc_insertion_line = true
                        }
                        rewritten_lines.push(line);
                    }
                } else {
                    rewritten_lines.push(line);
                }
            }
        }
    }

    let toc_entries: Vec<String> = if !flag_remove && is_toc_insertion_line {
        if let Some(first_header_line) = header_lines.first() {
            upper_header_level = header_line
                .captures(first_header_line)
                .map_or(0, |m| m[1].len())
        }
        header_lines
            .iter()
            .skip(1)
            .map(|hline| to_toc_entry(upper_header_level, &header_line, hline))
            .collect()
    } else {
        Vec::new()
    };

    if *flag_write {
        let md_filepath = matches.get_one::<String>("FILE").unwrap();
        let file = File::create(md_filepath)?;
        let mut writer = BufWriter::new(file);
        for rewritten_line in rewritten_lines {
            writeln!(writer, "{}", rewritten_line)?;
            if !flag_remove && is_toc_insertion_line && toc_insertion_line.is_match(&rewritten_line)
            {
                for entry in &toc_entries {
                    writeln!(writer, "{}", entry)?
                }
            }
        }
    } else {
        let mut writer = BufWriter::new(io::stdout().lock());
        for rewritten_line in rewritten_lines {
            writeln!(writer, "{}", rewritten_line)?;
            if !flag_remove && is_toc_insertion_line && toc_insertion_line.is_match(&rewritten_line)
            {
                for entry in &toc_entries {
                    writeln!(writer, "{}", entry)?
                }
            }
        }
    }

    Ok(())
}
