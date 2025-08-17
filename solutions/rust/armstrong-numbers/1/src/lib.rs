pub fn is_armstrong_number(num: u32) -> bool {
    if num == 0 {
        return true;
    }
    
    let mut digits = 1;
    loop {
        let checked_pow = 10_u32.checked_pow(digits);
        if let Some(pow) = checked_pow {
            if pow > num {
                break;
            }
            digits += 1;
        } else {
            return false;
        };
    }
    
    let num = num;
    let mut sum = 0;

    for i in 0..digits {
        sum += (num / 10_u32.pow(i) % 10).pow(digits);
    }

    println!("num: {}, sum: {}, digits: {}", num, sum, digits);

    num == sum
}
