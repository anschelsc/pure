 Go
====

This implementation is in the Go programming language, which can be found online at <http://golang.org>. If you have Go installed, running `gomake` in this directory will build an executable called `pure`. `pure [filename]` evaluates a Pure program from `filename`, or from standard input if none is given.

Variable Elimination
-------- -----------

To eliminate a variable from a Pure program (what the Unlambda page calls [abstraction elimination][1]), run `pure -e v`, where v is the variable to eliminate. For example,

	% ./pure -e x
	`xx
	``sii

To test this, we can run

	% ./pure
	` ``sii x
	`xx

and see that the variable elimination was, indeed, correct.

`pure` can also eliminate multiple variables: the [Unlambda page][2] suggests ```^n^f^x`f``nfx``` as the successor function for Church Numerals. We can eliminate all three variables in one go with

	% ./pure -e nfx
	`f ``nfx
	`s``s`ksk

Indeed, this is the elimination listed on the Unlambda page itself.

[1]: http://www.madore.org/~david/programs/unlambda/#lambda_elim
[2]: http://www.madore.org/~david/programs/unlambda/#howto_num
