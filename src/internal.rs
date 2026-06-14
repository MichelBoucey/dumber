pub fn add_section_chunk(s: &mut String, hc: &i8, cht: &usize, ht: &usize) {
    if *hc > 0 && *cht >= *ht {
        s.push_str(&(hc.to_string() + "."));
    }
}
