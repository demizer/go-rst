some punctuation is allowed around inline markup, e.g.
/*emphasis*/, -*emphasis*-, and :*emphasis*: (delimiters),
(*emphasis*), [*emphasis*], <*emphasis*>, {*emphasis*} (open/close pairs)
*emphasis*., *emphasis*,, *emphasis*!, and *emphasis*\ (closing delimiters),

but not
)*emphasis*(, ]*emphasis*[, >*emphasis*>, }*emphasis*{ (close/open pairs),
(*), [*], '*' or '"*"' ("quoted" start-string),
x*2* or 2*x* (alphanumeric char before),
\*args or * (escaped, whitespace behind start-string),
or *the\* *stars\* *inside* (escaped, whitespace before end-string).

However, '*args' will trigger a warning and may be problematic.

what about *this**?
