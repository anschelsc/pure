#ifndef PURE_TYPES
#define PURE_TYPES

#include <ostream>
#include <istream>
#include <string>

class Func;

class Piece {
public:
	virtual ~Piece() {}
	virtual const Func * eval() const = 0;

	friend std::ostream & operator<<(std::ostream &, const Piece &);
protected:
	virtual std::ostream & write(std::ostream &) const = 0;
};

Piece * parse(std::istream &);

class Func {
public:
	virtual const Func * apply(const Piece *) const = 0;
	virtual const Piece * defuse() const = 0;
};

#endif
