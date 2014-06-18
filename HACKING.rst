=============================================================
Guidlines for Contributing to the Go reStructuredText Project
=============================================================
:Modified: Wed Jun 18 00:58 2014

---------------
Getting started
---------------

The go-rst lexer and parser are based off of the text/template lexer and
parser. The principle of each is the same. For a detailed introduction to how
they work, see the video `Lexical scanning in Go
<https://www.youtube.com/watch?v=HxaD_trXwRE>`_.

There are some key differences:

1. lexing is done line by line, character by character.
#. The AST generated is much simpler.
#. The parser is much larger and more complicated due to the complexity of
   reStructuredText.

General layout
==============

The path of execution through the package for parsing is as follows:

1. Parse() is called passing input.
#. Input text is normalized to condense unicode characters.
#. Tree.Parse() is called.
#. Tree.Parse() initiates the lexer and calls startParse().
#. Tree.Parse() calls Tree.parse() which starts the parsing.
#. Tree.parse() blocks waiting for a token on the receive channel.
#. lexer.lexStart() is called to start the lexing.
#. lexer.emit() emits a token on the channel, sending a pointer to the item
   created.
#. Tree.parse() receives the pointer to item and if it is actionable
   immediately, creates a Node and appends it to Tree.Nodes, otherwise it looks
   ahead for the next tokens to build a proper Node. Pointers to tokens
   received from the lexer not used immediately are saved to the Tree.token
   buffer.
#. Once the lexer is finished lexing, the send channel is closed.
#. The parser uses the remaining tokens in the buffer and returns the parse
   Tree.

-------
Testing
-------

To run the projects tests, simply use::

  go test

To run a specific test, use::

  go test -test.run <test_name>

The name used can be all or some of the name. For example, to run the first
lexer test for section headers, ``<test_name>`` can be either
``TestLexSectionTitleGood0000`` or ``TitleGood0000``.

To run a test with debug output, use::

  go test -test.run <test_name> -debug

To examine test coverage, use::

  go test -coverprofile c.out && go tool cover -html c.out

Testdata
========

Test data for all section tests is contained in the testdata directory. To
understand how the test data is used, please see the README.rst in the testdata
directory.

---------
Debugging
---------

Debugging is often necessary when adding new parts to the lexer or parser.
go-rst imports the `go-spew library <https://github.com/davecgh/go-spew>`_ for
pretty printing lexer Items or parser Nodes. The spew ConfigState is stored in
the ``spd`` global variable, so dumping objects to stdout is as simple as::

  spd.Spew(<input>)

Besides using go-spew, debug output can be sent to stdout (or anywhere with
some changes) using the `go-elog package
<https://github.com/demizer/go-elog>`_. go-elog is a replacement library for
the standard log package that adds logging levels and support for other output
streams (io.Writer). Logging debug output is easy using the Debug, Debugln, and
Debugf functions::

  log.Debugln("Hello!")

The log.Debugln() output contains the file and line number of where the debug
output was called from.
