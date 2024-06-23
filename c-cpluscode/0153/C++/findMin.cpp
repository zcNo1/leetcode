/*
 * @Author       : zc
 * @Date         : 2021-04-08 23:46:09
 * @LastEditors  : zc
 * @LastEditTime : 2021-04-09 00:04:14
 * @Description  : 0153
 * @FilePath     : \leetcode\0153\C++\findMin.cpp
 */

class Solution {
public:
    int findMin(vector<int>& nums) {
        int l = 0, r = nums.size() - 1;
        int mid = 0;
        while (l < r) {
            mid = (r + l) / 2;
            if (nums[mid] > nums[r]) {
                l = mid + 1;
            } else {
                r = mid;
            }

        }
        
        return nums[l];
    }
};