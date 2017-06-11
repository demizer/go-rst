================================
go-rst - reStructuredText for Go
================================

.. image:: https://travis-ci.org/demizer/go-rst.svg?branch=master
    :target: https://travis-ci.org/demizer/go-rst
.. image:: https://coveralls.io/repos/github/demizer/go-rst/badge.svg?branch=master
    :target: https://coveralls.io/github/demizer/go-rst?branch=master
.. image:: https://goreportcard.com/badge/github.com/demizer/go-rst
    :target: https://goreportcard.com/report/github.com/demizer/go-rst
.. image:: https://godoc.org/github.com/demizer/go-rst?status.svg
    :target: http://godoc.org/github.com/demizer/go-rst

A reStructuredText parser for the Go Programming Language.

------
Status
------

.. The following is auto-generated using the tools/update-progress.sh
.. STATUS START

go-rst implements **9%** of the official specification (26 of 283 Items)

.. STATUS END

See https://github.com/demizer/go-rst/tree/master/doc/README.rst for complete details.

Road Map
========

0.1 (In Progress)
-----------------

* Parsing support:
  - Hyperlink Reference
  - inline markup
  - bullet list
* CLI: rst2html: Translate document to HTML
* Render basic documents using Hugo (https://github.com/spf13/hugo/issues/1436)

0.2
---

* Parsing support:
  - CODE directive
  - Enumerated list
  - Literal blocks
* Syntax highlighting using Sourcegraph's highlighting engine

0.3
---

* Parsing support:
  - Blockquote
  - definition list

0.4
---

* CLI: confluence2rst: Tool to convert a Confluence page into reStructuredText
* CLI: rst2confluence: Tool to convert reStructuredText to Confluence markup

-----------------
How to contribute
-----------------

See the `doc` directory for more documentation and tip and tricks.

* **Convert tests into JSON**

  The docutils tests are implemented in a "psuedo xml" which is non-standard.
  Translating the tests into JSON has the benefit of making the reStructuredText
  tests programming language neutral so that reStructuredText parsers can be
  implemented in other programming languages. See
  https://github.com/demizer/go-rst/tree/master/testdata
  for more information.

* **Implement an element**

  Implement an element from the list above.

* **Write documentation**

  All projects need good documentation.

* **Test and report**

  Not actually possible in the current state, but using the library and writing
  bug reports is always helpful.
