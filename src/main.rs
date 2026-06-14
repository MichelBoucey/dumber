use regex::Regex;
use std::fs::File;
use std::io::{self, prelude::*, BufReader};
use crate::clap::cli;
mod internal;
mod clap;

fn main() -> io::Result<()> {
    let mut header_counters: [i8; 7] = [0; 7];
    let mut section: String = String::from("");
    let header_line = Regex::new(r"^(#{1,6})\s+([\d\.]*)\s*(.*)$").unwrap();
    let no_title_skip: bool = false;
    let mut first_h1_done: bool = false;
    // let toc_insertion_line = Regex::new(r"^<!--\s+\bToC\b\s+-->\s*$").unwrap();
    // let toc_line = Regex::new(r"^\s*-\s\[[\d\.]*\]\(#\d*").unwrap();
    // let header_lines: Vec<String> = Vec::new();
    // let mut rewritten_line: String;

    let matches = cli().get_matches();
    let file = matches.get_one::<String>("file").expect("required");
    let file = File::open(file)?;
    let reader = BufReader::new(file);

    for result in reader.lines() {
        let line = result.unwrap();
        let caps = header_line.captures(&line);

        if let Some(cs) = caps {
            if header_line.is_match(&line) {
                if cs.len() == 4 {
                    let header = &cs[1];
                    let title = &cs[3];
                    let current_header_type = &cs[1].len();

                    if first_h1_done || no_title_skip {
                        header_counters[*current_header_type] += 1
                    }

                    if !first_h1_done && *current_header_type == 1 {
                        first_h1_done = true;
                        // println!("{}", "Hey!");
                    }

                    for (header_type,_) in header_counters.iter().enumerate().skip(1) {
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

                    println!(
                        "{}",
                        header.to_owned() + " " + &*section + title
                    );

                    for v in header_counters.iter_mut().skip(current_header_type + 1) {
                        *v = 0;
                    }

                    section = "".to_string();
                }
            }
        } else {
            println!("{}", line);
        }
    }

    Ok(())
}
