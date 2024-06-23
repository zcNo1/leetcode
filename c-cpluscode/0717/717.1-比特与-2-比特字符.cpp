/*
 * @Author       : zc
 * @Date         : 2022-02-20 14:19:30
 * @LastEditors  : zc
 * @LastEditTime : 2022-02-20 18:07:38
 * @Description  : file content
 * @FilePath     : \leetcode\0717\717.1-比特与-2-比特字符.cpp
 */
/*
 * @lc app=leetcode.cn id=717 lang=cpp
 *
 * [717] 1比特与2比特字符
 */

#include <iostream>
#include <vector>

using namespace std;
// @lc code=start
class Solution {
public:
    bool isOneBitCharacter(vector<int>& bits) {
        vector<int>::const_iterator it = bits.end() - 1;

        if (*it == 1) {
            return false;
        }
        if (bits.size() == 1) {
            return true;
        }
        it--;
        int cnt1 = 0;

        for (;it >= bits.begin();--it) {
            if (*it == 0) {
                break;
            } else {
                cnt1++;
            }
        }

        if (cnt1 / 2==0){
            return true;
        } else {
            return false;
        }
    }
};

int main(void) {
    Solution s;

    return 0;
}
// @lc code=end

