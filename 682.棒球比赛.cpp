/*
 * @lc app=leetcode.cn id=682 lang=cpp
 *
 * [682] 棒球比赛
 */

#include <iostream>
#include <vector>
#include <stack>
using namespace std;
// @lc code=start
class Solution {
public:
    int calPoints(vector<string>& ops) {
        stack<int> sta;
        int ret = 0;
        for (int idx = 0; idx < ops.size(); idx++) {
            if (ops[idx] == "+") {
                int n1 = sta.top();
                sta.pop();
                int n2 = sta.top();
                sta.push(n1);
                sta.push(n1 + n2);
                continue;
            }
            if (ops[idx] == "D") {
                sta.push(sta.top() * 2);
                continue;
            }
            if (ops[idx] == "C") {
                sta.pop();
                continue;
            }
            sta.push(strtol(ops[idx].c_str(), NULL, 10));
        }
        while(!sta.empty()){
            ret += sta.top();
            sta.pop();
        }

        return ret;
        
        
    }
};
// @lc code=end

