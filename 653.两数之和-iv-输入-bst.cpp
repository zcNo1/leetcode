/*
 * @lc app=leetcode.cn id=653 lang=cpp
 *
 * [653] 两数之和 IV - 输入 BST
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */
#include <iostream>
#include <vector>
#include <stack>
using namespace std;
class Solution {
public:
    bool findTarget(TreeNode* root, int k) {
        vector<int> list;
        midTrans(root, list);
        int l = 0, r =list.size() - 1;
        int sum = 0;
        while (l < r) {
            sum = list[l] + list[r];
            if (sum == k) {
                return true;
            } else if (sum < k) {
                l++;
            } else {
                r--;
            }
        }
        return false;
    }

    void midTrans(TreeNode* root, vector<int> &list) {
        if (!root) {
            return;
        }
        midTrans(root->left, list);
        list.push_back(root->val);
        midTrans(root->right, list);
        return;
    }
};
// @lc code=end

