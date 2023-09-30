#include <iostream>
#include <math.h>
using namespace std;

int main()
{
    int a,b,c;
    cin >> a >> b >> c;
    int res1{ abs(a - b) }, res2{ abs(a - c) };
    if (res1 < res2) {
        cout << "B "<< res1;
    }
    else {
        cout << "C " << res2;
    }
}
