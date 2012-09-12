#include <stdlib.h>
#include <assert.h>
#include <ctype.h>

#include "pure.h"

path *pc;
void (*todo)(void);
ast *program;

int main() {
	program = parse(stdin);
	pc = malloc(sizeof *pc);
	pc->up = NULL;
	pc->here = program;
	todo = &f_raw;
	while (pc)
		todo();
	output(stdout, program);
	printf("\n");
	release(program);
}

void output(FILE *f, ast *program) {
	if (program->is_node) {
		fputc('`', f);
		output(f, program->data.t.left);
		output(f, program->data.t.right);
		return;
	}
	fputc(program->data.c, f);
}

ast *parse(FILE *f) {
	char next;
	while (isspace(next = fgetc(f)))
		;
	switch (next) {
	case EOF:
		return NULL;
	case '`':
		(void) 0; // Because a label can't be attached to a declaration.
		ast *left = parse(f);
		ast *right = parse(f);
		if (!(left && right))
			return NULL;
		return node(left, right);
	default:
		return leaf(next);
	}
}

void release(ast *a) {
	if (--a->count)
		return;
	if (a->is_node) {
		release(a->data.t.left);
		release(a->data.t.right);
	}
	free(a);
}

ast *node(ast *left, ast *right) {
	ast *ret = malloc(sizeof *ret);
	assert(ret);
	ret->is_node = 1;
	ret->data.t.left = left;
	ret->data.t.right = right;
	ret->count = 1;
	return ret;
}

ast *leaf(char c) {
	ast *ret = malloc(sizeof *ret);
	assert(ret);
	ret->is_node = 0;
	ret->data.c = c;
	ret->count = 1;
	return ret;
}

void f_raw(void) {
	if (pc->here->is_node) {
		descend();
		return;
	}
	switch (pc->here->data.c) {
	case 's':
		ascend();
		todo = &f_s;
		return;
	case 'k':
		ascend();
		todo = &f_k;
		return;
	case 'i':
		ascend();
		todo = &f_i;
		return;
	default:
		while (pc)
			ascend();
		return;
	}
}

void f_s(void) {
	ascend();
	todo = &f_s1;
}

void f_s1(void) {
	ascend();
	todo = &f_s2;
}

void f_s2(void) {
	ast *left1, *left2, *x, *y, *z;
	left1 = pc->here->data.t.left;
	z = pc->here->data.t.right;
	left2 = left1->data.t.left;
	y = left1->data.t.right;
	release(left2->data.t.left); // the 's'
	x = left2->data.t.right;
	++z->count;
	replace(node(node(x, z), node(y, z)));
	free(left1);
	free(left2);
	todo = &f_raw;
}

void f_k(void) {
	ascend();
	todo = &f_k1;
}

void f_k1(void) {
	ast *left = pc->here->data.t.left;
	release(pc->here->data.t.right); // ignored parameter
	release(left->data.t.left); // the 'k'
	replace(left->data.t.right);
	free(left);
	todo = &f_raw;
}

void f_i(void) {
	release(pc->here->data.t.left); // the 'i'
	replace(pc->here->data.t.right);
	todo = &f_raw;
}

void ascend(void) {
	path *to_free = pc;
	pc = pc->up;
	free(to_free);
}

void descend(void) {
	path *newp = malloc(sizeof *newp);
	assert(newp);
	newp->up = pc;
	newp->here = pc->here->data.t.left;
	pc = newp;
}

void replace(ast *a) {
	free(pc->here);
	if (pc->up)
		pc->up->here->data.t.left = a;
	else
		program = a;
	pc->here = a;
}
