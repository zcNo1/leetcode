/*
 * @Author       : zc
 * @Date         : 2021-06-13 17:49:23
 * @LastEditors  : zc
 * @LastEditTime : 2021-06-14 11:18:54
 * @Description  : file content
 * @FilePath     : \leetcode\146.lru-缓存机制.cpp
 */
/*
 * @lc app=leetcode.cn id=146 lang=cpp
 *
 * [146] LRU 缓存机制
 */

// @lc code=start
#include <iostream>
#include <unordered_map>
#include <list>
typedef struct Cache_ {
    int value;                      // 存放的数据value
    std::list<int>::iterator it;    // 存放关键字key的list迭代器
}Cache;
class LRUCache {
public:
    LRUCache(int capacity) {
        if (capacity <= 0) {
            cap = 1;
        } else {
            cap = capacity;
        }
    }
    
    int get(int key) {
        auto it = map.find(key);
        if (it == map.end()) {
            return -1;
        }
        // 旧的key的list迭代器，重新插入到list尾端，生成新的迭代器
        l.erase(it->second.it);
        it->second.it = l.insert(l.end(), key);
        return it->second.value;
    }
    
    void put(int key, int value) {
        auto it = map.find(key);
        if (it != map.end()) {
            l.erase(it->second.it);
            it->second.it = l.insert(l.end(), key);
            it->second.value = value;
            return;
        }
        if (l.size() >= cap) {
            map.erase(l.front());
            l.pop_front();
        }
        std::list<int>::iterator lit = l.insert(l.end(), key);
        Cache cache= {value, lit};
        map.insert(std::make_pair(key, cache));
    }
private:
    int cap;
    std::list<int> l;
    std::unordered_map<int, Cache> map;
};

/**
 * Your LRUCache object will be instantiated and called as such:
 * LRUCache* obj = new LRUCache(capacity);
 * int param_1 = obj->get(key);
 * obj->put(key,value);
 */
// @lc code=end

