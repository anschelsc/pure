\ Structures:
\ single: eval ast. char
\ pair: eval ast. 'left 'right

\ block: apply freeze 'ast
\ s1: apply freeze 'x
\ s2: apply freeze 'x 'y
\ k1: apply freeze 'x
\ s, k, and i are premade
: eval ( ast -- func) dup @ execute ;
: ast. ( ast) dup cell+ @ execute ;

: apply ( func ast -- func) over @ execute ;
: freeze ( func -- ast) dup cell+ @ execute ;

variable sast variable kast variable iast

variable 'pair
: pair 'pair @ execute ;

: freeze-s ( func -- ast) drop sast @ ;
: freeze-k ( func -- ast) drop kast @ ;
: freeze-i ( func -- ast) drop iast @ ;

: freeze-s1 ( func -- ast) 2 cells + @ sast @ pair ;
: freeze-s2 ( func -- ast) 2 cells + dup @ sast @ pair swap cell+ @ swap pair ;
: freeze-k1 ( func -- ast) 2 cells + @ kast @ pair ;

: freeze-block ( func -- ast) 2 cells + @ ;

: apply-s2 ( func ast -- func)
   over 2 cells + @ ( func z x)
   eval over apply ( func z `xz)
   swap rot 3 cells + @ ( `xz z y)
   pair apply ; ( ``xz`yz)
: apply-s1 ( func ast -- func) here -rot ['] apply-s2 , ['] freeze-s2 ,  swap 2 cells + @ , , ;
: apply-s ( func ast -- func) nip here swap ['] apply-s1 , ['] freeze-s1 ,  , ;
create sfunc ' apply-s , ' freeze-s ,

: apply-k1 ( func ast -- func) drop 2 cells + @ eval ;
: apply-k ( func ast -- func) nip here swap ['] apply-k1 , ['] freeze-k1 ,  , ;
create kfunc ' apply-k , ' freeze-k ,

variable 'blockify
: apply-block ( func ast -- func) swap 2 cells + @ pair 'blockify @ execute ;
: blockify ( ast -- func) here swap ['] apply-block , ['] freeze-block ,  , ;
' blockify 'blockify !

: apply-i ( func ast -- func) nip eval ;
create ifunc ' apply-i , ' freeze-i ,

: eval-p ( ast -- func) 2 cells + dup @ eval swap cell+ @ apply ;
: eval-s ( ast -- func) dup 2 cells + @
   dup [char] s = if 2drop sfunc exit then
   dup [char] k = if 2drop kfunc exit then
   [char] i = if drop ifunc exit then
   blockify ;

: ast.-s ( ast) 2 cells + @ emit ;
: ast.-p ( ast) [char] ` emit 2 cells + dup @ ast. cell+ @ ast. ;

: single ( char -- ast) here swap ['] eval-s , ['] ast.-s ,  , ;
: pair ( right left -- ast) here -rot ['] eval-p , ['] ast.-p ,  , , ;
' pair 'pair !

char s single sast !
char k single kast !
char i single iast !

variable vcount
0 vcount !
: ensure ( --) vcount @ 0< abort" Syntax error." ;
: up ( --) 1 vcount +! ;
: down ( --) -1 vcount +! ensure ;
: end-check ( --) vcount @ 1- abort" Syntax error." ;

: consume ( adr -- ast) c@ dup [char] ` = if down drop pair exit then up single ;
: parse ( adr count -- ast) over + 1- do i consume -1 +loop end-check ;
