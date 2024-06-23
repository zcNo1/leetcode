/*
 * @Author       : zc
 * @Date         : 2021-04-09 22:28:25
 * @LastEditors  : zc
 * @LastEditTime : 2021-04-09 23:29:58
 * @Description  : file content
 * @FilePath     : \leetcode\0154\C++\findMin.cpp
 */
#include <iostream>
#include <vector>
using namespace std;
class Solution {
public:
    int findMin(vector<int>& nums) {
        int l = 0, r = nums.size() - 1, m = 0;
        while (l < r) {
            m = (l + r) / 2;
            if (nums[m] < nums[r]) {
                r = m;
            } else if (nums[m] > nums[r]) {
                l = m + 1;
            } else {
                if (nums[r] < nums[l]) {
                    r = m;
                } else if (nums[r] == nums[l]) {
                    r--;
                    l++;
                } else {
                    break;
                }
            }
        }
        return nums[l];
    }
};
