/*
 * @Author       : zc
 * @Date         : 2022-02-13 20:23:24
 * @LastEditors  : zc
 * @LastEditTime : 2022-02-13 21:21:44
 * @Description  : file content
 * @FilePath     : \leetcode\1114.按序打印.cpp
 */
/*
 * @lc app=leetcode.cn id=1114 lang=cpp
 *
 * [1114] 按序打印
 */

#include<iostream>
#include<pthread.h>
// @lc code=start
class Foo {
private:
    pthread_mutex_t ;
    pthread_mutex_t ma,mb;
public:
    Foo() {
        pthread_mutex_init(&ma, NULL);
        pthread_mutex_init(&mb, NULL);
        pthread_mutex_lock(&ma);
        pthread_mutex_lock(&mb);
    }


    void first(function<void()> printFirst) {
        // printFirst() outputs "first". Do not change or remove this line.
        printFirst();
        pthread_mutex_unlock(&ma);
    }

    void second(function<void()> printSecond) {
        pthread_mutex_lock(&ma);
        
        // printSecond() outputs "second". Do not change or remove this line.
        printSecond();
        
        pthread_mutex_unlock(&ma);
        pthread_mutex_unlock(&mb);
    }

    void third(function<void()> printThird) {
        pthread_mutex_lock(&mb);
        
        // printThird() outputs "third". Do not change or remove this line.
        printThird();
    }
};
// @lc code=end

