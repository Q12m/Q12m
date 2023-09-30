#include <iostream>
#include <cmath>
using namespace std;

int main()
{
    double pi = sqrt(12) * (1 - 1 / 9.0 + 1 / (5 * pow(3, 2)) - 1 / (7 * pow(3, 3)) + 1 / (9 * pow(3, 4)) - 1 / (11 * pow(3, 5)));
    cout << pi;
    return 0;
}
