use regex::Regex;

pub fn add_section_chunk(s: &mut String, hc: &i8, cht: &usize, ht: &usize) {
    if *hc > 0 && *cht >= *ht {
        s.push_str(&(hc.to_string() + "."));
    }
}

pub fn to_toc_entry(u: usize, r: &Regex, l: &str) -> String {
    let m = r.captures(l).unwrap();
    let c = m[1].len() - u;
    let mut entry = String::with_capacity(c * 4 + m[2].len() + m[3].len() * 2 + 6);

    for _ in 0..c {
        entry.push_str("    ");
    }
    entry.push_str("- [");
    entry.push_str(&m[2]);
    entry.push_str("](#");
    entry.push_str(&m[2].replace(".", ""));
    entry.push('-');
    entry.push_str(&m[3].replace(" ", "-").to_lowercase());
    entry.push_str(") ");
    entry.push_str(&m[3]);

    entry
}
