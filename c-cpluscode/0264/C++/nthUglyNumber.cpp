/*
 * @Author       : zc
 * @Date         : 2021-04-11 20:50:28
 * @LastEditors  : zc
 * @LastEditTime : 2021-04-12 00:05:17
 * @Description  : file content
 * @FilePath     : \leetcode\0264\C++\nthUglyNumber.cpp
 */
#include <iostream>
#include <vector>

using namespace std;
class Solution {
public:
    int nthUglyNumber(int n) {
        vector<int> dp;
        dp.push_back(1);
        n--;
        int p2 = 0, p3 = 0, p5 = 0;
        int n2, n3, n5, dpn = 1;
        while (n) {
            n2 = dp[p2] * 2;
            n3 = dp[p3] * 3;
            n5 = dp[p5] * 5;
            dpn = min(min(n2, n3), n5);
            dp.push_back(dpn);
            if (dpn == n2) {
                p2++;
            }
            if (dpn == n3) {
                p3++;
            }
            if (dpn == n5) {
                p5++;
            }
            n--;
        }
        return dpn;
    }
};
    int nthUglyNumber(int n) {
        vector<int> dp;
        dp.push_back(1);
        n--;
        int p2 = 0, p3 = 0, p5 = 0;
        int n2, n3, n5, dpn = 1;
        while (n) {
            n2 = dp[p2] * 2;
            n3 = dp[p3] * 3;
            n5 = dp[p5] * 5;
            dpn = min(min(n2, n3), n5);
            cout << "n2: " << n2 << "\tn3: " << n3 << "\tn5: " << n5 << "\tdpn: " << dpn << endl;
            dp.push_back(dpn);
            if (dpn == n2) {
                p2++;
            }
            if (dpn == n3) {
                p3++;
            }
            if (dpn == n5) {
                p5++;
            }
            n--;
        }
        return dpn;
    }
int main(void) {
    int n;
    while (cin >> n)
    cout << nthUglyNumber(n) << endl;
}