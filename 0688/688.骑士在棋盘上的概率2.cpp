/*
 * @Author       : zc
 * @Date         : 2022-02-17 22:55:40
 * @LastEditors  : zc
 * @LastEditTime : 2022-02-20 14:17:27
 * @Description  : file content
 * @FilePath     : \leetcode\0688\688.骑士在棋盘上的概率2.cpp
 */
/*
 * @lc app=leetcode.cn id=688 lang=cpp
 *
 * [688] 骑士在棋盘上的概率
 */
#include <iostream>
#include <list>
#include <vector>

using namespace std;

// @lc code=start
class Solution {
public:
    double knightProbability(int n, int k, int row, int column) {
        if (row < 0 || row >= n || column < 0 || column >= n) {
            return 0.0;
        }

        if (k==0) {
            return 1.0;
        }
        int ret;
        vector<vector<vector<double>>> dp(k+1, vector<vector<double>>(n, vector<double>(n, 1.0)));
        vector<vector<int>> dirs = {{-2, -1}, {-2, 1}, {2, -1}, {2, 1}, {-1, -2}, {-1, 2}, {1, -2}, {1, 2}};

        for (int kk = 1; kk <= k; kk++) {
            for (int xx = 0; xx < n; xx++) {
                for (int yy = 0; yy < n; yy++) {
                    dp[kk][xx][yy] = 0.0;
                    for (auto & dir : dirs) {
                        int ni = xx + dir[0], nj = yy + dir[1];
                        if (ni >= 0 && ni < n && nj >= 0 && nj < n) {
                            dp[kk][xx][yy] += dp[kk-1][ni][nj] / 8.0;
                        }
                    }
                }
            }
        }
        
        return dp[k][row][column];
    }
};

int main(void) {
    Solution s;

    int n, k, row, col;

    cin >> n >> k >> row >> col;

    cout << s.knightProbability(n,k,row,col) << endl;
    return 0;
}
// @lc code=end

