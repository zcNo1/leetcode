/*
 * @Author       : zc
 * @Date         : 2021-06-29 21:40:53
 * @LastEditors  : zc
 * @LastEditTime : 2021-06-29 22:48:57
 * @Description  : file content
 * @FilePath     : \leetcode\168.excel表列名称.cpp
 */
/*
 * @lc app=leetcode.cn id=168 lang=cpp
 *
 * [168] Excel表列名称
 */
#include <iostream>
#include <vector>
#include <string>
using namespace std;
// @lc code=start
class Solution {
public:
    string convertToTitle(int columnNumber) {
        int quotient = columnNumber;
        int remainder = 0;
        string str;
        while (quotient != 26) {
            quotient = quotient / 26;
            remainder = quotient % 26;
            if (remainder != 0){
                str.push_back('A' + remainder - 1);
            }
        }
        return string(str.rbegin(),str.rend());
    }
};
// @lc code=end

