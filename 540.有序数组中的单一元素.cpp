/*
 * @Author       : zc
 * @Date         : 2022-02-14 20:41:42
 * @LastEditors  : zc
 * @LastEditTime : 2022-02-14 21:01:40
 * @Description  : file content
 * @FilePath     : \leetcode\540.有序数组中的单一元素.cpp
 */
/*
 * @lc app=leetcode.cn id=540 lang=cpp
 *
 * [540] 有序数组中的单一元素
 */

#include <iostream>
#include <vector>

using namespace std;

// @lc code=start
class Solution {
public:
    int singleNonDuplicate(vector<int>& nums) {
        int low = 0, high = nums.size() - 1;
        while (low < high) {
            int mid = (high - low) / 2 + low;
            if (nums[mid] == nums[mid ^ 1]) {
                low = mid + 1;
            } else {
                high = mid;
            }
        }
        return nums[low];
    }
};
// @lc code=end

