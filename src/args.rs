use clap::{Arg, ArgAction, Command};

pub fn cli() -> Command {
    Command::new("dumber")
        .author("Michel Boucey, michel.boucey@gmail.com")
        .about("A tool to (un)number sections and add/remove toc(s) of a Markdown document")
        .arg_required_else_help(false)
        .arg(
            Arg::new("write")
                .action(ArgAction::SetTrue)
                .short('w')
                .long("write")
                .help("Write changes to the .md file (default to stdout)"),
        )
        .arg(
            Arg::new("remove")
                .action(ArgAction::SetTrue)
                .short('r')
                .long("remove")
                .help("Remove changes from a modified .md file (default to stdout)"),
        )
        .arg(
            Arg::new("all")
                .action(ArgAction::SetTrue)
                .short('a')
                .long("all-headers")
                .help("Numbering all section headers, starting from the main document title, first H1"),
        )
        .arg(
            Arg::new("version")
                .action(ArgAction::SetTrue)
                .short('v')
                .long("version")
                .help("Print version"),
        )
        .arg(
            Arg::new("FILE")
                .help("The Markdown file to process"),
        )
}
