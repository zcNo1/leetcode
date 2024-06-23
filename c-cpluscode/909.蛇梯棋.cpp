/*
 * @Author       : zc
 * @Date         : 2021-06-27 14:44:37
 * @LastEditors  : zc
 * @LastEditTime : 2021-06-29 00:01:20
 * @Description  : file content
 * @FilePath     : \leetcode\909.蛇梯棋.cpp
 */
/*
 * @lc app=leetcode.cn id=909 lang=cpp
 *
 * [909] 蛇梯棋
 */
#include <iostream>
#include <vector>
#include <queue>
using namespace std;
typedef struct {
    int loc;
    int step;
}LOC;
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
        int end = col * row;
        int ret_step = -1;

        queue<LOC> location;
        vector<bool> vis(end + 1, false);
        get_val_rc(board);
        LOC loc = {1, 0};
        location.push(loc);
        while (!location.empty()) {
            LOC cur = location.front();
            location.pop();
            for (int n = 1; n <= 6; n++) {
                int next_loc = cur.loc + n;
                if(next_loc > end) {
                    goto CLEAN;
                }
                int step = cur.step + 1;
                int next_next = get_board(board, next_loc);
                if (next_next != -1) {
                    next_loc = next_next;
                }
                if (next_loc == end) {
                    ret_step = step;
                    goto CLEAN;
                }
                if (!vis[next_loc]) {
                    LOC next = {next_loc, step};
                    location.push(next);
                    vis[next_loc] = true;
                }
                
            }
        }
CLEAN:
        return ret_step;
        
    }
};
// @lc code=end