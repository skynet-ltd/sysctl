#include <stdlib.h>

struct response
{
    size_t len;
    char *buffer;
};

int count_dots(const char *name);
struct response *query_sysctl(const char *name, char **error);