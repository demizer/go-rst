================================
go-rst - reStructuredText for Go
================================
:Modified: Fri Jun 13 15:54 2014

A reStructuredText parser for and implemented in the Go programming language.

**This library is far from complete and barely usable.**


=================
How to contribute
=================

* **Convert tests into JSON**

  The docutils tests are implemented in a "psuedo xml" which is non-standard.
  Translating the tests into JSON has the benefit of making the reStructuredText
  tests programming language neutral so that reStructuredText parsers can be
  implemented in other programming languages. See
  https://github.com/demizer/go-rst/tree/master/testdata
  for more information.

* **Implement an element**

  Implement an element from the list above.

* **Write some documentation**

  All projects need good documentation!

* **Test and report**

  Not actually possible in the current state, but using the library and writing
  bug reports is always helpful.

============
Contributors
============

These people have donated their valuable time in contributing to this library
and are here graciously recognized for their contributions!

* **Jesus Alvarez**

------
Status
------

Work is currently underway on translating the docutils tests from psuedo xml
into JSON.

====  ============================== ======
Done  Element                        Detail
====  ============================== ======
0%    Exported API                   Not even designed yet. Look to other text file parsers for inspiration. Suggestions welcome!
0%    HTML transpiler                Tree.Nodes can be easily walked to produce output html.
0%    Demo homepage                  To be implemented in Go and hosted at restructuredtext.com
5%    Documentation                  Most functions documented. Nothing on API.
100%  Section Headers                All tests from docutils converted to JSON and implemented.
5%    Transitions                    Basic lexing.
25%   Paragraphs                     Basic lexing, no inline markup, no tests.
0%    Bullet lists
5%    Enumerated lists               Some lexing implemented.
0%    Definition lists
0%    Field lists
0%    Biblio. fields
0%    RCS keywords
0%    Option lists
5%    Literal blocks                 Some parsing.
0%    Indented literal blocks
0%    Quoted literal blocks
0%    Line blocks
10%   Block quotes                   Some parsing.
0%    Doctest blocks                 Will use Go instead of Python. Much lower priority than everything else.
0%    Grid tables                    Will be gruesome to implement.
0%    Simple tables
0%    Footnotes
0%    Auto-Numbered Footnotes
0%    Auto-Symbol Footnotes
0%    Mixed Auto Footnotes
0%    Citations
0%    Hyperlink targets
0%    Anonymous Hyperlinks
0%    Substitution Definitions
2%    Comments                       Basic lexing and parsing.
0%    Implicit Hyperlink Targets
0%    Inline:Emphasis
0%    Inline:Strong
0%    Inline:Interpreted text
0%    Inline:Literals
0%    Inline:Embedded URIs
0%    Inline:Internal Targets
0%    Inline:Footnote References
0%    Inline:Citation References
0%    Inline:Substitution References
0%    Inline:Standalone HyperlinkS
0%    Units:Length
0%    Units:Percentage
0%    Directive:Admonitions
0%    Directive:Image
0%    Directive:Figure
0%    Directive:Topic
0%    Directive:Sidebar
0%    Directive:Code                 Needs a syntax parser for many programming languages.
0%    Directive:Math
0%    Directive:Rubric
0%    Directive:Epigraph
0%    Directive:Highlights
0%    Directive:Pull-quote
0%    Directive:Compound Paragraph
0%    Directive:Container
0%    Directive:Table
0%    Directive:CSV Table
0%    Directive:List Table
0%    Directive:Contents             Table of contents.
0%    Directive:Secnum               Automatic section numbering.
0%    Directive:Header
0%    Directive:Footer
0%    Directive:Meta                 HTML Meta Tags
0%    Directive:Replacement Text
0%    Directive:Unicode              Numerical unicode character codes.
0%    Directive:Date
0%    Directive:Class                For HTML output.
====  ============================== ======
