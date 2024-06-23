/*
 * @lc app=leetcode.cn id=693 lang=cpp
 *
 * [693] 交替位二进制数
 */

// @lc code=start
class Solution {
public:
    bool hasAlternatingBits(int n) {
        int start = 1;
        if (n%2==0) {
            start = 2;
        }
        int s = 0;
        while (n > s) {
            s += start;
            start *= 4;
        }
        if (n ==s) {
            return true;
        }
        return false;
    }
};
// @lc code=end

