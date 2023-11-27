#include <iostream>
#include <random>
class Cat {
public:
	int flag = rand();
	bool is_alive() {
		return static_cast<bool>(flag%2);
	}
};
class Box {
public:
	Cat open() {
		Cat cat;
		return cat;
	}

};
