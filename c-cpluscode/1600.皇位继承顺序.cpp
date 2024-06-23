/*
 * @Author       : zc
 * @Date         : 2021-06-20 13:59:58
 * @LastEditors  : zc
 * @LastEditTime : 2021-06-20 15:51:09
 * @Description  : file content
 * @FilePath     : \leetcode\1600.皇位继承顺序.cpp
 */
/*
 * @lc app=leetcode.cn id=1600 lang=cpp
 *
 * [1600] 皇位继承顺序
 */

// @lc code=start
#include<iostream>
#include<vector>
#include<unordered_map>
#include<list>
#include<new>
#include<cstdlib>
using namespace std;
typedef struct Family_{
    string name;
    bool is_dead;
    list<Family*> children;
    Family_() {
    }
}Family;

class ThroneInheritance {
private:
    Family* king;
    vector<string> order;
    unordered_map<string, Family*> map_order;
    Family* buildFamily(string fatherName) {
        Family *fam = new (nothrow) Family;
        if (fam == NULL) {
            return NULL;
        }
        fam->is_dead = false;
        fam->name = fatherName;
        fam->children = list<Family*>(0);
        map_order.insert(make_pair(fatherName, fam));
        return fam;
    }
    void getOrder(Family* father_fam) {
        if(!father_fam->is_dead) {
            order.push_back(father_fam->name);
        }
        
        auto child = father_fam->children;
        auto child_fam = child.begin();
        for(; child_fam != child.end(); ++child_fam) {
            getOrder(*(child_fam));
        }

        return;
    }
public:
    ThroneInheritance(string kingName) : map_order(0), king(NULL) {
        auto fam = buildFamily(kingName);
        if (!fam) {
            return;
        }
        king = fam;
    }
    void birth(string parentName, string childName) {
        // find father
        auto father = map_order.find(parentName);
        auto &father_fam = father->second;
        // build child family
        auto child_fam = buildFamily(childName);
        if (!child_fam) {
            return;
        }
        // add child to father family
        father_fam->children.push_back(child_fam);
    }
    
    void death(string name) {
        auto it = map_order.find(name);
        it->second->is_dead = true;
    }
    
    vector<string> getInheritanceOrder() {
        order.clear();
        getOrder(king);
        return order;
    }
};

/**
 * Your ThroneInheritance object will be instantiated and called as such:
 * ThroneInheritance* obj = new ThroneInheritance(kingName);
 * obj->birth(parentName,childName);
 * obj->death(name);
 * vector<string> param_3 = obj->getInheritanceOrder();
 */
// @lc code=end

