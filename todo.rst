=====================
go-rst - Things to Do
=====================

-----
INBOX
-----

* How to create a language for LLVM:

  http://llvm.org/docs/tutorial/LangImpl1.html#language

* Why using regex for parsers is not encouraged:

  http://commandcenter.blogspot.com/2011/08/regular-expressions-in-lexing-and.html

--------------------------------------------
Make a scanner to parse rst file into tokens
--------------------------------------------
:Added: Wed Feb 05 23:08 2014

+ [X] Implement basic test_section_headers.dat parsing in lex_test.go

  - [X] Load file

  - [X] Parse file into sections

  - [X] Parse json into arbitrary datastructure

+ [I] Implement basic tokenizer for the three tests in test_section_headers.dat

  - [_] TODO
