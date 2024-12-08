use std::str::FromStr;

fn parse_input(input: &str) -> Vec<(usize, Vec<usize>)> {
    let mut data = Vec::new();
    for line in input.lines() {
        let (target_str, nums_str) = line.split_once(": ").unwrap();
        let nums: Vec<usize> = nums_str
            .split_whitespace()
            .map(|s| FromStr::from_str(s).unwrap())
            .collect();
        let target: usize = FromStr::from_str(target_str).unwrap();

        data.push((target, nums));
    }

    data
}

pub fn bridge_repair(input: &str) -> usize {
    let equations = parse_input(input);
    let mut sum = 0;
    for (target, nums) in equations {
        if can_solve(0, 0, target, &nums, false) {
            sum += target;
        }
    }

    sum
}

pub fn bridge_repair_2(input: &str) -> usize {
    let equations = parse_input(input);
    let mut sum = 0;
    for (target, nums) in equations {
        if can_solve(0, 0, target, &nums, true) {
            sum += target;
        }
    }

    sum
}

fn can_solve(
    n: usize,
    current_val: usize,
    target: usize,
    nums: &Vec<usize>,
    concat_true: bool,
) -> bool {
    if n == nums.len() {
        return current_val == target;
    }

    let mut solved = can_solve(n + 1, current_val + nums[n], target, nums, concat_true)
        || can_solve(n + 1, current_val * nums[n], target, nums, concat_true);

    if concat_true {
        solved = solved
            || can_solve(
                n + 1,
                concat(current_val, nums[n]),
                target,
                nums,
                concat_true,
            );
    }

    solved
}

fn concat(a: usize, b: usize) -> usize {
    (a.to_string() + b.to_string().as_str()).parse().unwrap()
}
