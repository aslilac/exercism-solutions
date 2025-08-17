#include "collatz_conjecture.h"

int steps(int n) {
    if (n < 1) {
        return -1;
    }
    
    if (n == 1) {
        return 0;
    }

    int count = 0;
    while (n > 1) {
        if (n % 2 == 0) {
            n = n / 2;
        } else {
            n = 3 * n + 1;
        }
        count += 1;
    }

    return count;
}