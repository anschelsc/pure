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

typedef union {
	AST *single;
	struct {
		AST *left, *right;
	} pair;
} FData;

typedef struct Func_S {
	FData data;
	struct Func_S (*apply)(FData, AST *);
} Func;

Func apply_block(FData, AST *);
Func apply_s(FData, AST *);
Func apply_s1(FData, AST *);
Func apply_s2(FData, AST *);
Func apply_k(FData, AST *);
Func apply_k1(FData, AST *);
Func apply_i(FData, AST *);

void fprint(FILE *, AST *);

AST *from_char(char);
AST *combine(AST *, AST *);

Func eval(AST *);
AST *freeze(Func);

Func apply(Func, AST *);

AST *parse(FILE *);

#endif
