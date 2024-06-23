/*
 * @Author       : zc
 * @Date         : 2022-02-12 22:18:40
 * @LastEditors  : zc
 * @LastEditTime : 2022-02-12 22:55:04
 * @Description  : file content
 * @FilePath     : \leetcode\1020.飞地的数量.cpp
 */
/*
 * @lc app=leetcode.cn id=1020 lang=cpp
 *
 * [1020] 飞地的数量
 */

// @lc code=start

#include <vector>
#include <iostream>
using namespace std;
#define LAND 1
#define ENCLAVES 1
class Solution {
public:
    int numEnclaves(vector<vector<int>>& grid) {
        if (grid.size() == 0 || grid[0].size() == 0) {
            printf("param wrong");
            return 0;
        }
        if (grid.size() == 1 || grid[0].size() == 1) {
            return 0;
        }

        int idx = 0, idy = 0;
        int m = grid.size(), n = grid[0].size();
        vector<vector<int>> record(m, vector<int>(n, 0));

        // 遍历上边界
        for (idx = 0, idy = 0; idy < n; ++idy) {
            transferEnclaves(grid, idx, idy, record);
        }

        // 遍历下边界
        for (idx = m - 1, idy = 0; idy < n; ++idy) {
            transferEnclaves(grid, idx, idy, record);
        }

        // 遍历左边界
        for (idx = 0, idy = 0; idx < m; ++idx) {
            transferEnclaves(grid, idx, idy, record);
        }

        // 遍历右边界
        for (idx = 0, idy = n - 1; idx < m; ++idx) {
            transferEnclaves(grid, idx, idy, record);
        }

        int nLand = getNumGrid(grid, LAND);
        int nEnclaves = getNumGrid(record, ENCLAVES);
        int ret = nLand - nEnclaves;

        return ret < 0 ? 0 : ret;
    }

private:
    // 从给定的点遍历所有可以去的陆地，并记录下来
    void transferEnclaves(const vector<vector<int>>& grid,
                        size_t idx, size_t idy,
                        vector<vector<int>>& record) {
        //如果不是陆地或者这块陆地已经来过，则直接退出
        if (grid[idx][idy] != LAND || record[idx][idy] == ENCLAVES) {
            return;
        }

        // 飞地记录
        record[idx][idy] = ENCLAVES;

        // 上面的地
        if (idy > 0) {
            transferEnclaves(grid, idx, idy - 1, record);
        }

        // 下面的地
        if (idy < grid[0].size() - 1) {
            transferEnclaves(grid, idx, idy + 1, record);
        }

        // 左面的地
        if (idx > 0) {
            transferEnclaves(grid, idx - 1, idy, record);
        }

        // 右面的地
        if (idx < grid.size() - 1) {
            transferEnclaves(grid, idx + 1, idy, record);
        }
    }

    // 获取记录中网格值为n的个数
    int getNumGrid(const vector<vector<int>>& grid, int n) {
        int cnt = 0;
        for (int idx = 0; idx < grid.size(); idx++) {
            for (int idy = 0; idy < grid[0].size(); idy++) {
                if (grid[idx][idy] == n) {
                    cnt++;
                }
            }
        }

        return cnt;
    }
};
// @lc code=end

