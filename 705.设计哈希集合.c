/*
 * @lc app=leetcode.cn id=705 lang=c
 *
 * [705] 设计哈希集合
 */

// @lc code=start

#include<stdio.h>
#include<stdlib.h>
typedef struct HashNode_
{
    int key;
    struct HashNode_ *next;
}HashNode;

typedef struct {
    HashNode *table;
} MyHashSet;

/** Initialize your data structure here. */

const int base = 769;

MyHashSet* myHashSetCreate() {
    MyHashSet *obj = NULL;
    obj = (MyHashSet*)calloc(1, sizeof(MyHashSet));
    if (!obj) {
        obj = NULL;
        return NULL;
    }
    obj->table = (HashNode*)calloc(base, sizeof(HashNode));
    if (!obj->table) {
        free(obj);
        obj = NULL;
        return NULL;
    }
    for (int i = 0; i < base; i++) {
        obj->table[i].key = 0;
        obj->table[i].next = NULL;
    }
    return obj;
}

void myHashSetAdd(MyHashSet* obj, int key) {
        int pos = key % base;
        HashNode *pHead = &(obj->table[pos]);
        while (pHead->next) {
            if (pHead->next->key == key) {
                return;
            }
            pHead = pHead->next;
        }
        HashNode *newNode = (HashNode*)calloc(1, sizeof(HashNode));
        newNode->key = key;
        newNode->next = pHead->next;
        pHead->next = newNode;

        return;
}

void myHashSetRemove(MyHashSet* obj, int key) {
        int pos = key % base;
        HashNode *pHead = &(obj->table[pos]);
        while (pHead->next) {
            if (pHead->next->key == key) {
                HashNode *tmp = pHead->next;
                pHead->next = tmp->next;
                free(tmp);
                break;
            }
            pHead = pHead->next;
        }
}

/** Returns true if this obj contains the specified element */
bool myHashSetContains(MyHashSet* obj, int key) {
        int pos = key % base;
        HashNode *pHead = &(obj->table[pos]);
        while (pHead->next) {
            if (pHead->next->key == key) {
                return true;
            }
            pHead = pHead->next;
        }

        return false;
}

static void free_hashnode(HashNode *pNode) {
    HashNode *tmp = NULL;
    while (pNode) {
        tmp = pNode;
        pNode = tmp->next;
        free(tmp);
        tmp = NULL;
    }
    return;
}

void myHashSetFree(MyHashSet* obj) {
    int idx = 0;
    if (!obj){
        return;
    }
    if (obj->table) {
        for(idx = 0; idx < base; idx++) {
            free_hashnode(obj->table[idx].next);
        }
        free(obj->table);
        obj->table = NULL;
    }
    // free(obj);
    // obj = NULL;
}

/**
 * Your MyHashSet struct will be instantiated and called as such:
 * MyHashSet* obj = myHashSetCreate();
 * myHashSetAdd(obj, key);
 
 * myHashSetRemove(obj, key);
 
 * bool param_3 = myHashSetContains(obj, key);
 
 * myHashSetFree(obj);
*/
// @lc code=end