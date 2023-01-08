/*
 * @Author       : zc
 * @Date         : 2022-02-12 22:55:47
 * @LastEditors  : zc
 * @LastEditTime : 2022-02-13 10:20:06
 * @Description  : file content
 * @FilePath     : \leetcode\1447.最简分数.cpp
 */
/*
 * @lc app=leetcode.cn id=1447 lang=cpp
 *
 * [1447] 最简分数
 */
#include <iostream>
#include <vector>
#include <string>

using namespace std;

#define LEN 1024

// @lc code=start
class Solution {
public:
    vector<string> simplifiedFractions(int n) {
        vector<string> ret;

        if (n <= 1) {
            return ret;
        }
        for (int idx = 2; idx < n; idx++) {
            int mid = (idx + 1) / 2;
            for (int idy = 1; idy < mid; idy++) {
                
            }
        }
    }

    string getString(int m, int n)
    {
        char str[LEN + 1] = { 0 };

        snprintf(str, LEN, "%d/%d", m, n);
        return str;
    }
};
// @lc code=end

