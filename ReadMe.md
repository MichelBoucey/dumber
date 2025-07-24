# Dumber

`dumber`, a (not so dumb) command line tool for **d**(igital n)**umber**(ing) Markdown document sections and creation of table(s) of contents accordingly.

With `dumber` (you already use `git`, don't you?) you can *add* or *remove* to your Mardown files:

- section numbers to header sections. This works on hash sign only. &lt;H1&gt; (#) stays unnumbered as the main title. Use `-a` for numbering *all* section tags.
- a table of contents with links on entries.

See an [example](./example.md).

_N.B._ : The table of contents generation is not tested nor implemented for UTF8.

## 1. Installation

### 1.1. From sources

```
make install
```

### 1.2. From ArchLinux AUR

Build the ArchLinux package for `dumber` from [AUR](https://aur.archlinux.org/packages/dumber).

## 2. Optional test suite

Rebuild and install a brand new `dumber` and run a small test suite.

```
make test
```

## 3. Usage

### 3.1. Command line options
```
user@machine $ dumber -h
Usage: dumber [OPTION] FILE

  -a	Numbering all sections from the main document title (H1)
  -h	Show help
  -r	Remove table of contents and section numbers from the .md file
  -v	Show version
  -w	Write section numbers to the .md file (default to stdout)
```

### 3.2. Add table(s) of contents

To add a table of contents you have to add a line with the HTML comment **&lt;!-- Toc --!&gt;**, where you want a table of contents to appear:

```
<!-- ToC -->
```

The table of contents will be written just after the HTML comment line, and you can add this comment line as many times as you want, if you are, like me, a big fan of tables of contents, or if the length of your document needs a second table of contents at its end.

