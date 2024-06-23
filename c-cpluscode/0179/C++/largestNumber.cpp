/*
 * @Author       : zc
 * @Date         : 2021-04-12 22:37:53
 * @LastEditors  : zc
 * @LastEditTime : 2021-05-16 13:43:48
 * @Description  : file content
 * @FilePath     : \leetcode\0179\C++\largestNumber.cpp
 */
#include "0000/C++/inc_C++.h"
#define STRLEN (20)
const int i3;
bool compare1 (string a, string b) {
    int i = 42;
    int *p;
    int *&r1 = p;
    const int &r2 = i;
    const int &r3 = 10;
    
    const int i2 = 0;
        size_t len1 = a.size();
        size_t len2 = b.size();
        if (len1 == len2) {
            cout << "len1 == len2" << endl;
            cout << "a[" << len1 << "]: " << a << "\tb[" << len2 << "]: " << b << endl;
            return a >= b;
        } else if (len1 < len2) {
            b[len1 + 1] = 0;
            a += a[0];
            cout << "len1 < len2" << endl;
            cout << "a[" << len1 << "]: " << a << "\tb[" << len2 << "]: " << b << endl;
            return a >= b; 
        } else {
            a[len2 + 1] = 0;
            b += b[0];
            cout << "len1 > len2" << endl;
            cout << "a[" << len1 << "]: " << a << "\t[" << len1 << "]: " << b << endl;
            return a > b;
        }
    }
class Solution {
public:
    int dtostrArr(vector<string>& numstr, vector<int>& nums) {
        char str[STRLEN + 1] = { 0 };
        for (auto it = nums.begin(); it < nums.end(); it++){
            snprintf(str, STRLEN, "%d", (*it));
            numstr.push_back(str);
            memset(str, 0, STRLEN);
        }
        return 0;
    }
    string largestNumber(vector<int>& nums) {
        vector<string> numstr;
        string ret = "";
        dtostrArr(numstr, nums);
        showNumstr(numstr);
        sort(numstr.begin(), numstr.end(), compare1);
        showNumstr(numstr);
        for(int i = 0; i < numstr.size(); i++) {
            ret += numstr[i];
        }
        return ret;
    }
    void showNumstr(vector<string>& numstr) {
        size_t cnt = numstr.size();
        cout << "begin" << endl;
        for (size_t i = 0; i < cnt; i++){
            cout << "numstr[" << i << "] = " << numstr[i] << endl;
        }
        cout << "end" << endl;
        return;
    }
};
