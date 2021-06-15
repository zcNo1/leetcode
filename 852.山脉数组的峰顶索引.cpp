/*
 * @Author       : zc
 * @Date         : 2021-06-15 23:32:46
 * @LastEditors  : zc
 * @LastEditTime : 2021-06-15 23:51:32
 * @Description  : file content
 * @FilePath     : \leetcode\852.山脉数组的峰顶索引.cpp
 */
/*
 * @lc app=leetcode.cn id=852 lang=cpp
 *
 * [852] 山脉数组的峰顶索引
 */

// @lc code=start
#include <iostream>
#include <vector>
using namespace std;
class Solution {
public:
    int peakIndexInMountainArray(vector<int>& arr) {
        int l = 0, r = arr.size() - 1, m = l + (r - l) / 2;
        while (l < r) {
            if (arr[m] < arr[m + 1]) {
                l = m + 1;
            } else if (arr[m] > arr[m + 1]) {
                r = m;
            }
            m = l + (r - l) / 2;
        }
        return l;
    }
};
// @lc code=end

