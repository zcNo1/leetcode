/*
 * @Author       : zc
 * @Date         : 2021-04-13 22:59:07
 * @LastEditors  : zc
 * @LastEditTime : 2021-04-13 23:19:05
 * @Description  : file content
 * @FilePath     : \leetcode\0783\C++\minDiffInBST.cpp
 */

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
#include "0000/C++/inc_C++.h"
class Solution {
public:
    int minDiffInBST(TreeNode* root) {
        if (root == NULL) {
            return 0;
        }
        if (root->left == NULL) {
            return INT_MAX;
        }
        if (root->right == NULL) {
            return root->val - root->left->val;
        }
        int l = minDiffInBST(root->left);
        int r = minDiffInBST(root->right);
        int m = min(root->val - root->left->val, root->right->val - root->val);
        return min(l, min(r, m));
    }
};