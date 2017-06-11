============================================================
Implementation of the Go reStructuredText Parser and Tooling
============================================================
:Modified: Sun Jun 11 10:46 2017

---------------
Test conversion
---------------

The docutils tests are implemented in a "psuedo xml" which is non-standard. Translating the tests into JSON has the benefit
of making the reStructuredText tests programming language neutral so that reStructuredText parsers can be implemented in
other programming languages.

To be written...

--------
Elements
--------

Tests not yet implemented are marked as such using "-xx" in the test name.

As of 2017-06-11 there are 336 unimplemented tests that have been imported from docutils.

Implementation
==============

To be written...

Debugging
=========

Debugging go-rst can be difficult and time consuming at times, especially if adding a new feature. Here are some tricks to
make the process a little easier.

Use the logger
--------------

The test logging is configured in `parse_test.go`.

  gb test -v -test.run=".*03.02.07.00.*_Parse.*" parse -debug | grep -v "name=lexer"
  rst2pseudoxml testdata/03-test-section/03.01.03.00-section-bad-subsection-order.rst --halt=5
  gb test -v -test.run=".*03.01.03.00.*_Parse.*" parse -debug | grep -v "name=lexer" | ag "NodeList" --passthrough

  This will dump all output regardless of parsing errors. Very useful to see how the reference parser uses system messages.

  rst2pseudoxml testdata/03-test-section/03.00.04.00-section-bad-unexpected-titles.rst --halt=5

