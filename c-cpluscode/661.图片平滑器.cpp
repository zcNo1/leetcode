/*
 * @lc app=leetcode.cn id=661 lang=cpp
 *
 * [661] 图片平滑器
 */
#include <iostream>
#include <vector>
#include <stack>
using namespace std;
// @lc code=start
class Solution {
private:
    int row, col;
public:
    vector<vector<int>> imageSmoother(vector<vector<int>>& img) {
        row = img.size();
        if (row == 0) {
            return vector<vector<int>>();
        }
        col = img[0].size();
        if (col == 0) {
            return vector<vector<int>>(row, vector<int>());
        }
        vector<vector<int>> ret_img(row, vector<int>(col, 0));
        for (int i = 0; i < row; i++) {
            for (int j = 0; j< col; j++) {
                ret_img[i][j] = cal(img, i, j);
            }
        }

        return ret_img;
    }
    int cal(vector<vector<int>>& img, int br, int bc) {
        int r = 0, c = 0;
        int cnt = 0, num = 0;
        for(int i = -1; i <= 1; i++) {
            for(int j = -1; j <= 1; j++) {
                r = br + i;
                c = bc + j;
                if (r >= 0 && r < row && c >=0 && c < col){
                    num += img[r][c];
                    cnt++;
                }
            }
        }

        if (cnt == 0) {
            return 0;
        }
        return num / cnt;
    }
};
// @lc code=end

