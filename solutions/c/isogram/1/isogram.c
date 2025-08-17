#include "isogram.h"

#include <stdint.h>
#include <string.h>

bool is_isogram(const char phrase[]) {
    if (phrase == NULL)
        return false;
    
    uint32_t state = 0;
    int len = strlen(phrase);

    for (int i = 0; i < len; ++i) {
        uint32_t check = state;
        char it = phrase[i] & 0xdf; // 0b1101'1111

        if (it >= 'A' && it <= 'Z') {
            state |= 1 << (it - 'A');
            if (check == state)
                return false;
        }
    }

    return true;
}
