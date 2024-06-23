/*
 * @Author       : zc
 * @Date         : 2021-06-22 23:26:37
 * @LastEditors  : zc
 * @LastEditTime : 2021-06-23 00:02:03
 * @Description  : file content
 * @FilePath     : \leetcode\剑指Offer38.字符串的排列.cpp
 */
#include<vector>
#include<iostream>
#include<algorithm>
using namespace std;
class Solution {
private:
    vector<string> ret;
    int sz;
    void traverseS(string rets, string s, int idx ) {
        for (int i = idx; i < s.size(); i++) {
            if (i > idx && s[i] == s[i-1]) {
                continue;
            }
            rets.push_back(s[i]);
            if (rets.size() == sz) {
                ret.push_back(rets);
            }
            string s1 = s;
            s1.erase(i, 1);
            traverseS(rets, s1, idx);
            rets.pop_back();
        }
    }
public:
    vector<string> permutation(string s) {
        sort(s.begin(), s.end());
        sz = s.size();
        string rets;
        traverseS(rets, s, 0);
        return ret;
    }
};