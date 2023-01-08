/*
 * @lc app=leetcode.cn id=2028 lang=cpp
 *
 * [2028] 找出缺失的观测数据
 */
#include <iostream>
#include <vector>
#include <stack>
using namespace std;
// @lc code=start
class Solution {
public:
    vector<int> missingRolls(vector<int>& rolls, int mean, int n) {
        int m = rolls.size();
        int na = (m + n) * mean;
        int nc = 0;

        for (int i = 0; i < m; i++) {
            nc += rolls[i];
        }
        if (nc + n > na || nc + 6 * n < na) {
            return vector<int>();
        }
        vector<int> ret(n, 1);
        int ns = na - n - nc;

        int cnt = 0;
        while (ns > 5) {
            ret[cnt] = 6;
            ns -= 5;
            cnt++;
        }

        ret[cnt] += ns;

        return ret;

    }
};
// @lc code=end

