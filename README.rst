================================
go-rst - reStructuredText for Go
================================

.. image:: https://drone.io/github.com/demizer/go-rst/status.png
    :target: https://drone.io/github.com/demizer/go-rst/latest
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

Current Work
------------

* Translating the docutils tests from psuedo xml into JSON.
* Inline markup parsing

TODO
----

See TODO.txt (or issues) for current active work. The following list is in order of importance, but not yet in progress:

* Integration with Hugo

  Need to finish inline-markup parsing to have something worth while for Hugo. See https://github.com/spf13/hugo/issues/1436

* rst2html conversion tool

* rst2confluence import/conversion tool

  Confluence is garbage, it would be so nice to be able to alleviate the pain using reStructuredText. This has the lowest
  priority.

-----------------
How to contribute
-----------------

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
