/*
 * @Author       : zc
 * @Date         : 2021-06-14 14:39:22
 * @LastEditors  : zc
 * @LastEditTime : 2021-06-14 16:12:55
 * @Description  : file content
 * @FilePath     : \leetcode\705.设计哈希集合.cpp
 */
/*
 * @Author       : zc
 * @Date         : 2021-06-14 14:39:22
 * @LastEditors  : zc
 * @LastEditTime : 2021-06-14 15:21:55
 * @Description  : file content
 * @FilePath     : \leetcode\705.设计哈希集合.cpp
 */
/*
 * @lc app=leetcode.cn id=705 lang=cpp
 *
 * [705] 设计哈希集合
 */

// @lc code=start
#include <iostream>
#include <vector>
#include <list>
#include <memory.h>


class MyHashSet {
private:
    std::vector<std::list<int>> set;
    static const int base = 769;
    static int hashVal(int key) {
        return key % base;
    }
public:
    /** Initialize your data structure here. */
    MyHashSet() : set(base) {

    }
    
    void add(int key) {
        int pos = hashVal(key);
        for (auto i : set[pos]){
            if (i == key) {
                return;
            }
        }
        set[pos].push_back(key);
        return;
    }
    
    void remove(int key) {
        int pos = hashVal(key);
        for (auto it = set[pos].begin(); it != set[pos].end(); ++it){
            if (*it == key) {
                set[pos].erase(it);
                return;
            }
        }
        return;
    }
    
    /** Returns true if this set contains the specified element */
    bool contains(int key) {
        int pos = hashVal(key);
        for (auto i : set[pos]){
            if (i == key) {
                return true;
            }
        }
        return false;
    }
};

/**
 * Your MyHashSet object will be instantiated and called as such:
 * MyHashSet* obj = new MyHashSet();
 * obj->add(key);
 * obj->remove(key);
 * bool param_3 = obj->contains(key);
 */
// @lc code=end

