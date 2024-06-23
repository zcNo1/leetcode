/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */
#include <stdio.h>
#include <stdlib.h>

 struct ListNode {
    int val;
    struct ListNode *next;
 };
/**
 * @brief 获取两个个位数及上一位进位数的和对10的余数
 * 
 * @param num1 第一个个位数
 * @param num2 第二个个位数
 * @return int 取到的余数
 */
int getVal(int num1, int num2)
{
    int sum_num = 0;    // 当前位相加
    static int next_num = 0;   // 进数，当前位相加除以10得到，为下一位的数
    sum_num = num1 + num2 + next_num;
    next_num = sum_num / 10;
    return sum_num % 10;
}

/**
 * @brief 由当前节点的值获取对应的输出节点
 * 
 * @param num1 数值1
 * @param num2 数值2
 * @return struct ListNode* 取到的节点
 */
struct ListNode* getNode(int num1, int num2)
{
    struct ListNode *temNext = NULL;
    temNext = (struct ListNode *)calloc(1, sizeof(struct ListNode));
    if (temNext == NULL)
    {
        return NULL;
    }
    temNext->val = getVal(num1, num2);
    temNext->next = NULL;

    return temNext;
}

struct ListNode* addTwoNumbers(struct ListNode* l1, struct ListNode* l2){
    struct ListNode *ret = NULL, *tem = NULL;

    if (!l1 || !l2)
    {
        goto CLEAN;
    }

    ret = getNode(l1->val, l2->val);;
    if (ret == NULL)
    {
        goto CLEAN;
    }
    l1 = l1->next;
    l2 = l2->next;

    tem = ret;
    while (l1 != NULL || l2 != NULL)
    {
        // tem为NULL证明上一个节点出错，走出错处理
        if (tem == NULL)
        {
            goto CLEAN;
        }
        // 如果l1短，则会先走到NULL，处理l2剩余的值
        if (l1 == NULL)
        {
            tem->next = getNode(0, l2->val);
            tem = tem->next;
            l2 = l2->next;
            continue;
        }
        // 如果l2短，则会先走到NULL，处理l1剩余的值
        if (l2 == NULL)
        {
            tem->next = getNode(l1->val, 0);
            tem = tem->next;
            l1 = l1->next;
            continue;
        }
        tem->next = getNode(l1->val, l2->val);
        tem = tem->next;
        l1 = l1->next;
        l2 = l2->next;
    }
    int temNum = getVal(0, 0);
    if (temNum != 0)
    {
        tem->next = (struct ListNode *)calloc(1, sizeof(struct ListNode));
        if (tem->next == NULL)
        {
            goto CLEAN;
        }
        tem->next->val = temNum;
        tem->next->next = NULL;
    }

    return ret;

CLEAN:
    while (ret != NULL)
    {
        tem = ret->next;
        free(ret);
        ret = tem;
    }
    return ret;
}