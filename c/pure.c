#include "pure.h"

#include <stdlib.h>
#include <stdio.h>
#include <ctype.h>

int main(int argc, char *argv[]) {
	AST *a = parse(stdin);
	if (a == NULL) {
		fprintf(stderr, "Syntax error.");
		return 1;
	}
	fprint(stdout, freeze(eval(a)));
	printf("\n");
	return 0;
}

AST *parse(FILE *f) {
	char c;
	while (isspace(c = fgetc(f)));
	if (c == EOF)
		return NULL;
	if (c == '`') {
		AST *left = parse(f);
		if (left == NULL) return NULL;
		AST *right = parse(f);
		if (left == NULL) return NULL;
		return combine(left, right);
	}
	return from_char(c);
}

void fprint(FILE *f, AST *a) {
	switch (a->type) {
	case CHAR:
		fputc(a->val.c, f);
		break;
	case PAIR:
		fputc('`', f);
		fprint(f, a->val.pair.left);
		fprint(f, a->val.pair.right);
		break;
	default:
		fprintf(stderr, "This should never happen!");
	}
}

AST *from_char(char c) {
	AST *ret = malloc(sizeof(AST));
	ret->type = CHAR;
	ret->val.c = c;
	return ret;
}

AST *combine(AST *left, AST *right) {
	AST *ret = malloc(sizeof(AST));
	ret->type = PAIR;
	ret->val.pair.left = left;
	ret->val.pair.right = right;
	return ret;
}

Func eval(AST *a) {
	Func ret;
	switch (a->type) {
	case CHAR:
		switch (a->val.c) {
		case 's':
			ret.apply = &apply_s;
			break;
		case 'k':
			ret.apply = &apply_k;
			break;
		case 'i':
			ret.apply = &apply_i;
			break;
		default:
			ret.apply = &apply_block;
			ret.data.single = a;
			break;
		}
		break;
	case PAIR:
		ret = apply(eval(a->val.pair.left), a->val.pair.right);
		break;
	}
	return ret;
}

AST *freeze(Func f) {
	if (f.apply == &apply_block)
		return f.data.single;
	if (f.apply == &apply_s)
		return from_char('s');
	if (f.apply == &apply_s1)
		return combine(from_char('s'), f.data.single);
	if (f.apply == &apply_s2) {
		return combine(combine(from_char('s'), f.data.pair.left), f.data.pair.right);
	}
	if (f.apply == &apply_k)
		return from_char('k');
	if (f.apply == &apply_k1)
		return combine(from_char('k'), f.data.single);
	if (f.apply == &apply_i)
		return from_char('i');
	//CAN'T HAPPEN
	return NULL;
}

Func apply(Func left, AST *right) {
	return left.apply(left.data, right);
}

Func apply_block(FData data, AST *right) {
	Func ret = {combine(data.single, right), &apply_block};
	return ret;
}

Func apply_s(FData _, AST *x) {
	Func ret = {x, &apply_s1};
	return ret;
}

Func apply_s1(FData data, AST *y) {
	data.pair.left = data.single;
	data.pair.right = y;
	Func ret = {data, &apply_s2};
	return ret;
}

Func apply_s2(FData data, AST *z) {
	return apply(apply(eval(data.pair.left), z), combine(data.pair.right, z));
}

Func apply_k(FData _, AST *x) {
	FData data;
	data.single = x;
	Func ret = {data, &apply_k1};
	return ret;
}

Func apply_k1(FData data, AST *_) {
	return eval(data.single);
}

Func apply_i(FData _, AST *x) {
	return eval(x);
}
