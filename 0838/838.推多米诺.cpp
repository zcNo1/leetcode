/*
 * @Author       : zc
 * @Date         : 2022-02-21 22:52:02
 * @LastEditors  : zc
 * @LastEditTime : 2022-02-21 23:35:43
 * @Description  : file content
 * @FilePath     : \leetcode\0838\838.推多米诺.cpp
 */
/*
 * @lc app=leetcode.cn id=838 lang=cpp
 *
 * [838] 推多米诺
 */

#include <iostream>
#include <string>

using namespace std;

// @lc code=start
class Solution {
public:
    string pushDominoes(string dominoes) {
        size_t f = 0;
        size_t s = 0;
        string ret;

        while(dominoes[s]=='.') {
            s++;
        }

        s_push(ret, s-f, dominoes[s]);
        f=s;
        s++;
        while(s<dominoes.size()){
            if (dominoes[s] =='.') {
                s++;
                continue;
            }
            if (dominoes[f] == dominoes[s]) {
                s_push(ret, s-f, dominoes[s]);
                f=s;
                s++;
                continue;
            }
            if ((s-f)%2==0) {
                size_t mid = (s-f)/2;
                s_push(ret, mid, dominoes[f]);
                s_push(ret, 1, '.');
                s_push(ret, mid-1, dominoes[s]);
            } else {
                size_t mid = (s-f+1)/2;
                s_push(ret, mid, dominoes[f]);
                s_push(ret, mid-1, dominoes[f]);
            }
            f=s;
            s++;
        }
        if (f<dominoes.size()) {
            s_push(ret, dominoes.size() - f, dominoes[f]);
        }
        return ret;
    }

    void s_push(string &str, size_t n, char ch)
    {
        while(n>0){
            str.push_back(ch);
            n--;
        }
    }
};

int main(void) {
    Solution s;

    cout << "123"<< s.pushDominoes(".L.R...LR..L..") << endl;
    return 0;
}
// @lc code=end

