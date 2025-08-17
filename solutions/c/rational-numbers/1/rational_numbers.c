#include "rational_numbers.h"

rational_t reduce(rational_t r) {
    if (r.numerator == 0) {
        r.denominator = 1;
        goto done;
    }

    if (r.denominator < 0) {
        r.numerator *= -1;
        r.denominator *= -1;
        goto simplify;
    }

    if (r.denominator == 0) {
        if (r.numerator < 0) {
            r.numerator = -1;
        } else {
            r.numerator = 1;
        }
        goto done;
    }

    simplify:;
    rational_t gcd = r;
    // int gca = r.denominator;
    if (gcd.numerator < 0) gcd.numerator *= -1;
    if (gcd.denominator < 0) gcd.denominator *= -1;
    while (gcd.denominator != 0) {
        int nu = gcd.numerator,
            de = gcd.denominator;
        gcd.denominator = nu % de;
        gcd.numerator = de;
    }

    r.numerator /= gcd.numerator;
    r.denominator /= gcd.numerator;

    done:
    return r;
}



rational_t add(rational_t r1, rational_t r2) {
    if (r1.denominator == r2.denominator) {
        return reduce((rational_t){
            r1.numerator + r2.numerator,
            r1.denominator });
    }
    
    return reduce((rational_t){
        r1.numerator * r2.denominator + r2.numerator * r1.denominator,
        r1.denominator * r2.denominator });
}

rational_t subtract(rational_t r1, rational_t r2) {
    if (r1.denominator == r2.denominator) {
        return reduce((rational_t){
            r1.numerator - r2.numerator,
            r1.denominator });
    }

    return reduce((rational_t){
        r1.numerator * r2.denominator - r2.numerator * r1.denominator,
        r1.denominator * r2.denominator });
}

rational_t multiply(rational_t r1, rational_t r2) {
    return reduce((rational_t){
        r1.numerator * r2.numerator,
        r1.denominator * r2.denominator });
}

rational_t divide(rational_t r1, rational_t r2) {
    return reduce((rational_t){
        r1.numerator * r2.denominator,
        r1.denominator * r2.numerator });
}



rational_t absolute(rational_t r) {
    if (r.numerator < 0) r.numerator *= -1;
    if (r.denominator < 0) r.denominator *= -1;
    return reduce(r);
}

rational_t exp_rational(rational_t r, int16_t n) {
    if (n < 0) {
        int an = n * -1;
        return reduce((rational_t){
            pow(r.denominator, an),
            pow(r.numerator, an) });
    }
    
    return reduce((rational_t){
        pow(r.numerator, n),
        pow(r.denominator, n) });
}

float exp_real(int16_t n, rational_t r) {
    return pow(pow(n, r.numerator), 1./r.denominator);
}
