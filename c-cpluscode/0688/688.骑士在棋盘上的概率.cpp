/*
 * @Author       : zc
 * @Date         : 2022-02-17 22:55:40
 * @LastEditors  : zc
 * @LastEditTime : 2022-02-19 21:04:54
 * @Description  : file content
 * @FilePath     : \leetcode\0688\688.骑士在棋盘上的概率.cpp
 */
/*
 * @lc app=leetcode.cn id=688 lang=cpp
 *
 * [688] 骑士在棋盘上的概率
 */
#include <iostream>
#include <list>

using namespace std;

// @lc code=start
class Solution {
    using GRID=pair<int,int>;
    using GRID_LIST=list<GRID>;
public:
    int cnt = 0;
    int total = 1;
public:
    double knightProbability(int n, int k, int row, int column) {
        if (row < 0 || row >= n || column < 0 || column >= n) {
            return 0.0;
        }

        if (k==0) {
            return 1.0;
        }

        dsf(n,k,row,column);
        double ret = static_cast<double>(cnt);

        while (k>0) {
            ret = ret / 8.0;
            k--;
        }

        return ret;
    }

    void dsf(int n, int k, int row, int col) {
        if (row < 0 || row >= n || col < 0 || col >= n) {
            return;
        }

        if (k <= 0) {
            cnt++;
            return;
        }

        k--;

        dsf(n, k, row-1, col-2);
        dsf(n, k, row-2, col-1);
        dsf(n, k, row-2, col+1);
        dsf(n, k, row-1, col+2);
        dsf(n, k, row+1, col+2);
        dsf(n, k, row+2, col+1);
        dsf(n, k, row+2, col-1);
        dsf(n, k, row+1, col-2);
    }
};

int main(void) {
    Solution s;

    int n, k, row, col;

    cin >> n >> k >> row >> col;

    cout << s.knightProbability(n,k,row,col) << endl;

    cout << s.total << " " << s.cnt << endl;
    return 0;
}
// @lc code=end

