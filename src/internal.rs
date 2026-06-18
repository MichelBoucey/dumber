use regex::Regex;

pub fn add_section_chunk(s: &mut String, hc: &i8, cht: &usize, ht: &usize) {
    if *hc > 0 && *cht >= *ht {
        s.push_str(&(hc.to_string() + "."));
    }
}

pub fn to_toc_entry(u: usize, r: Regex, l: String) -> String {
    let m = r.captures(&l).unwrap();
    let c = m[1].len() - u;

    "    ".to_string().repeat(c)
        + "- ["
        + &m[2]
        + "](#"
        + &m[2].to_string().replace(".", "")
        + "-"
        + &m[3].replace(" ", "-").to_lowercase()
        + ") "
        + &m[3]
}
