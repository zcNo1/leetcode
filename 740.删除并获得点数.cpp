/*
 * @Author       : zc
 * @Date         : 2021-05-07 20:19:12
 * @LastEditors  : zc
 * @LastEditTime : 2021-05-07 20:59:02
 * @Description  : file content
 * @FilePath     : \leetcode\740.删除并获得点数.cpp
 */
/*
 * @lc app=leetcode.cn id=740 lang=cpp
 *
 * [740] 删除并获得点数
 */

// @lc code=start
// #include <0000/C++/inc_C++.h>
class Solution {
public:
    int deleteAndEarn(vector<int>& nums) {
        sort(nums.begin(), nums.end());
        int sz = nums.size();
        vector<int> dp(sz,0);
        int ret = nums[0];
        int pre = 0;
        dp[0] = nums[0];
        for (int idx = 1; idx < sz; idx++) {
            if (nums[idx] == nums[idx-1]){
                dp[idx] = dp[idx-1] + nums[idx];
            } else if (nums[idx] == nums[idx-1] + 1){
                dp[idx] = pre + nums[idx];
            } else {
                dp[idx] = max(pre, dp[idx-1]) + nums[idx];
            } 
            if (nums[idx] != nums[idx-1]) {
                pre = ret;
            }         
            ret = max(ret, dp[idx]);
        }
        return ret;

    }
};
// @lc code=end

