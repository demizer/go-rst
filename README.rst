================================
go-rst - reStructuredText for Go
================================
:Modified: Thu Nov 27 21:44 2014

A reStructuredText parser for and implemented in the Go programming language.

**This library is far from complete and barely usable.**


=================
How to contribute
=================

* **Convert tests into JSON**

  The docutils tests are implemented in a "psuedo xml" which is non-standard.
  Translating the tests into JSON has the benefit of making the reStructuredText
  tests programming language neutral so that reStructuredText parsers can be
  implemented in other programming languages. See
  https://github.com/demizer/go-rst/tree/master/testdata
  for more information.

* **Implement an element**

  Implement an element from the list above.

* **Write some documentation**

  All projects need good documentation!

* **Test and report**

  Not actually possible in the current state, but using the library and writing
  bug reports is always helpful.

============
Contributors
============

These people have donated their valuable time in contributing to this library
and are here graciously recognized for their contributions!

* **Jesus Alvarez**

------
Status
------

Work is currently underway on translating the docutils tests from psuedo xml
into JSON.

+-------------------------------------------------+
|          **Whitespace 10% Complete**            |
+---------+------------------------+--------------+
|**Done** | **Sub Element**        | **Note**     |
+---------+------------------------+--------------+
| 10%     | tab-to-space           |              |
+---------+------------------------+--------------+
|  0%     | form-feed-to-space     |              |
+---------+------------------------+--------------+
|  0%     | vertical-tab-to-space  |              |
+---------+------------------------+--------------+
|          **Whitespace 10% Complete**            |
+---------+------------------------+--------------+
|**Done** | **Sub Element**        | **Note**     |
+---------+------------------------+--------------+
| 10%     | tab-to-space           |              |
+---------+------------------------+--------------+
|  0%     | form-feed-to-space     |              |
+---------+------------------------+--------------+
|  0%     | vertical-tab-to-space  |              |
+---------+------------------------+--------------+

