#include <stdio.h>

typedef struct ast {
	char is_node;
	union {
		char c;
		struct {
			struct ast *left, *right;
		} t;
	} data;
	size_t count;
} ast;

ast *leaf(char);
ast *node(ast *, ast *);

void release(ast *);

ast *parse(FILE *);
void output(FILE *, ast *);

typedef struct path {
	ast *here;
	struct path *up;
} path;

// To travel the path:
void ascend(void);
void descend(void);

// Change the current ast
void replace(ast *);

void f_raw(void);
void f_s(void);
void f_s1(void);
void f_s2(void);
void f_k(void);
void f_k1(void);
void f_i(void);
