use clap::{Arg, ArgAction, Command};

pub fn cli() -> Command {
    Command::new("dumber")
        .about("A tool to (un)mumber sections and add/move toc(s) of a Markdown document")
        .override_usage("Usage: dumber [OPTIONS] <FILE>")
        .arg_required_else_help(true)
        .allow_external_subcommands(true)
        .arg(
            Arg::new("write")
                .action(ArgAction::SetTrue)
                .short('w')
                .long("write")
                .help("Write changes to the .md file (default to stdout)"),
        )
        .arg(
            Arg::new("all")
                .action(ArgAction::SetTrue)
                .short('a')
                .long("all-headers")
                .help("Numbering all section headers, starting from the main document title, first H1"),
        )
        .arg(
            Arg::new("remove")
                .action(ArgAction::SetTrue)
                .short('r')
                .long("remove")
                .help("Remove changes from the .md file (default to stdout)"),
        )
        .arg(
            Arg::new("FILE")
                .required(true)
                .help("The Markdown file to process"),
        )
}
