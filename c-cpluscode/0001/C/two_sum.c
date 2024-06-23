// /**
//  * @brief 两数之和，暴力搜索，但是时间测试没有通过
//  * 
//  * Note: The returned array must be malloced, assume caller calls free().
//  * @param nums array want to get child sum
//  * @param numsSize array size
//  * @param target target sum
//  * @param returnSize return match array size
//  * 
//  * @return array contain two num, which from nums and sum is target
//  */
// int *twoSum(int *nums, int numsSize, int target, int *returnSize)
// {
//     if (!nums || !returnSize || numsSize < 2)
//     {
//         fprintf(stderr, "%s", "param is invaild");
//         return NULL;
//     }

//     int *returnNums = NULL;
//     returnNums = (int *)calloc(2, sizeof(int));
//     if (!returnNums)
//     {
//         return NULL;
//     }

//     int i = 0, j = 0;
//     for (i = 0; i < numsSize; i++)
//     {
//         for (j = i + 1; j < numsSize; j++)
//         {
//             // if sum of two element is equal target, output their index and size of 2
//             if (*(nums + i) + *(nums + j) == target)
//             {
//                 returnNums[0] = i;
//                 returnNums[1] = j;
//                 *returnSize = 2;
//                 return returnNums;
//             }
//         }
//     }

//     // if all element dont meet requirement, return NULL and size of 0
//     *returnSize = 0;
//     if (returnNums)
//     {
//         free(returnNums);
//         returnNums = NULL;
//     }
//     return NULL;
// }

#include <stdio.h>
#include "uthash.h"

#define FAIL (int)-1
#define SUCCESS (int)0

#define EXIST (int)1

/**
 * @brief 含有hash表的结构体
 * 
 * @param key 数组中的数字
 * @param val 数组中数字的位置
 */
typedef struct two_sum_st
{
    int key; // 数组中的数字
    int val; // 数组中数字的位置

    UT_hash_handle hh;
} two_sum_st;

// 
static two_sum_st *hash_head = NULL;

/**
 * @brief 返回hash中对应数字的位置，内部实现
 * 
 * @param key 数字
 * @return two_sum_st* 成功/失败(没有该键) 找到的结构体/NULL
 */
static two_sum_st *find_num_(int key)
{
    two_sum_st *mem;
    HASH_FIND_INT(hash_head, &key, mem); /* s: output pointer */
    return mem;
}

/**
 * @brief 返回hash中对应数字的位置
 * 
 * @param key 数字
 * @return int 成功/失败(没有该键) 数字的位置/FAIL(-1)
 */
int find_num(int key)
{
    two_sum_st *hash_child = NULL;
    hash_child = find_num_(key);
    if (hash_child == NULL || hash_child->val < 0)
    {
        return FAIL;
    }

    return hash_child->val;
}


/**
 * @brief 往hash表中添加元素，如果存在该元素则替换其val，内部实现
 * 
 * @param key 数组中的数字
 * @param val 数组中数字的位置
 * @return int 成功/失败 SUCCESS(0)/FAIL(-1)
 */
static int add_num_(int ikey, int ival)
{
    two_sum_st *mem = NULL;
    mem = find_num_(ikey);
    if (mem != NULL)
    {
        mem->val = ival;
        return SUCCESS;
    }
    mem = (two_sum_st *)calloc(1, sizeof(two_sum_st));
    if (mem == NULL)
    {
        return FAIL;
    }
    mem->key = ikey;
    mem->val = ival;
    HASH_ADD_INT(hash_head, key, mem); // 插入数据

    return SUCCESS;
}

/**
 * @brief 往hash表中添加元素，如果存在该元素则替换其val
 * 
 * @param key 数组中的数字
 * @param val 数组中数字的位置
 * @return int 成功/失败 SUCCESS(0)/FAIL(-1)
 */
int add_num(int key, int val)
{
    return add_num_(key, val);
}

/**
 * @brief 两数之和，利用hash表去做
 * 
 * @param nums 
 * @param numsSize 
 * @param target 
 * @param returnSize 
 * @return int* 
 */
int *twoSum(int *nums, int numsSize, int target, int *returnSize)
{
    if (!nums || !returnSize || numsSize < 2)
    {
        // fprintf(stderr, "%s", "param is invaild");
        return NULL;
    }

    int *returnNums = NULL;
    returnNums = (int *)calloc(2, sizeof(int));
    if (!returnNums)
    {
        return NULL;
    }


    two_sum_st *tem;
    int index = -1;
    int i = 0;

    for (i = 0; i < numsSize; i++)
    {
        printf("this %d number\n", i);
        index = find_num(target - nums[i]);
        printf("index is %d\n", index);
        if (index >= 0)
        {
            returnNums[0] = index;
            returnNums[1] = i;
            *returnSize = 2;
            return returnNums;
        }
        add_num(nums[i], i);
        for(tem=hash_head; tem != NULL; tem=(two_sum_st*)(tem->hh.next)) 
        {
            printf("user %d, cookie %d\n", tem->key, tem->val);
        }
    }

    // if all element dont meet requirement, return NULL and size of 0
    *returnSize = 0;
    if (returnNums)
    {
        free(returnNums);
        returnNums = NULL;
    }
    return NULL;
}

#define CNT 2
int main(void)
{
    int nums[CNT] = {3,3};
    int target = 6;
    int returnSize = 0;
    int i = 0;
    int *ret;

    ret = twoSum(nums, CNT, target, &returnSize);
    if (ret != NULL)
    {
        for (;i < 2; i++)
        {
            printf("result is nums[%d] = %d\n", ret[i], nums[ret[i]]);
        }
    }
    

    printf("%d\n", returnSize);

    return 0;
}