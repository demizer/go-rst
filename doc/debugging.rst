================
Debugging go-rst
================

Debugging go-rst can be difficult and time consuming at times, especially if adding a new feature. Here are some tricks to
make the process a little easier.

--------------
Use the logger
--------------


The test logging is configured in `parse_test.go`.

  gb test -v -test.run=".*03.02.07.00.*_Parse.*" parse -debug | grep -v "name=lexer"
  rst2pseudoxml testdata/03-test-section/03.01.03.00-section-bad-subsection-order.rst --halt=5
  gb test -v -test.run=".*03.01.03.00.*_Parse.*" parse -debug | grep -v "name=lexer" | ag "NodeList" --passthrough

  This will dump all output regardless of parsing errors. Very useful to see how the reference parser uses system messages.

  rst2pseudoxml testdata/03-test-section/03.00.04.00-section-bad-unexpected-titles.rst --halt=5

