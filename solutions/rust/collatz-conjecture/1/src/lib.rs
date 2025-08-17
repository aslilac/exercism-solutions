pub fn collatz(n: u64) -> Option<u64> {
    if (n < 1) {
        return None;
    }
    
    let mut i = n;
    let mut cycles = 0;
    while i > 1 {
        i = if i % 2 == 0 {
            i / 2
        } else {
            i.checked_mul(3)?.checked_add(1)?
        };
        cycles += 1;
    }
    return Some(cycles)
}
