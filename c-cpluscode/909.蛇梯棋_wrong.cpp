/*
 * @Author       : zc
 * @Date         : 2021-06-27 14:44:37
 * @LastEditors  : zc
 * @LastEditTime : 2021-06-28 22:50:20
 * @Description  : file content
 * @FilePath     : \leetcode\909.蛇梯棋.cpp
 */
/*
 * @lc app=leetcode.cn id=909 lang=cpp
 *
 * [909] 蛇梯棋
 */
// 动态规划法
// 位置 loc,
// x = board[0].length;
// y = board.length;
// i_loc = y - loc / x;
// j_loc = loc % x;
// pre_n = loc - n // n = 1,2,3,4,5,6
// if (board[i_pre_n][j_pre_n] == -1)
//      dp[i_loc][j_loc] = dp[i_pre_n][j_pre_n] + n;
// else 
//      continue;
// if (board[i_loc][j_loc] != -1)
//      next = board[i_loc][j_loc];
//      dp[i_next][j_next] = dp[i_next][j_next] == -1 
//          ? dp[i_loc][j_loc] 
//          : min(dp[i_next][j_next], dp[i_loc][j_loc]);
#include <iostream>
#include <vector>
using namespace std;
// @lc code=start
class Solution {
private:
    vector<vector<int>> val_rc;
    void get_val_rc(vector<vector<int>>& val) {
        int col = val[0].size();
        int row = val.size();
        int sz = col * row;
        val_rc = vector<vector<int>>(2, vector<int>(sz, -1));
        int idx = row - 1, idy = 0;
        bool is_right = true;
        for (int loc = 1; loc <= sz; loc++) {
            val_rc[0][loc - 1] = idx;
            val_rc[1][loc - 1] = idy;
            if (idy == 0 && !is_right) {
                is_right = true;
                idx--;
                continue;
            } else if (idy == col - 1 && is_right) {
                is_right = false;
                idx--;
                continue;
            }
            if (is_right) {
                idy++;
            } else {
                idy--;
            }
        }
        return;
    }
    int get_dp(vector<vector<int>>& val, int loc) {
        if (loc <= 0) {
            return 0;
        }
        return get_val(val, loc);
    }
    int get_board(vector<vector<int>>& val, int loc) {
        if (loc <= 0) {
            return -1;
        }
        return get_val(val, loc);
    }
    int get_val(vector<vector<int>>& val, int loc) {
        if (val_rc[0][loc - 1] != -1 && val_rc[1][loc - 1] != -1) {
            return val[val_rc[0][loc - 1]][val_rc[1][loc - 1]];
        }
        return -1;
    }
    void set_val(vector<vector<int>>& val, int loc, int num) {
        if (val_rc[0][loc - 1] != -1 && val_rc[1][loc - 1] != -1) {
            int &v = val[val_rc[0][loc - 1]][val_rc[1][loc - 1]];
            if (v != -1) {
                v = min(num, v);
            } else {
                v = num;
            }   
        }
        return;
    }
public:
    int snakesAndLadders(vector<vector<int>>& board) {
        int col = board[0].size();
        int row = board.size();
        int sz = col * row;

        vector<vector<int>> dp(row, vector<int>(col, -1));
        get_val_rc(board);
        int next_board = get_board(board, 1);
        if (next_board != -1) {
            set_val(dp, next_board, 0);
        } else {
            set_val(dp, 1, 0);
        }
        for (int loc = 2; loc <= sz; loc++) {
            // 设置最小值的初始值为sz+1步
            int min_n = sz;
            // 计算前六格的值哪一个可以跳到该点
            // 取其中的最小值 + 1 作为该点的步数
            for (int n = 1; n <= 6; n++) {
                int pre_dp = get_dp(dp, loc - n);
                // 如果该点前面的某点真实步数为-1，证明无法从某点跳到该点，不列入统计
                if (pre_dp != -1) {
                    min_n = min(min_n, pre_dp);
                }
            }
            // 如果最小值不为初始值，证明该点可以有真实步数
            // 如果该点可以跳到下一个点，则该点步数为-1，下一个点为真实步数
            // 否则，这个点为真实步数
            if (min_n != sz) {
                int step = min_n + 1;
                int next_board = get_board(board, loc);
                if (next_board != -1) {
                    set_val(dp, next_board, step);
                } else {
                    int loc_dp = get_dp(dp, loc);
                    if (loc_dp != -1) {
                        step = min(step, loc_dp);
                    }
                    set_val(dp, loc, step);
                }
            }
            cout << loc << " " << get_board(board, loc) << " " <<\
             get_dp(dp, loc) << " " << endl;
        }
        cout << 1 << " " << get_board(board, 1) << " " <<\
             get_dp(dp, 1) << " " << endl;

        return get_dp(dp, sz);
    }
};
// @lc code=end

[   [   -1, -1, -1, -1, 48, 5,  -1  ],
    [   12, 29, 13, 9,  -1, 2,  32  ],
    [   -1, -1, 21, 7,  -1, 12, 49  ],
    [   42, 37, 21, 40, -1, 22, 12  ],
    [   42, -1, 2,  -1, -1, -1, 6   ],
    [   39, -1, 35, -1, -1, 39, -1  ],
    [   -1, 36, -1, -1, -1, -1, 5   ]]