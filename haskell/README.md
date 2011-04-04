Haskell
=======

This is the Haskell implementation of Pure. If you have [GHC][1], running `ghc --make pure` should build an executable called `pure`, which evaluates a Pure program from standard input and writes the result to standard output. You can alse do variable elimination (described in more detail in the Go `README`) by listing the variables on the command line:

	% ./pure nfx
	`f``nfx
	`s``s`ksk

[1]: http://www.haskell.org/ghc/