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
    let mem: Vec<&str> = input.split("don't()").collect();
    let mut sum = 0;
    sum += mull_it_over(mem[0]);

    for m in 1..mem.len() {
        let cmd = mem[m];
        let cmd: Vec<&str> = cmd.split("do()").collect();
        for i in 1..cmd.len() {
            sum += mull_it_over(cmd[i]);
        }
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
