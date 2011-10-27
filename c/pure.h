#ifndef PURE
#define PURE

#include <stdio.h>

enum AST_t { CHAR, PAIR };

typedef struct AST_S {
	enum AST_t type;
	union {
		char c;
		struct {
			struct AST_S *left;
			struct AST_S *right;
		} pair;
	} val;
} AST;

enum Func_t { BLOCK, S, S1, S2, K, K1, I };

typedef struct {
	enum Func_t type;
	union {	//We don't need any data if type is S, K, or I
		AST *block;
		struct {
			AST *x;
		} s1;
		struct {
			AST *x;
			AST *y;
		} s2;
		struct {
			AST *x;
		} k1;
	} val;
} Func;

void fprint(FILE *, AST *);

AST *from_char(char);
AST *combine(AST *, AST *);

Func *eval(AST *);
AST *freeze(Func *);

Func *apply(Func *, AST *);

AST *parse(FILE *);

#endif