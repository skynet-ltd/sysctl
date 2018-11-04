#include "query_sysctl.h"
#include <string.h>
#include <sys/sysctl.h>

#include <stdio.h>

int count_dots(const char *name)
{
    int dots = 0;
    char *copy = (char *)name;
    while (*copy)
    {
        if (*copy == '.')
        {
            dots++;
        }
        copy++;
    }
    return dots;
}

struct response *query_sysctl(const char *name, char **error)
{
    int dots = count_dots(name);
    int *mbi = (int *)malloc(sizeof(int) * (dots + 1));
    size_t len;

    int res = sysctlnametomib(name, mbi, &len);
    if (res != 0)
    {
        *error = (char *)"can't convert to mbi";
        free(mbi);
        return NULL;
    }

    res = sysctl(mbi, dots + 1, NULL, &len, NULL, 0);
    if (res != 0)
    {
        *error = (char *)"can't get result buffer length";
        free(mbi);
        return NULL;
    }

    char *buffer = (char *)malloc(sizeof(char) * len);
    res = sysctl(mbi, dots + 1, buffer, &len, NULL, 0);
    if (res != 0)
    {
        free(mbi);
        free(buffer);
        *error = (char *)"can't query sysctl";
        return NULL;
    }

    return &(struct response){.len = len, .buffer = buffer};
}