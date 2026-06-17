use std::process::Command;

fn main() {
    let githash = Command::new("git")
        .args(["rev-parse", "--short", "HEAD"])
        .output()
        .unwrap();
    println!(
        "{}",
        "cargo:rustc-env=GIT_COMMIT_SHORT_HASH=".to_owned()
            + String::from_utf8_lossy(&githash.stdout).trim()
    );
}
