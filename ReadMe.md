# Dumber [![Rust CI](https://github.com/MichelBoucey/dumber/actions/workflows/rust.yml/badge.svg?branch=main)](https://github.com/MichelBoucey/dumber/actions/workflows/rust.yml)

`dumber`, a (not so dumb) command line tool for **d**(igital n)**umber**(ing) Markdown document sections and creation of table(s) of contents accordingly.

With `dumber` (you already use `git`, don't you?) you can add to, update, or remove from your Mardown files:

- section numbers to header sections. This works on hash sign only. &lt;H1&gt; (#) stays unnumbered as the main title. Use `-a` for numbering *all* section tags.
- table(s) of contents with links on entries.

See an [example](https://github.com/MichelBoucey/dumber/blob/main/example.md).

_N.B._ : The table of contents generation is not tested nor implemented for UTF8.

## 1. Installation

### 1.1. From sources

```
make install
```

### 1.2. From crates.io

```
cargo install dumber
```

### 1.3. From ArchLinux AUR

Build the ArchLinux package for `dumber` from [AUR](https://aur.archlinux.org/packages/dumber).

Or with `yay`, run `yay -S dumber`.

## 2. Optional test suite

Rebuild and install a brand new `dumber` and run a small test suite.

```
make test
```

## 3. Usage

### 3.1. Command line options
```
user@box $ dumber --help
A tool to (un)number sections and add/remove toc(s) of a Markdown document

Usage: dumber [OPTIONS] [FILE]

Arguments:
  [FILE]  The Markdown file to process

Options:
  -w, --write        Write changes to the .md file (default to stdout)
  -r, --remove       Remove changes from a modified .md file (default to stdout)
  -a, --all-headers  Numbering all section headers, starting from the main document title, first H1
  -v, --version      Print version
  -h, --help         Print help
```

### 3.2. Add table(s) of contents

To add a table of contents you have to add a line with the HTML comment **&lt;!-- Toc --!&gt;**, where you want a table of contents to appear:

```
<!-- ToC -->
```

The table of contents will be written just after the HTML comment line, and you can add this comment line as many times as you want, if you are, like me, a big fan of tables of contents, or if the length of your document needs a second table of contents at its end.

