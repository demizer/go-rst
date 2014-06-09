=========================================
The reStructuredText Test Data For GO-RST
=========================================
:Modified: Mon Jun 09 01:48 2014

These test files have been transcoded from the docutils "pseudo xml" format
into standard JSON.

Some of these tests have been changed to conform to the parser and lexer
provided by the go-rst package. The docutils parser is much more complex, so
some test results don't apply to the go-rst parser.

--------------------
Test Layout Overview
--------------------

* The tests are broken down into catagories, with each directory containing
  sub-catagories.
* There are currently three files per test: the rst file, the expected lexer
  output "items.json", and the expected parser output "nodes.json".
* The sub-directories of each category end with "good" or "bad" to indicate how
  the parser is expected to parse the test. Directories ending with "good" are
  proper syntax and are expected to be parsed correctly. Directories ending
  with "bad" usually means the parser is going to generate a system message or
  two.
* Each test file begins with a number syntax formatted with two leading digits,
  a decimal, and two trailing digits: "00.00" This is to allow for
  incrementally adding additional variations of a single test while keeping
  the file names unique.

---------------
Differing Tests
---------------

1. Test: test_section/06_title_with_overline_bad/03.01_incomplete_sections_no_title.rst

   From: docutils/test/test_parsers/test_rst/test_section_headers.py line: 787

   The expected results by the docutils package do not make any sense at all.
   It seems the test is only to make sure the parser does not crash. So I
   modified the expected results to conform to the current output of the go-rst
   parser. Naturally the output is very different.
