#include <iostream>
#include "types.h"

using namespace std;

int main() {
	Piece * p = parse(cin);
	cout << *p->eval()->defuse() << endl;
	delete p;
}
