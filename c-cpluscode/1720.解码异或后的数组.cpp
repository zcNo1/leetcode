/*
 * @Author       : zc
 * @Date         : 2021-05-06 23:17:12
 * @LastEditors  : zc
 * @LastEditTime : 2021-05-06 23:29:14
 * @Description  : file content
 * @FilePath     : \leetcode\1720.解码异或后的数组.cpp
 */
/*
 * @lc app=leetcode.cn id=1720 lang=cpp
 *
 * [1720] 解码异或后的数组
 */
#include <0000/C++/inc_C++.h>
// @lc code=start
class Solution {
public:
    vector<int> decode(vector<int>& encoded, int first) {
        int sz = encoded.size();
        vector<int> ret(sz + 1);
        ret[0] = first;
        for (int idx = 0; idx < sz; idx++) {
            ret[idx + 1] = ret[idx] ^ encoded[idx];
        }

        return ret;
    }
};
// @lc code=end

