#include <iostream>
using namespace std;

int main() {
	int num;
	cin >> num;
	int now = 1, max = 2, pre = 0;
	bool flag = true;
	for (int i = 0; i <  num;) {
		for (int j = 0; j < now; j++) {
			i++;
			cout << i << " ";
			if (i == num)break;
			
			 
			 
		}
	 
		cout << "\n";
		if (now-pre==1 && max-now > 0 && flag) {
			pre = now;
			now++;
		}
		else if(now-pre == 1 && max==now){
			max++;
			pre = now; 
			now--;
			flag = false;
		}
		else if (flag == false && pre - now > 0) {
			pre = now;
			now--;
			if (now == 0) {
				now = pre + 1;
				pre = 1;
				flag = true;
			}

		}
	 

	}
}
