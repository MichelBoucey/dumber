use clap::{Arg, Command};

pub fn cli() -> Command {
    Command::new("dumber")
        .about("Numbering Markdown sections")
        .override_usage("Usage: dumber [OPTIONS] file")
        .arg_required_else_help(true)
        .allow_external_subcommands(true)
        .arg(
            Arg::new("file")
                .required(true)
                .help("The Markdown file to process")
        )
}
