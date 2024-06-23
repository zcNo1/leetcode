/*
 * @Author       : zc
 * @Date         : 2022-02-13 10:21:58
 * @LastEditors  : zc
 * @LastEditTime : 2022-02-13 12:59:46
 * @Description  : file content
 * @FilePath     : \leetcode\digua1.cpp
 */
#include <iostream>
#include <vector>

using namespace std;

void get(int k, int n, vector<int> &record, int &curRet);
int getjiechen(int n);

int main(void)
{
    int n, k;
    cin >> n >> k;
    int ret = 0;
    vector<int> record(n, 0);
    for (int i = 0; i < n; i++)
    {
        record[i] = i+1; 
    }

    get(k-1, n,record, ret);
    cout << ret << endl;

    return 0;
}

void get(int k, int n, vector<int> &record, int &curRet)
{
    // 如果n等于1，证明就剩最后一个数字，直接累计，退出递归
    if (n == 1) {
        curRet = curRet * 10 + record[0];
        return;
    }
    // 获取n-1的阶乘
    int sub = getjiechen(n - 1);
    int shang = k / sub;
    int yu = k % sub;
    // 累计数字
    curRet = curRet * 10 + record[shang];
    // 删除已经记录的数字
    record.erase(record.begin() + shang);

    // 递归
    get(yu, n - 1, record, curRet);
}

int getjiechen(int n)
{
    if (n < 1){
        return 0;
    }
    int ret = 1;
    while (n > 1) {
        ret *= n;
        n--;
    }
    return ret;
}