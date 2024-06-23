/*
 * @Author       : zc
 * @Date         : 2021-06-14 11:04:05
 * @LastEditors  : zc
 * @LastEditTime : 2021-06-14 11:10:01
 * @Description  : file content
 * @FilePath     : \leetcode\374.猜数字大小.cpp
 */
/*
 * @lc app=leetcode.cn id=374 lang=cpp
 *
 * [374] 猜数字大小
 */

// @lc code=start
/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * int guess(int num);
 */

class Solution {
public:
    int guessNumber(int n) {
        int l = 1, r = n, m = l + (r - l) / 2;
        int ret = -1;
        while(l < r) {
            ret = guess(m);
            if (ret == 0) {
                return m;
            } else if (ret < 0) {
                r = m;
            } else {
                l = m + 1;
            }
            m = l + (r - l) / 2;
        }
        return l;
    }
};
// @lc code=end

