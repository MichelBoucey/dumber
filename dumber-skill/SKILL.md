---
name: Dumber
description: Number Markdown section headings and add, update, or remove tables of contents using the `dumber` CLI. Use when asked to number or unnumber headings, generate a table of contents, or strip a generated ToC from a Markdown file.
license: BSD-3-Clause
---

# Dumber

`dumber` is a command line tool that **d**(igital n)**umber**s the ATX (`#`)
headings of a Markdown document and generates or removes table(s) of contents
(ToC) accordingly.

Reach for it whenever a task is about:

- adding or refreshing numbers on Markdown headings,
- inserting or updating a table of contents,
- stripping numbering and/or a generated ToC from a file.

Prefer `dumber` over editing numbers or ToC entries by hand: it keeps the
hierarchy consistent and makes the change reversible.

## Availability check

Confirm the tool is installed before using it:

```
dumber --version
dumber --help
```

If it is missing, install with one of:

```
cargo install dumber          # from crates.io
yay -S dumber                 # Arch Linux (AUR)
make install                  # from a source checkout
```

Pre-built Linux x86_64 binaries are also published on the project's GitHub
releases page.

## Command reference

```
dumber [OPTIONS] [FILE]
```

`FILE` is the Markdown file to process, or `-` for standard input. When
`FILE` is omitted, standard input is read.

| Option             | Short | Effect                                                                 |
| ------------------ | ----- | ---------------------------------------------------------------------- |
| `--write`          | `-w`  | Write changes back to the `.md` file. Without it, output goes to stdout (a preview). |
| `--remove`         | `-r`  | Remove changes: strip heading numbers and delete generated ToC lines.  |
| `--all-headers`    | `-a`  | Number **all** section headers, including the main document title (first H1). |
| `--version`        | `-v`  | Print the version and exit.                                            |
| `--help`           | `-h`  | Print help and exit.                                                   |

## Core rules

- **Preview is the default.** With no `-w`, `dumber` prints the rewritten
  document to stdout and leaves the file untouched. Use this to inspect the
  result before committing to it.
- **`-w` requires a file.** `--write` cannot be combined with standard input;
  it exits with an error.
- **ATX headings only.** Only lines starting with `#` at column 0 are treated
  as headings (`#` through `######`, with a space after the hashes). Setext
  headings (`===` / `---` underlines) and indented hashes are ignored.
- **The first H1 is the title.** It is left unnumbered by default, and the
  numbering of the remaining sections starts at `1.`. Use `-a` to number the
  first H1 as well. Any later H1 is numbered normally.
- **Hierarchical numbering.** Sections are numbered by level, e.g. `1.`,
  `1.2.`, `1.2.3.`, and deeper counters reset when a shallower heading
  appears.
- **Removal is a clean round-trip.** `-r` strips existing heading numbers and
  drops the ToC lines that `dumber` generated, restoring the document's
  structure.

## Tables of contents

A ToC is generated at each occurrence of the HTML comment marker:

```
<!-- ToC -->
```

Notes:

- The marker must match exactly: `<!--`, whitespace, the word `ToC`,
  whitespace, `-->`. Keep the spaces (`<!--ToC-->` without spaces will not
  match).
- Place the marker on its own line wherever a ToC should appear; a document
  may contain several markers.
- Entries are written immediately **after** the marker line, indented by
  heading depth, and links are anchored to the heading text.
- On every run, previously generated ToC entries are recognised and
  regenerated (or removed with `-r`), so running the tool repeatedly is safe.

## Recommended agent workflow

1. **Inspect** the target Markdown file and note its heading structure and any
   `<!-- ToC -->` markers.
2. **Dry run** without `-w` and read the stdout preview.
3. **Review** that numbering and ToC entries are correct.
4. **Apply** with `-w` once the preview looks right.
5. **Verify** the written file, or revert with `-r -w` if the result is not
   wanted.

## Examples

Given `doc.md`:

```markdown
# Project

<!-- ToC -->

# Getting started

## Install

## Configure

# Reference

## Flags
```

```sh
# Preview: number headings and fill the ToC, to stdout only
dumber doc.md

# Number every header, including the first H1 title
dumber -a doc.md

# Apply the numbering and ToC in place
dumber -w doc.md

# Read from stdin (preview only; -w is not allowed with stdin)
cat doc.md | dumber -
dumber - < doc.md

# Preview removing numbers and the generated ToC
dumber -r doc.md

# Strip numbers and the generated ToC in place
dumber -r -w doc.md
```

## Limits

- ToC generation is not implemented or tested for UTF-8; non-ASCII headings may
  produce imperfect anchors.
- Only `#`-style (ATX) headings are considered; table-of-contents links follow
  the heading text, so headings that differ only in punctuation can collide.
