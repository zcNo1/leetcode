/*
 * @Author       : zc
 * @Date         : 2021-04-10 20:30:27
 * @LastEditors  : zc
 * @LastEditTime : 2021-04-10 20:56:59
 * @Description  : file content
 * @FilePath     : \leetcode\0263\C++\isUgly.cpp
 */
class Solution {
public:
    int diveN(int n, int dive) {
        while ( n % dive == 0 ) {
            n /= dive;
        }
        return n;
    }
    bool isUgly(int n) {
        if (n <= 0) {
            return false;
        }
        n = diveN(diveN(diveN(n, 5), 3), 2);
        return n == 1;
    }
};