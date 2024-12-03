pub fn mull_it_over(input: &str) -> i32 {
    let mem: Vec<&str> = input.split("mul(").collect();
    let mut sum = 0;
    for m in mem {
        if m.len() < 4 {
            continue;
        }
        if let Some((num1, num2)) = get_nums(m) {
            sum += num1 * num2;
        }
    }

    sum
}

pub fn mull_it_over_2(input: &str) -> i32 {
    let mut pos = 0;
    let mut sum = 0;
    let mut enabled = true;
    let chars: Vec<char> = input.chars().collect();
    let len = chars.len();

    while pos < len {
        if !chars[pos].is_ascii_alphabetic() {
            pos += 1;
            continue;
        }

        if pos + 3 < len && &input[pos..pos + 4] == "do()" {
            enabled = true;
            pos += 4;
            continue;
        }

        if pos + 6 < len && &input[pos..pos + 7] == "don't()" {
            enabled = false;
            pos += 7;
            continue;
        }

        if pos + 3 < len && &input[pos..pos + 4] == "mul(" {
            let start = pos + 4;
            if let Some(end_pos) = input[start..].find(')') {
                if let Some((num1, num2)) = get_nums(&input[pos + 4..start + end_pos + 1]) {
                    if enabled {
                        sum += num1 * num2;
                    }
                    pos = start + end_pos + 1;
                    continue;
                } else {
                    pos += 4;
                    continue;
                }
            } else {
                pos += 4;
                continue;
            }
        }
        pos += 1;
    }

    sum
}

fn get_nums(s: &str) -> Option<(i32, i32)> {
    let split = s.split_once(')');
    if split.is_none() {
        return None;
    }

    let (s, _) = split.unwrap();

    if s.len() < 3 || s.len() > 7 {
        return None;
    }

    let num_split = s.split_once(',');
    if num_split.is_none() {
        return None;
    }
    let (num1_str, num2_str) = num_split.unwrap();

    if (num1_str.len() < 1 || num1_str.len() > 3) && (num2_str.len() < 1 || num2_str.len() > 3) {
        return None;
    }

    let num1 = num1_str.parse::<i32>();
    let num2 = num2_str.parse::<i32>();

    if num1.is_err() || num2.is_err() {
        return None;
    }

    return Some((num1.unwrap(), num2.unwrap()));
}
