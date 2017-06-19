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

See `implementation status`_ for a breakdown of what's implemented.

There is also a `Road Map`_.

-----
Usage
-----

This library does not have any functionality beyond running parser tests.

Tests
=====

From the root of the project,

::

    GO_RST_SKIP_NOT_IMPLEMENTED=1 go test -v ./pkg/...

There are many tests that are imported from docutils, but not implemented yet.

-----------------
How to contribute
-----------------

See the `doc`_ directory for more documentation and tip and tricks.

* **Convert tests into JSON**

  `Import a test suite from docutils`_

* **Implement parsing for a document element**

  `How to implement an element`_

.. _Road Map: https://github.com/demizer/go-rst/blob/master/doc/implementation.rst#roadmap
.. _implementation status: https://github.com/demizer/go-rst/tree/master/doc/README.rst
.. _Doc: https://github.com/demizer/go-rst/tree/master/doc
.. _Import a test suite from docutils: https://github.com/demizer/go-rst/tree/master/doc/implementation.rst#testing
.. _How to implement an element: https://github.com/demizer/go-rst/blob/master/doc/implementation.rst#implementing-a-test
