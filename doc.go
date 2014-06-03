/*
A reStructuredText parser implemented in Go!

This implementation follows as close as possible to the reference docutils
implementation with the following noted (tentative) exceptions and
improvements:

Most likely will be faster than the Python implementation.

The parser/lexer is based on the text/template package.

No support for doctest blocks... Maybe some day with Go?

The docutils tests are being re-written as JSON objects--which means language
portability for the tests. Hopefully the rewritten tests can be used to create
more implementations of reStructuredText!

Minimal output formats; only HTML and PDF will be supported initially.  Maybe
ODF in the future. Conversion of the document object should be "straight
forward" like it is in docutils.

The code-block directive will have limited language syntax support initially.

There will be other yet unforeseen limitations or improvements.

Status

Last updated: Mon Jun 02 23:49 2014

Currently progress is moving forward nicely and the text/template lexer/parser
design is holding up pretty well with the increased demand of reStructuredText.
Most of the section tests are implemented and passing. Work is currently being
done on implementing the blockquote, paragraph, and inline markup tests. The
package should be ready to output HTML for basic documents in six months time.
*/
package rst
