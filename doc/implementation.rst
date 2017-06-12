============================================================
Implementation of the Go reStructuredText Parser and Tooling
============================================================
:Modified: Mon Jun 12 00:13 2017

--------
Overview
--------

Implementation details of the Go reStructuredText Parser are documented here. This document details setting up new tests and
tips for debugging the parser engine.

.. contents::

-------
Testing
-------

New tests are added using a combination of JSON and simple Go. The naming and directory structure of the tests are important.

Tests are imported from docutils then implemented in the parser. This is a manual process.

Import
======

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

Conversion
==========

In the reference implementation of reStructuredText, the tests are implemented in a "psuedo xml". Not every language has a
psuedo xml parser (like Go), so work was started to translating the tests into JSON. The tests for the Go reStructuredText
parser are more accurate because the tokenizer tests are included. The tests tell the parser _how_ to parse a document.

Converting and adding new tests is currently a manual process.

Some of these tests have been changed to conform to the parser and lexer provided by the go-rst package. The docutils parser
is much more complex, so some test results don't apply to the go-rst parser.

Test Layout Overview
====================

Test names are serialized, the best effort was made to get the tests sorted in order of importance. Each test comes with a
unique identifier embedded in the name. Each test file begins with a number syntax formatted with "double dot quad",
`00.00.00.00`. This is to allow for incrementally adding additional variations of a single test while keeping the file names
unique.

There are currently three files per test: the rst file, the expected lexer output "items.json", and the expected parser
output "nodes.json".

Test names contain the words "good" or "bad" to indicate how the parser is expected to parse the test. Tests marked with
"good" are proper syntax and are expected to parse correctly. Tests marked with "bad" usually result in the parser generating
a system messages.

Tests that are not implemented by the go-rst parser have "-xx" appended to the name of the test. Unimplemented tests are also
tracked in the corresponding Go test file and are blocked from being run by the Go test program with a global variable.

testdata directory layout and naming format
-------------------------------------------

::

  ▾ testdata/
    ▸ 00-test-comment/
    ▸ 01-test-reference-hyperlink-targets/
    ▸ 02-test-paragraph/01-good/
    ▸ 03-test-blockquote/
    ▸ 04-test-section/
    ▸ 05-test-literal-block/
    ▾ 06-test-inline-markup/
      ▸ 00-inline-markup-recognition-rules/
      ▾ 01-strong/
        ▾ 01-good/
            06.01.00.00-strong-items.json
            06.01.00.00-strong-nodes.json
            06.01.00.00-strong.rst
            06.01.01.00-strong-with-apostrophe-items.json
            06.01.01.00-strong-with-apostrophe-nodes.json
            06.01.01.00-strong-with-apostrophe.rst
            06.01.02.00-strong-quoted-items.json
            06.01.02.00-strong-quoted-nodes.json
            06.01.02.00-strong-quoted.rst
            06.01.03.00-strong-asterisk-items.json
            06.01.03.00-strong-asterisk-nodes.json
            06.01.03.00-strong-asterisk.rst
            06.01.03.01-strong-asterisk-items.json
            06.01.03.01-strong-asterisk-nodes.json
            06.01.03.01-strong-asterisk.rst
            06.01.04.00-strong-across-lines-items.json
            06.01.04.00-strong-across-lines-nodes.json
            06.01.04.00-strong-across-lines.rst
        ▾ 02-bad/
            06.01.00.01-strong-unclosed-items.json
            06.01.00.01-strong-unclosed-nodes-xx.json
            06.01.00.01-strong-unclosed.rst
            06.01.00.02-strong-unclosed-items.json
            06.01.00.02-strong-unclosed-nodes-xx.json
            06.01.00.02-strong-unclosed.rst
            06.01.03.02-strong-kwargs-items.json
            06.01.03.02-strong-kwargs-nodes-xx.json
            06.01.03.02-strong-kwargs.rst
      ▸ 02-emphasis/
      ▸ 03-literal/
      ▸ 04-reference/
      ▸ 05-embedded-uri/
      ▸ 06-embedded-aliases/
      ▸ 07-inline-targets/
      ▸ 08-footnote-reference/
      ▸ 09-citation-reference/
      ▸ 10-substitution-reference/
      ▸ 11-standalone-hyperlink/
    ▸ 07-test-list-bullet/
    ▸ 08-test-list-enumerated/

Element ID numbers
~~~~~~~~~~~~~~~~~~

Individual elements are numbered sequentially, in the order of importance needed to render a usable document.

The official reStructuredText spec is not divided into numbered sections for implementation writers (like the commonmark
spec) so this order is at best an approximation.

::

    ▸ 00-test-comment/
    ▸ 01-test-reference-hyperlink-targets/
    ▸ 02-test-paragraph/01-good/
    ▸ 03-test-blockquote/
    ▸ 04-test-section/
    ▸ 05-test-literal-block/

Test names
~~~~~~~~~~

`06.01.03.01-strong-asterisk.rst` can be broken down in the following way:

1. The first double digit, `06` in the example indicates the group the test belongs to.

   This number is the same as the number set as an element ID above.

#. The second double digit, `01` indicates the first sub group of the test.

   There are none for the hyperlink target tests, but the inline markup tests and section tests have plenty.

   For example, here is what the inline markup tests subgroups look like::

     ▾ 06-test-inline-markup/
       ▸ 00-inline-markup-recognition-rules/
       ▸ 01-strong/
       ▸ 02-emphasis/
       ▸ 03-literal/
       ▸ 04-reference/

#. The third double digit, `03` indicates the second sub group of the test.

   The third sub group groups tests that are similar, but just a little different from each other.

#. The fourth and last double digit, `01` indicates the variation of the test.

#. The name comes after the ID

   Names should be descriptive and short. `two-paragraphs-three-lines`, `strong-asterisk` and `strong-across-lines` follow
   these guidelines.

#. Tests that are not yet implemented are denoted with `-xx` appended to the end of the test name.

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

Converting an existing test
===========================

NOTE: See the table above for tests that have not yet been imported into go-rst

The docutils reference implementation contains hundreds of tests, as of 2017-06-11 not all of the tests have been converted
to JSON.

NOTE: If importing tests from docutils, it's best to import all the tests in one commit so that tests are not forgotten.

1. Download the docutils reference implementation from http://repo.or.cz/docutils.git

#. Open the project in a text editor and go to the `test/test_parsers/test_rst` directory

   http://repo.or.cz/docutils.git/tree/HEAD:/docutils/test/test_parsers/test_rst

#. Inspect the testdata directory of the go-rst and determine which tests are not already imported.

   See the _`Status` table for a quick overview of import/implementation status from the docutils reference parser.

Adding a new test
=================

Adding a new test is easy.

Test conflicts with reference implementation
============================================

While implementing the go-rst parser, differences found from the official implementation are noted here.

Differences are mostly related to the style of parsing as the default docutils parser engine is based off of regular
expresssions, and the go-rst parser is hand-written by the finesh artisans.

Test: 04.02.06.01-incomplete-sections-no-title.rst
--------------------------------------------------

From: docutils/test/test_parsers/test_rst/test_section_headers.py line: 787

The expected results by the docutils package do not make any sense at all.  It seems the test is only to make sure the parser
does not crash. So I modified the expected results to conform to the current output of the go-rst parser. Naturally the
output is very different.

Test: 06.02.01.00-emphasis-with-emphasis-apostrophe.rst
-------------------------------------------------------

From: docutils/test/test_parsers/test_rst/test_inline_markup.py line: 33

Tests apostrophe handling, I think... Not really sure of the purpose of this test.
rst2html shows the following output, which appears broken:

.. code:: html

   <p>l'<em>emphasis</em> with the <em>emphasis</em>' apostrophe.
   lu2019*emphasis* with the <em>emphasis</em>u2019 apostrophe.</p>

Test: 06.00.00.00-double-underscore.rst
---------------------------------------

From: http://repo.or.cz/w/docutils.git/blob/HEAD:/docutils/test/test_parsers/test_rst/test_inline_markup.py#l1594

The markup::

    text-*separated*\u2010*by*\u2011*various*\u2012*dashes*\u2013*and*\u2014*hyphens*.
    \u00bf*punctuation*? \u00a1*examples*!\u00a0*\u00a0no-break-space\u00a0*.

Tests recognition rules with unicode literals. \u00a0 is "No Break Space".

Output from rst2html.py (docutils v0.12)::

    <p>text-<em>separated</em>u2010*by*u2011*various*u2012*dashes*u2013*and*u2014*hyphens*.
    u00bf*punctuation*? u00a1*examples*!u00a0*u00a0no-break-spaceu00a0*.</p>

According to the reStructuredText spec, whitespace after an inline markup start string are not allowed, but this test clearly
shows that it is. The troublesome section is ``\u00a0*\u00a0no-break-space\u00a0*`` as the parser cannot detect the '*' start
string (based on the spec). As mentioned in the previous trouble item, the docutils parser does not correctly use unicode
literals.

I have modified this test to remove the troublesome section.

Test: 06.00.03.00-emphasis-wrapped-in-unicode.rst
-------------------------------------------------

The following test is clearly valid:

.. code:: reStructuredText

    text separated by
    *newline*
    or *space* or one of
    \xa0*NO-BREAK SPACE*\xa0,
    \u1680*OGHAM SPACE MARK*\u1680,

but the official docutils parser parses it incorrectly::

    <document source="test data">
        <paragraph>
            text separated by
            <emphasis>
                newline
            \n\
            or \n\
            <emphasis>
                space
            or one of
            \xa0
            <emphasis>
                NO-BREAK SPACE
            \xa0,
            \u1680
            <emphasis>
                OGHAM SPACE MARK
            \u1680,

go-rst parses it correctly:

.. code:: json

    [
        {
            "type": "NodeParagraph",
            "nodeList": [
                {
                    "type": "NodeText",
                    "text": "text separated by",
                },
                {
                    "type": "NodeInlineEmphasis",
                    "text": "newline",
                },
                {
                    "type": "NodeText",
                    "text": "or ",
                },
                {
                    "type": "NodeInlineEmphasis",
                    "text": "space",
                },
                {
                    "type": "NodeText",
                    "text": " or one of\n\u00a0",
                },
                {
                    "type": "NodeInlineEmphasis",
                    "text": "NO-BREAK SPACE",
                },
                {
                    "type": "NodeText",
                    "text": "\u00a0,\n\u1680",
                },
                {
                    "type": "NodeInlineEmphasis",
                    "text": "OGHAM SPACE MARK",
                },
                {
                    "type": "NodeText",
                    "text": "\u1680,",
                },
            ]
        }
    ]

Notice the the usage of `\n` to merge NodeText nodes. The official parser does this correctly for test 02.00.01.00, but fails
miserably on this test.

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

---------------------------
Document conversion tooling
---------------------------

To be written...
