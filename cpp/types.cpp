#include "types.h"

using namespace std;

ostream & operator<<(ostream & os, const Piece & p) {
	return p.write(os);
}

class Char : public Piece {
public:
	explicit Char(char);
	const Func * eval() const;
protected:
	ostream & write(ostream &) const;
private:
	char c;
};

class Pair : public Piece {
public:
	Pair(const Piece *, const Piece *);
	const Func * eval() const;
protected:
	ostream & write(ostream &) const;
private:
	const Piece * x, * y;
};

class Block : public Func {
public:
	explicit Block(const Piece *);
	const Func * apply(const Piece *) const;
	const Piece * defuse() const;
private:
	const Piece * p;
};

class S : public Func {
public:
	const Func * apply(const Piece *) const;
	const Piece * defuse() const;
};

class S1 : public Func {
public:
	explicit S1(const Piece *);
	const Func * apply(const Piece *) const;
	const Piece * defuse() const;
private:
	const Piece * x;
};

class S2 : public Func {
public:
	S2(const Piece *, const Piece *);
	const Func * apply(const Piece *) const;
	const Piece * defuse() const;
private:
	const Piece * x, * y;
};

class K : public Func {
public:
	const Func * apply(const Piece *) const;
	const Piece * defuse() const;
};

class K1 : public Func {
public:
	explicit K1(const Piece *);
	const Func * apply(const Piece *) const;
	const Piece * defuse() const;
private:
	const Piece * x;
};

class I : public Func {
public:
	const Func * apply(const Piece *) const;
	const Piece * defuse() const;
};

Char::Char(char c_) : c(c_) {}

const Func * Char::eval() const {
	switch (c) {
	case 's':
		return new S;
	case 'k':
		return new K;
	case 'i':
		return new I;
	}
	return new Block(this);
}

ostream & Char::write(ostream & os) const {
	return os << c;
}

Pair::Pair(const Piece * x_, const Piece * y_) : x(x_), y(y_) {}

const Func * Pair::eval() const {
	return x->eval()->apply(y);
}

ostream & Pair::write(ostream & os) const {
	return os << '`' << *x << *y;
}

Piece * parse(istream & is) {
	char c;
	is >> c;
	while (isspace(c))
		is >> c;
	if (c != '`')
		return new Char(c);
	Piece * x = parse(is);
	Piece * y = parse(is);
	return new Pair(x, y);
}

Block::Block(const Piece * p_) : p(p_) {}

const Func * Block::apply(const Piece * other) const {
	return new Block(new Pair(p, other));
}

const Piece * Block::defuse() const {
	return p;
}

const Func * S::apply(const Piece * x) const {
	return new S1(x);
}

const Piece * S::defuse() const {
	return new Char('s');
}

S1::S1(const Piece * x_) : x(x_) {}

const Func * S1::apply(const Piece * y) const {
	return new S2(x, y);
}

const Piece * S1::defuse() const {
	return new Pair(new Char('s'), x);
}

S2::S2(const Piece * x_, const Piece * y_) : x(x_), y(y_) {}

const Func * S2::apply(const Piece * z) const {
	return x->eval()->apply(z)->apply(new Pair(y, z));
}

const Piece * S2::defuse() const {
	return new Pair(new Pair(new Char('s'), x), y);
}

const Func * K::apply(const Piece * x) const {
	return new K1(x);
}

const Piece * K::defuse() const {
	return new Char('k');
}

K1::K1(const Piece * x_) : x(x_) {}

const Func * K1::apply(const Piece *) const {
	return x->eval();
}

const Piece * K1::defuse() const {
	return new Pair(new Char('k'), x);
}

const Func * I::apply(const Piece * x) const {
	return x->eval();
}

const Piece * I::defuse() const {
	return new Char('i');
}
