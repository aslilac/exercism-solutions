pub fn brackets_are_balanced(string: &str) -> bool {
    let mut cleaned = string.to_string();
    cleaned.retain(|a| matches!(a, '{' | '}' | '[' | ']' | '(' | ')'));
    let mut chars = cleaned.chars();
    let mut stack = Vec::new();
    
    while let Some(a) = chars.next() {
        match a {
            '[' => stack.push(']'),
            '{' => stack.push('}'),
            '(' => stack.push(')'),
            _ => {
                if stack.pop() != Some(a) {
                    return false;
                }
            },
        }
    };

    return stack.is_empty();
}
