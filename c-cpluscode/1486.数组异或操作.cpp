/*
 * @Author       : zc
 * @Date         : 2021-05-07 20:10:17
 * @LastEditors  : zc
 * @LastEditTime : 2021-05-07 20:14:21
 * @Description  : file content
 * @FilePath     : \leetcode\1486.数组异或操作.cpp
 */
/*
 * @lc app=leetcode.cn id=1486 lang=cpp
 *
 * [1486] 数组异或操作
 */

// @lc code=start
class Solution {
public:
    int xorOperation(int n, int start) {
        int idx = 1, ret = start;
        while(idx < n) {
            ret ^= start + 2 * idx;
            idx++;
        }
        return ret;
    }
};
// @lc code=end

