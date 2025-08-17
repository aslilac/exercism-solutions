#include "reverse_string.h"
#include <stdlib.h>
#include <string.h>

char *reverse(const char *text) {
    int len = strlen(text);
    char* txet = malloc(len);

    for (int i = 0; i < len; ++i) {
        txet[i] = text[len - 1 - i];
    }

    return txet;
}
