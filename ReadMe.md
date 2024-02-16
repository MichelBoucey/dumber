<h1>Dumber</h1>

`dumber`, a (not so dumb) command line tool for **d**(igital n)**umber**(ing) Markdown document sections and creation of table(s) of contents accordingly.

With `dumber` (you already use `git`, don't you?) you can *add* or *remove* to your Mardown files:

- section numbers to header sections. This works on hash sign only, so one can exclude HTML section tags, as &lt;H1&gt;, to stay unnumbered.
- a table of contents with links on entries.

See an [example](./example.md).

_N.B._ : The table of content generation is not tested and implemented for UTF8 yet.

## 1. Installation

```
make install
```

## 2. Usage

### 2.1. Command line options
```
user@machine $ dumber -h
Usage: dumber [OPTION] FILE

  -h    Show help
  -r    Remove section numbers and table of contents from the .md file
  -v    Show version
  -w    Write section numbers to the .md file (default to stdout)
```

### 2.2. Add table(s) of contents

To add a table of contents you have to add a line with the HTML comment **&lt;!-- Toc --!&gt;**, where you want a table of contents to appear:

```
<!-- ToC -->
- [1.](#1-installation) Installation
- [2.](#2-usage) Usage
    - [2.1.](#21-options) Options
    - [2.2.](#22-add-table(s)-of-content) Add table(s) of content
```

The table of contents will be written just after the HTML comment line. And you can add this comment line as many times as you want, if you are, like me, a big fan of tables of contents, or if the length of your document needs a second table of contents at its end.

