============================================================
Implementation of the Go reStructuredText Parser and Tooling
============================================================
:Modified: Thu Sep 14 11:38 2017

--------
Overview
--------

Implementation details of the Go reStructuredText Parser are documented here. This document details setting up new tests and
tips for debugging the parser engine.

.. contents::

-------
Roadmap
-------

0.1 (In Progress)
=================

* Parsing support:

  - Hyperlink Reference

  - inline markup

  - bullet list

* CLI: rst2html: Translate document to HTML

* Render basic documents using Hugo (https://github.com/spf13/hugo/issues/1436)

0.2
===

* Parsing support:

  - CODE directive

  - Enumerated list

  - Literal blocks

* Syntax highlighting using Sourcegraph's highlighting engine

0.3
===

* Parsing support:

  - Blockquote

  - definition list

0.4
===

* CLI: confluence2rst: Tool to convert a Confluence page into reStructuredText

0.5
===

* CLI: rst2confluence: Tool to convert reStructuredText to Confluence markup

-------
Testing
-------

New tests are added using a combination of JSON and simple Go. The naming and directory structure of the tests are important.

Tests are imported from docutils then implemented in the parser. This is a semi-manual process.

Conversion
==========

In the reference implementation of reStructuredText, the tests are implemented in a "psuedo xml". Not every language has a
psuedo xml parser (like Go), so work was started to translating the tests into JSON. The tests for the Go reStructuredText
parser are more accurate because the tokenizer tests are included. The tests tell the parser *how* to parse a document.

Converting and adding new tests is currently a manual process.

Some of these tests have been changed to conform to the parser and lexer provided by the go-rst package. The docutils parser
is much more complex, so some test results don't apply to the go-rst parser.

Import
------

The docutils reference implementation contains hundreds of tests. They can be seen at:

http://repo.or.cz/docutils.git/tree/HEAD:/docutils/test/test_parsers/test_rst

Status
------

The following table details tests that have been imported and implemented.

======================================  ========  ===========
Test file                               Imported  Implemented
test_SimpleTableParser.py               NO        NO
test_TableParser.py                     NO        NO
test_block_quotes.py                    YES       NO
test_bullet_lists.py                    YES       NO
test_character_level_inline_markup.py   NO        NO
test_citations.py                       NO        NO
test_comments.py                        YES       IN PROGRESS
test_definition_lists.py                YES       NO
test_doctest_blocks.py                  NO        NO
test_east_asian_text.py                 NO        NO
test_enumerated_lists.py                YES       NO
test_field_lists.py                     NO        NO
test_footnotes.py                       NO        NO
test_functions.py                       NO        NO
test_inline_markup.py                   YES       IN PROGRESS
test_interpreted.py                     NO        NO
test_interpreted_fr.py                  NO        NO
test_line_blocks.py                     NO        NO
test_literal_blocks.py                  YES       NO
test_option_lists.py                    NO        NO
test_outdenting.py                      NO        NO
test_paragraphs.py                      YES       YES
test_section_headers.py                 YES       YES
test_substitutions.py                   NO        NO
test_tables.py                          NO        NO
test_targets.py                         YES       IN PROGRESS
test_transitions.py                     NO        NO
======================================  ========  ===========

Test Layout Overview
====================

Test names are serialized. The best effort was made to get the tests sorted in order of importance for parser implementation.
Each test name includes a "double dot quad" identifier—this allows for incrementally adding additional variations of a single
test while keeping the file names unique.

There are currently three files per test: the rst file, the expected lexer output "items.json", and the expected parser
output "nodes.json".

Test names contain the words "good" or "bad" to indicate how the parser is expected to parse the test. Tests marked with
"good" are proper syntax and are expected to parse correctly. Tests marked with "bad" usually result in the parser generating
a system messages.

Tests that are not implemented by the go-rst parser have "-xx" appended to the name of the test. Unimplemented tests are also
tracked in the corresponding Go test file and are blocked from being run by the Go test program with a global variable.

testdata directory layout and naming format
-------------------------------------------

This is important! The names and directory layout of these files are used to generate the Go testing code.

::

  ▾ testdata/
    ▸ 00-test-comment/
    ▸ 01-test-reference-hyperlink-targets/
    ▸ 02-test-paragraph/
    ▸ 03-test-blockquote/
    ▾ 04-test-section/
      ▸ 00-section-title/
      ▾ 01-section-title-overline/
          ...
          04.01.05.00-bad-incomplete-section-items.json
          04.01.05.00-bad-incomplete-section-nodes.json
          04.01.05.00-bad-incomplete-section.rst
          ....

Element ID numbers
~~~~~~~~~~~~~~~~~~

Individual elements are numbered sequentially, in the order of importance needed to render a usable document.

The official reStructuredText spec is not divided into numbered sections for implementation writers (like the commonmark
spec) so this order is at best an approximation.

Good tests and bad tests
++++++++++++++++++++++++

Good tests are expected to produce valid output from the parser. Bad tests result in the parser returning error messages,
also called "System Messages" in reStructuredText. 

Test names
~~~~~~~~~~

**04.01.05.00-bad-incomplete-section.rst** can be broken down in the following way:

1. The first double digit, `04` in the example indicates the group the test belongs to.

   The parent directory (element group) contains this number.

#. The second double digit, `01` indicates the first sub group of the test

#. `05` indicates the second sub group of the test

   The second sub group groups tests that are similar, but just a little different from each other.

   For example, `06.01.00.XX` would be the first sub-subgroup for regular strong elements in a paragraph. `06.01.02.XX` would
   group quoted strong elements.

#. The fourth and last double digit, `00` indicates the variation of the test

#. The name comes after the ID

   Names should be descriptive and short. `two-paragraphs-three-lines`, `strong-asterisk` and `strong-across-lines` are good
   examples of names.

#. Tests that are not yet implemented are denoted with `-xx` appended to the end of the test name

   Un-implemented tests are also blocked from running in the Go test files using a global variable.

items.json
----------

The items.json files describes tokens generated by the lexer. It contains a json array of the following object:

.. code:: json

    {
        "id": 9,
        "type": "itemInlineEmphasis",
        "text": "emphasis",
        "startPosition": 5,
        "line": 4,
        "length": 8
    }

id
  A sequential numerical identifier given to the lexed item.

type
  The type of token found by the lexer.

text
  The actual text of the token. This excludes the actual markup. For emphasized text written in the document as
  ``*emphasis``, the text would only contain ``emphasis``.

startPosition
  The start position in the line of the lexed token. This is the byte position in the line of text.

line
  The line location within the file.

length
  The actual length of the lexed token. This is the number of runes in the text and is not the length in bytes.

nodes.json
----------

This files describes the document tree generated by the parser and roughly has the same fields as items.json.

For example, `00.00.00.00-comment-nodes.json` contains:

.. code:: json

   [
       {
           "type": "NodeComment",
           "text": "A comment.",
           "startPosition": 4,
           "line": 1,
           "length": 10
       },
       {
           "type": "NodeParagraph",
           "nodeList": [
               {
                   "type": "NodeText",
                   "text": "Paragraph.",
                   "startPosition": 1,
                   "line": 3,
                   "length": 10
               }
           ]
       }
   ]

Notice a paragraph node contains child nodes.

Import a test suite
===================

The docutils reference implementation contains hundreds of tests, as of 2017-06-11 not all of the tests have been converted
to JSON.

.. note:: If importing tests from docutils, it's best to import all the tests in one commit so that tests are not forgotten.


Get the docutils code
---------------------

Download the docutils reference implementation from http://repo.or.cz/docutils.git

Open the project in a text editor and go to the `test/test_parsers/test_rst` directory

   http://repo.or.cz/docutils.git/tree/HEAD:/docutils/test/test_parsers/test_rst

The first test
--------------

See the `Status`_ table for a quick overview of import/implementation status from the docutils reference parser. The
`testdata` also contains empty directories that will indicate which tests have not yet been imported from the docutils test
suite.

For this example, the Option List test suite will be imported.

Open `test_option_lists.py`_, the file begins with a Python array containing the reStructuredText source and the pseudo XML:

.. code:: python

    totest['option_lists'] = [
    ["""\
    Short options:

    -a       option -a

    -b file  option -b

    -c name  option -c
    """,
    """\
    <document source="test data">
        <paragraph>
            Short options:
        <option_list>
            <option_list_item>
                <option_group>
                    <option>
                        <option_string>
                            -a
                <description>
                    <paragraph>
                        option -a
            <option_list_item>
                <option_group>
                    <option>
                        <option_string>
                            -b
                        <option_argument delimiter=" ">
                            file
                <description>
                    <paragraph>
                        option -b
            <option_list_item>
                <option_group>
                    <option>
                        <option_string>
                            -c
                        <option_argument delimiter=" ">
                            name
                <description>
                    <paragraph>
                        option -c
    """],

We are primarily concerned with the reStructuredText source. We can always generate the psuedo XML separately with the
`rst2psuedoxml` docutils CLI tool.

Creating the test files
~~~~~~~~~~~~~~~~~~~~~~~

Next, create the test files that will contain the reStructuredText source for this test. These files will be used to generate
the Go testing code.

Navigate to the `testdata` directory, notice the `10-test-list-option` already exists. Now take a look at the spec, notice
there are at least four syntaxes option lists can use:

  There are several types of options recognized by reStructuredText:

  * Short POSIX options consist of one dash and an option letter.
  * Long POSIX options consist of two dashes and an option word; some systems use a single dash.
  * Old GNU-style "plus" options consist of one plus and an option letter ("plus" options are deprecated now, their use discouraged).
  * DOS/VMS options consist of a slash and an option letter or word.

  -- reStructuredText Specification

With this information, we can expect four subgroups for these tests. Here is the directory structure that should be created::

   ▾ 10-test-list-option/
     ▸ 00-short-posix/
     ▸ 01-long-posix/
     ▸ 02-gnu-plus/
     ▸ 03-dos/

Now that the directory structure is setup, we can create the files for our first test:

.. code:: console

   $ touch 10-test-list-option/00-short-posix/10.00.00.00-three-short-options{-nodes.json,-items.json,.rst}

Our directory structure now looks like::

  ▾ 10-test-list-option/
    ▾ 00-short-posix/
        10.00.00.00-three-short-options-items-xx.json
        10.00.00.00-three-short-options-nodes-xx.json
        10.00.00.00-three-short-options.rst

Open `10.00.00.00-three-short-options.rst` and copy the reStructuredText source from above into that file. Use the
`rst2psuedoxml` command to ensure the reStructuredText source file is valid. The command should return the same psuedo xml
shown in the other part of the test suite above:

.. code:: console

   $ rst2pseudoxml --halt=5 10-test-list-option/00-short-posix/10.00.00.00-three-short-options.rst

In this case, the output is the same, so the reStructuredText source is good.

Creating the Go test files
~~~~~~~~~~~~~~~~~~~~~~~~~~

The Go test code tests named ``rst_test.go`` in each of the lexer and parser packages.

The files can be regenerated using the ``go generate`` command::

    go generate

Using a Go Test functions with a unique names makes it possible to use the filtering capabilities of the Go test binary as
shown below.

View the ``rst_test.go`` file for the lexer and parser.

This test begins by geting the absolute path to the test using the name of the test without the `.rst` extension. The test
file is read and tokenized and results are checked against expected lexer tokens file
(`10.00.00.00-three-short-options-items.json`) using the `JSON diff library JD`_. The JSON diff library outputs in a special
"diff language" which is simple enough to learn. See the examples on the libraries Github page.

The environment variable check makes it possible to skip tests that are not implemented. This is used in Travis CI and
Coveralls to prevent the build and test from failing.

The parser test also compares the parser output to the expected parse nodes file
(`11.00.00.00-three-short-options-nodes.json`) by diffing JSON objects.

Running the tests
~~~~~~~~~~~~~~~~~

To run our tests explicitly, we can run the test directly with:

.. code:: console

   $ go test -v ./pkg/token -test.run=".*10.00.00.00.Lex.*" -debug
   === RUN   Test_10_00_00_00_LexOptionListGood_NotImplemented
   --- FAIL: Test_10_00_00_00_LexOptionListGood_NotImplemented (0.01s)
           token_test.go:76: "testdata/10-test-list-option/00-short-posix/10.00.00.00-three-short-options-items.json" is empty!
   FAIL
   exit status 1
   FAIL    github.com/demizer/go-rst/pkg/token     0.010s

Since the expected tokens (items) have not been written, this test fails as expected. Now run the parser test:

.. code:: console

   $ go test -v ./pkg/parser -test.run=".*10.00.00.00.Parse.*" -debug
   === RUN   Test_10_00_00_00_ParseOptionListShortGood_NotImplemented
   --- FAIL: Test_10_00_00_00_ParseOptionListShortGood_NotImplemented (0.00s)
       parse_test.go:104: "testdata/10-test-list-option/00-short-posix/10.00.00.00-three-short-options-nodes.json" is empty!
       FAIL
       exit status 1
       FAIL    github.com/demizer/go-rst/pkg/parser    0.007s

It fails as expected.

JSON Diff output
++++++++++++++++

Edit `10.00.00.00-three-short-options-items.json` and add some dummy tokens:

.. code:: json

   [
       {
           "id": 1,
           "type": "itemCommentMark",
           "text": "..",
           "line": 1,
           "length": 2,
           "startPosition": 1
       },
       {
           "id": 1,
           "type": "itemCommentMark",
           "text": "..",
           "line": 1,
           "length": 2,
           "startPosition": 1
       },
       {
           "id": 1,
           "type": "itemCommentMark",
           "text": "..",
           "line": 1,
           "length": 2,
           "startPosition": 1
       },
       {
           "id": 1,
           "type": "itemCommentMark",
           "text": "..",
           "line": 1,
           "length": 2,
           "startPosition": 1
       },
       {
           "id": 1,
           "type": "itemCommentMark",
           "text": "..",
           "line": 1,
           "length": 2,
           "startPosition": 1
       },
       {
           "id": 1,
           "type": "itemCommentMark",
           "text": "..",
           "line": 1,
           "length": 2,
           "startPosition": 1
       },
       {
           "id": 1,
           "type": "itemCommentMark",
           "text": "..",
           "line": 1,
           "length": 2,
           "startPosition": 1
       },
       {
           "id": 1,
           "type": "itemCommentMark",
           "text": "..",
           "line": 1,
           "length": 2,
           "startPosition": 1
       }
   ]

Run the test again, it will fail with::

   --- FAIL: Test_10_00_00_00_LexOptionListGood_NotImplemented (0.01s)
           token_test.go:53: The Actual Lexer Tokens and the Expected Lexer tokens do not match!
                   @ [7,"id"]
                   - 1
                   + 8
                   @ [7,"length"]
                   - 2
                   + 0
                   @ [7,"line"]
                   - 1
                   + 7
                   @ [7,"startPosition"]
                   ...

Most of the output has been cut off except for the start of the output. See the Github project page for the JD library on how to read the output.

And now the test has been imported into the Go reStructuredText Test Suite.

Import all the tests
--------------------

It's important to import all the Option List tests in this fashion so that we don't forget any tests!

The next section shows how to implement parsing make these tests pass.

Implementing a test
===================

Adding a new test is easy.

Debugging
=========

Debugging go-rst can be difficult and time consuming at times, especially if adding a new feature. Here are some tricks to
make the process a little easier.

Use the logger
--------------

The following command Will show lexer and parser output in debug format::

  go test -v ./pkg/parser -test.run=".*06.00.05.00.*_Parse.*" parse -debug

The lexer output can become annoying when trying to debug the parser. To exclude the output, use the ``-exclude`` argument:

::

  GO_RST_SKIP_NOT_IMPLEMENTED=1 go test -v ./pkg/parser -test.run=".*06.00.05.00.*_Parse.*" -debug -exclude=lexer

-------
Tooling
-------

To be written...

.. _test_option_lists.py: http://repo.or.cz/docutils.git/blob/HEAD:/docutils/test/test_parsers/test_rst/test_option_lists.py
.. _JSON diff library JD: https://github.com/josephburnett/jd
