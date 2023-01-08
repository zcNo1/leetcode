/*
 * @Author       : zc
 * @Date         : 2021-06-30 22:13:13
 * @LastEditors  : zc
 * @LastEditTime : 2021-06-30 22:16:11
 * @Description  : file content
 * @FilePath     : \leetcode\22.cpp
 */
/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */

#include <iostream>
#include <string>
#include <algorithm>
#include <cstring>

using namespace std;

struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};
const char str_s = '[';
const char str_e = ']';
const char sep = ',';
const char char_null = 'n';
const char str_null = "null";
class Codec {
public:

    // Encodes a tree to a single string.
    string serialize(TreeNode* root) {
        
    }

    // Decodes your encoded data to tree.
    TreeNode* deserialize(string data) {
        strstr(              )
        
    }
};

// Your Codec object will be instantiated and called as such:
// Codec codec;
// codec.deserialize(codec.serialize(root));