This test was originally from 02.00-lots-of-escaping-unicode.rst

It was the last portion, and according to the spec, whitespace is not allowed
after the emphasis start string. So this test should produce errors when
parsed.

\u00a1*examples*!\u00a0*\u00a0no-break-space\u00a0*.
