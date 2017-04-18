================
The Big Refactor
================

-----
About
-----

Deep flaws in the lexer/parser design were discovered when trying to integrate test 06.00.03.00

1. tokenizing is done one slice at a time, instead of one unicode code point at a time.

   This makes it difficult to handle unicode combining characters.

   A workaround was attemted by trying to normalize the input text, but this is not ideal. Mainly because of decomposition;
   '\u2000' decomposes to '\u2002` when normalized with the NFC form (`Unicode Normalization Forms`_), which according to the
   unicode spec are functionally identical. For a document syntax such as reStructuredText, this is not good because we want
   to preserve the original intent.

#. The project directory structure is becoming unweildly.

   It's hard to understand the "flow" of the library. This will make it hard to gain new contributors.

   Also, having separate files for `lexer_*` and `parser_*` add increased complexity.

#.

#. Plugin support (directives)

   Need to start thinking about this...

   The https://github.com/golang-commonmark/markdown repository has a great design.

----
TODO
----

This might take a while. Let's plan.

1. rename files to remove "parse\_" and "lexer\_"

#. Use lifo stack for lexer/parser state instead of saveState()

#. Lexer should use scanner.scan()

   * Should loop by unicode code points

   * Remove the need for norm.NFC

#. remove EOL (utf.RuneError)

#. remove the "parse" directory

.. _Unicode Normalization Forms: http://unicode.org/reports/tr15/
