#include <iostream>
#include <string>
#include <unordered_map>


using namespace std;


std::string lab1(std::string symbol) {

	std::unordered_map<std::string, int> data;

	data["I"] = 1;
	data["V"] = 5;
	data["X"] = 10;
	data["L"] = 50;
	data["C"] = 100;
	data["D"] = 500;
	data["M"] = 1000;


	for (int i = 0; i < symbol.length(); i++) {

		if (symbol == "I") {
			cout << "1" << endl;
		}

		if (symbol == "V") {
			cout << "5" << endl;
		}

		if (symbol == "X") {
			cout << "10" << endl;
		}

		if (symbol == "L") {
			cout << "50" << endl;
		}

		if (symbol == "C") {
			cout << "100" << endl;
		}

		if (symbol == "D") {
			cout << "500" << endl;
		}

		if (symbol == "M") {
			cout << "1000" << endl;
		}
	}
	return symbol;
}


int main() {

	std::string inputValue;
	cout << "Input a roman numeral: " << endl;
	cin >> inputValue;

	lab1(inputValue);
}
