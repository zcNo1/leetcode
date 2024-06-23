/*
 * @Author       : zc
 * @Date         : 2022-02-13 10:35:23
 * @LastEditors  : zc
 * @LastEditTime : 2022-02-13 10:41:44
 * @Description  : file content
 * @FilePath     : \leetcode\1189.气球-的最大数量.cpp
 */
/*
 * @lc app=leetcode.cn id=1189 lang=cpp
 *
 * [1189] “气球” 的最大数量
 */

#include <iostream>
#include <vector>
#include <string>

using namespace std;

// @lc code=start
class Solution {
public:
    int maxNumberOfBalloons(string text) {
        vector<int> record(5,0);

        for (int i = 0; i < text.size(); i++) {
            switch (text.at(i))
            {
            case 'a':
                record[0]++;
                break;
            case 'b':
                record[1]++;
                break;
            case 'n':
                record[2]++;
                break;
            case 'o':
                record[3]++;
                break;
            case 'l':
                record[4]++;
                break;
            
            default:
                break;
            }
        }

        int ret = record[0];
        for (int i = 1; i < 3; i++) {
            if (ret > record.at(i)) {
                ret = record.at(i);
            }
        }
        for (int i = 3; i < 4; i++) {
            if (ret > record.at(i) / 2) {
                ret = record.at(i) / 2;
            }
        }

        return ret;
    }
};
// @lc code=end

