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
	AST *ret = (AST *) malloc(sizeof(AST));
	ret->type = CHAR;
	ret->val.c = c;
	return ret;
}

AST *combine(AST *left, AST *right) {
	AST *ret = (AST *) malloc(sizeof(AST));
	ret->type = PAIR;
	ret->val.pair.left = left;
	ret->val.pair.right = right;
	return ret;
}

Func *eval(AST *a) {
	Func *ret;
	switch (a->type) {
	case CHAR:
		ret = (Func *) malloc(sizeof(Func));
		switch (a->val.c) {
		case 's':
			ret->type = S;
			break;
		case 'k':
			ret->type = K;
			break;
		case 'i':
			ret->type = I;
			break;
		default:
			ret->type = BLOCK;
			ret->val.block = a;
			break;
		}
		break;
	case PAIR:
		ret = apply(eval(a->val.pair.left), a->val.pair.right);
		break;
	}
	return ret;
}

AST *freeze(Func *f) {
	switch (f->type) {
	case BLOCK:
		return f->val.block;
	case S:
		return from_char('s');
	case S1:
		return combine(from_char('s'), f->val.s1.x);
	case S2:
		return combine(combine(from_char('s'), f->val.s2.x), f->val.s2.y);
	case K:
		return from_char('k');
	case K1:
		return combine(from_char('k'), f->val.k1.x);
	case I:
		return from_char('i');
	}
	//CAN'T HAPPEN
	return NULL;
}

Func *apply(Func *left, AST *right) {
	Func *ret;
	switch (left->type) {
	case BLOCK:
		ret = (Func *) malloc(sizeof(Func));
		ret->type = BLOCK;
		ret->val.block = combine(left->val.block, right);
		break;
	case S:
		ret = (Func *) malloc(sizeof(Func));
		ret->type = S1;
		ret->val.s1.x = right;
		break;
	case S1:
		ret = (Func *) malloc(sizeof(Func));
		ret->type = S2;
		ret->val.s2.x = left->val.s1.x;
		ret->val.s2.y = right;
		break;
	case S2:
		ret = apply(apply(eval(left->val.s2.x), right), combine(left->val.s2.y, right));
		break;
	case K:
		ret = (Func *) malloc(sizeof(Func));
		ret->type = K1;
		ret->val.k1.x = right;
		break;
	case K1:
		ret = eval(left->val.k1.x);
		break;
	case I:
		ret = eval(right);
		break;
	}
	return ret;
}