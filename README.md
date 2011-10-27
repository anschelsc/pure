Pure
====

Pure is a purely functional subset of the [Unlamda programming language][1]. 
Specifically, it contains just the combinators s, k, and i. These have no 
side-effects whatsoever, including I/O. As a consequence, there is no concept 
of "order of evaluation". For this reason, Unlambda's "abstraction elimination" 
becomes much more efficient.

Directory Structure
--------- ---------

This directory contains a subdirectory for each language in which Pure is 
implemented. Specific documentation for each implementation is in the 
subdirectory.

Author
------

Everything in this directory is (c) 2011 Anschel Schaffer-Cohen, 
<anschelsc@gmail.com>, and licensed under the BSD-style license in the LICENSE 
file. If you want to contribute and/or comment, feel free to do so via 
[Github][2] or by sending me email.

[1]: http://www.madore.org/~david/programs/unlambda/
[2]: http://github.com/anschelsc/pure/
