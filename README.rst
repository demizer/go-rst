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

|

A reStructuredText parser for the Go Programming Language.

This project is not yet usable. See the Road Map below.

------
Status
------

.. The following is auto-generated using the tools/update-progress.sh
.. STATUS START

go-rst implements **9%** of the official specification (26 of 283 Items)

.. STATUS END

See `implementation status`_ for complete details.

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

0.5
---

* CLI: rst2confluence: Tool to convert reStructuredText to Confluence markup

-----------------
How to contribute
-----------------

See the `doc`_ directory for more documentation and tip and tricks.

* **Convert tests into JSON**

  `How to convert tests`_

* **Implement parsing for a document element**

  `How to implement an element`_

.. _implementation status: https://github.com/demizer/go-rst/tree/master/doc/README.rst
.. _Doc: https://github.com/demizer/go-rst/tree/master/doc
.. _How to convert tests: https://github.com/demizer/go-rst/tree/master/doc/implementation.rst#test-conversion
.. _How to implement an element: https://github.com/demizer/go-rst/tree/master/implementation.rst#elements
