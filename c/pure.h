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

typedef struct Func_S {
	void *data;
	struct Func_S (*apply)(void *, AST *);
} Func;

Func apply_block(void *, AST *);
Func apply_s(void *, AST *);
Func apply_s1(void *, AST *);
Func apply_s2(void *, AST *);
Func apply_k(void *, AST *);
Func apply_k1(void *, AST *);
Func apply_i(void *, AST *);

void fprint(FILE *, AST *);

AST *from_char(char);
AST *combine(AST *, AST *);

Func eval(AST *);
AST *freeze(Func);

Func apply(Func, AST *);

AST *parse(FILE *);

#endif