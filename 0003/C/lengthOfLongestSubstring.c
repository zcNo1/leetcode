int lengthOfLongestSubstring(char * s){
    if (!s)
    {
        return 0;
    }
    int len = strlen(s), i = 0, j = 0;
    int maxDis = 0, dis = 0;
    int cnt[256] = { 0 };   //初始化为0
    int *dn = (int*)calloc(len + 1, sizeof(int));    //以当前字符为结尾的最长不重复子串的长度
    if (dn == NULL)
    {
        return -1;
    }
    i = 1;
    while(*s != '\0')
    {
        j = *s - '\0';
        // 获取该字符与上一次出现的距离 dis
        dis = i - cnt[j];
        // dn表示以当前字符为结尾的最长不重复子串的长度
        // ss表示以当前字符为结尾的最长不重复子串
        // 如果ss[i]与ss[i-1]里面的字符串不重复，则dn[i] = dn[i - 1]+1
        // 如果ss[i]与ss[i-1]里面的字符串重复，dis必定小于dn[i - 1]+1且dn[i]必定等于dis
        dn[i] = dn[i - 1] + 1 < dis ? dn[i - 1] + 1 : dis;
        // 获取最大的不重复子串长度
        maxDis = dn[i] > maxDis ? dn[i] : maxDis;
        cnt[j] = i;
        i++;
        s++;
    }

    return maxDis;
}