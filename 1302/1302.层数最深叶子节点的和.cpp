/*
 * @Author       : zc
 * @Date         : 2022-02-20 18:17:49
 * @LastEditors  : zc
 * @LastEditTime : 2022-02-20 18:27:24
 * @Description  : file content
 * @FilePath     : \leetcode\1302\1302.层数最深叶子节点的和.cpp
 */
/*
 * @lc app=leetcode.cn id=1302 lang=cpp
 *
 * [1302] 层数最深叶子节点的和
 */

// @lc code=start

// Definition for a binary tree node.
  struct TreeNode {
      int val;
      TreeNode *left;
      TreeNode *right;
      TreeNode() : val(0), left(nullptr), right(nullptr) {}
      TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
      TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
  };

#include <iostream>
#include <list>
using namespace std;
class Solution {
public:
    int deepestLeavesSum(TreeNode* root) {
        list<TreeNode*> li;

        li.push_back(root);

        int ret = 0, l_ret = 0;
        int sz = li.size();
        int cnt = 0;
        while (!li.empty()) {
            ret += li.front()->val;
            if (li.front()->left)
                li.push_back(li.front()->left);
            if (li.front()->right)
                li.push_back(li.front()->right);
            li.pop_front();
            cnt++;
            if (sz == cnt) {
                cnt = 0;
                sz = li.size();
                l_ret = ret;
                ret = 0;
            }
        }

        return l_ret;
    }
};
// @lc code=end

