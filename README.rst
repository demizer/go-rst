================================
go-rst - reStructuredText for Go
================================
:Modified: Fri Jun 13 15:48 2014

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

Jesus Alvarez

------
Status
------

Work is currently underway on translating the docutils tests from psuedo xml
into JSON.

=========== ==========  ============================== ======
Last Update Completion  Element                        Detail
=========== ==========  ============================== ======
Jun 13 2014 0%          Exported API                   Not even designed yet. Look to other text file parsers for inspiration. Suggestions welcome!
Jun 13 2014 0%          HTML transpiler                Tree.Nodes can be easily walked to produce output html.
Jun 13 2014 0%          Demo homepage                  To be implemented in Go and hosted at restructuredtext.com
Jun 13 2014 5%          Documentation                  Most functions documented. Nothing on API.
Jun 13 2014 100%        Section Headers                All tests from docutils converted to JSON and implemented.
Jun 13 2014 5%          Transitions                    Basic lexing.
Jun 13 2014 25%         Paragraphs                     Basic lexing, no inline markup, no tests.
Jun 13 2014 0%          Bullet lists
Jun 13 2014 5%          Enumerated lists               Some lexing implemented.
Jun 13 2014 0%          Definition lists
Jun 13 2014 0%          Field lists
Jun 13 2014 0%          Biblio. fields
Jun 13 2014 0%          RCS keywords
Jun 13 2014 0%          Option lists
Jun 13 2014 5%          Literal blocks                 Some parsing.
Jun 13 2014 0%          Indented literal blocks
Jun 13 2014 0%          Quoted literal blocks
Jun 13 2014 0%          Line blocks
Jun 13 2014 10%         Block quotes                   Some parsing.
Jun 13 2014 0%          Doctest blocks                 Will use Go instead of Python.  Much lower priority than everything else.
Jun 13 2014 0%          Grid tables                    Will be gruesome to implement.
Jun 13 2014 0%          Simple tables
Jun 13 2014 0%          Footnotes
Jun 13 2014 0%          Auto-Numbered Footnotes
Jun 13 2014 0%          Auto-Symbol Footnotes
Jun 13 2014 0%          Mixed Auto Footnotes
Jun 13 2014 0%          Citations
Jun 13 2014 0%          Hyperlink targets
Jun 13 2014 0%          Anonymous Hyperlinks
Jun 13 2014 0%          Substitution Definitions
Jun 13 2014 2%          Comments                       Basic lexing and parsing.
Jun 13 2014 0%          Implicit Hyperlink Targets
Jun 13 2014 0%          Inline:Emphasis
Jun 13 2014 0%          Inline:Strong
Jun 13 2014 0%          Inline:Interpreted text
Jun 13 2014 0%          Inline:Literals
Jun 13 2014 0%          Inline:Embedded URIs
Jun 13 2014 0%          Inline:Internal Targets
Jun 13 2014 0%          Inline:Footnote References
Jun 13 2014 0%          Inline:Citation References
Jun 13 2014 0%          Inline:Substitution References
Jun 13 2014 0%          Inline:Standalone HyperlinkS
Jun 13 2014 0%          Units:Length
Jun 13 2014 0%          Units:Percentage
Jun 13 2014 0%          Directive:Admonitions
Jun 13 2014 0%          Directive:Image
Jun 13 2014 0%          Directive:Figure
Jun 13 2014 0%          Directive:Topic
Jun 13 2014 0%          Directive:Sidebar
Jun 13 2014 0%          Directive:Code                 Needs a syntax parser for many programming languages.
Jun 13 2014 0%          Directive:Math
Jun 13 2014 0%          Directive:Rubric
Jun 13 2014 0%          Directive:Epigraph
Jun 13 2014 0%          Directive:Highlights
Jun 13 2014 0%          Directive:Pull-quote
Jun 13 2014 0%          Directive:Compound Paragraph
Jun 13 2014 0%          Directive:Container
Jun 13 2014 0%          Directive:Table
Jun 13 2014 0%          Directive:CSV Table
Jun 13 2014 0%          Directive:List Table
Jun 13 2014 0%          Directive:Contents             Table of contents.
Jun 13 2014 0%          Directive:Secnum               Automatic section numbering.
Jun 13 2014 0%          Directive:Header
Jun 13 2014 0%          Directive:Footer
Jun 13 2014 0%          Directive:Meta                 HTML Meta Tags
Jun 13 2014 0%          Directive:Replacement Text
Jun 13 2014 0%          Directive:Unicode              Numerical unicode character codes.
Jun 13 2014 0%          Directive:Date
Jun 13 2014 0%          Directive:Class                For HTML output.
=========== ==========  ============================== ======
