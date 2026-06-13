use regex::Regex;
use std::fs::File;
use std::io::{self, prelude::*, BufReader};

fn main() -> io::Result<()> {
    let mut _header_counts: [i8; 7];
    let _header_lines: Vec<String> = Vec::new();
    let header_line = Regex::new(r"^(#{1,6})\s+([\d\.]*)\s*(.*)$").unwrap();
    let mut _rewritten_line: String;
    let mut _section: &mut String;
    // let _toc_insertion_line = Regex::new(r"").unwrap();
    // let _toc_line = Regex::new(r"").unwrap();

    let file = File::open("test/test.md")?;
    let reader = BufReader::new(file);

    for result in reader.lines() {
        let line = result.unwrap();
        let caps = header_line.captures(&line);

        if caps.is_some() {
            let cs = caps.unwrap();
            if header_line.is_match(&line) {
                if cs.len() == 4 {
                    let _header = &cs[1];
                    let _current_header_type = &cs[1].len();
                    println!(
                        "{}",
                        cs[1].to_owned() + " " + "1.2." + " " + &cs[3].to_owned()
                    );
                }
            }
        } else {
            println!("{}", line);
        }
    }

    Ok(())
}
