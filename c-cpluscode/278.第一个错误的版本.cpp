/*
 * @Author       : zc
 * @Date         : 2021-06-13 11:27:33
 * @LastEditors  : zc
 * @LastEditTime : 2021-06-13 11:55:17
 * @Description  : file content
 * @FilePath     : \leetcode\278.第一个错误的版本.cpp
 */
/*
 * @lc app=leetcode.cn id=278 lang=cpp
 *
 * [278] 第一个错误的版本
 */

// @lc code=start
// The API isBadVersion is defined for you.
// bool isBadVersion(int version);

class Solution {
public:
    int firstBadVersion(int n) {
        if (n <= 1) {
            return n;
        }
        int l = 1, r = n;
        // 为了避免整数溢出，必须这样计算二分法的mid值
        int m = l + (r - l) / 2;
        while (l < r) {
            if (isBadVersion(m)) {
                r = m;
            } else {
                l = m + 1;
            }
            // 为了避免整数溢出，必须这样计算二分法的mid值
            m = l + (r - l) / 2;
        }
        return l;
    }
};
// @lc code=end

