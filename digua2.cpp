/*
 * @Author       : zc
 * @Date         : 2022-02-13 10:21:58
 * @LastEditors  : zc
 * @LastEditTime : 2022-02-13 12:49:31
 * @Description  : file content
 * @FilePath     : \leetcode\digua2.cpp
 */
#include <iostream>
#include <vector>

using namespace std;


void get_all_lines(vector<pair<int, int>> &points, 
        vector<pair<pair<int, int>, pair<int, int>>> &lines)
{
    for(int i = 0; i < points.size() - 1; i++)
    {
        for (int j = i + 1; j < points.size(); j++) {
            lines.push_back(make_pair(points[i], points[j]));
        }
    }
}

bool is_equal(pair<pair<int, int>, pair<int, int>>& l1, 
            pair<pair<int, int>, pair<int, int>>& l2)
{
    return (l1.first.first - l1.second.first) *
            (l2.first.first - l2.second.first) +
            (l1.first.second)
}

bool is_vec(pair<pair<int, int>, pair<int, int>>& l1, 
            pair<pair<int, int>, pair<int, int>>& l2)
{
    
}

bool is_mid(pair<pair<int, int>, pair<int, int>>& l1, 
            pair<pair<int, int>, pair<int, int>>& l2)
{
    bool x = l1.first.first + l1.second.first == 
                l2.first.first + l2.second.first;
    bool y = l1.first.second + l1.second.second == 
                l2.first.second + l2.second.second;
    return x && y;
}

int get_zheng(vector<pair<pair<int, int>, pair<int, int>>> &lines)
{
    int cnt = 0;
    for(int i = 0; i < lines.size() - 1; i++)
    {
        for (int j = i + 1; j < lines.size(); j++) {
            if (is_mid(lines[i], lines[j]) &&
                is_vec(lines[i], lines[j]) &&
                is_equal(lines[i], lines[j])) {
                cnt++;
            }
        }
    }

    return cnt;
}

int main(void)
{
    vector<pair<int, int>> points;
    vector<pair<pair<int, int>, pair<int, int>>> lines;

    get_all_lines(points, lines);
    int cnt = get_zheng(lines);

    cout << cnt << endl;

    return 0;
}


