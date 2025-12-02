use std::fs;

fn parse() -> Result<Vec<(i64, i64)>, String> {
    let input_str = fs::read_to_string("./src/input.txt")
        .map_err(|e| format!("Error occurred while reading input: {e}"))?;

    let ranges: Vec<(i64, i64)> = input_str
        .split(',')
        .filter(|s| !s.trim().is_empty())
        .map(|range_str| {
            let parts: Vec<&str> = range_str.split('-').collect();
            if parts.len() != 2 {
                return Err(format!("Invalid range format: {}", range_str));
            }
            let start = parts[0]
                .trim()
                .parse::<i64>()
                .map_err(|_| format!("Invalid start number: {}", parts[0]))?;
            let end = parts[1]
                .trim()
                .parse::<i64>()
                .map_err(|_| format!("Invalid end number: {}", parts[1]))?;
            Ok((start, end))
        })
        .collect::<Result<Vec<_>, _>>()?;

    Ok(ranges)
}

fn sum_invalid_ids(ids: &[(i64, i64)]) -> i64 {
    let mut sum = 0;
    for &(start, end) in ids {
        for num in start..=end {
            let num_str = num.to_string();
            if num_str.len() % 2 != 0 {
                continue;
            }

            let mid = num_str.len() / 2;
            let left = &num_str[..mid];
            let right = &num_str[mid..];

            if left == right {
                sum += num;
            }
        }
    }
    sum
}

fn is_invalid_id(num_str: &str) -> bool {
    let len = num_str.len();

    for repeat_len in 1..=len / 2 {
        let repeats = len / repeat_len;
        if !len.is_multiple_of(repeat_len) {
            continue;
        }

        let sequence = &num_str[..repeat_len];
        let mut is_repeat = true;

        for i in 1..repeats {
            let offset = i * repeat_len;
            if &num_str[offset..offset + repeat_len] != sequence {
                is_repeat = false;
                break;
            }
        }

        if is_repeat {
            return true;
        }
    }

    false
}

fn sum_invalid_ids_2(ids: &[(i64, i64)]) -> i64 {
    let mut sum = 0;
    for &(start, end) in ids {
        for num in start..=end {
            let num_str = num.to_string();
            if is_invalid_id(&num_str) {
                sum += num;
            }
        }
    }
    sum
}

fn main() {
    match parse() {
        Ok(ranges) => {
            let result = sum_invalid_ids(&ranges);
            println!("{}", result);
        }
        Err(e) => {
            eprintln!("Error: {e}");
            std::process::exit(1);
        }
    }
    match parse() {
        Ok(ranges) => {
            let result = sum_invalid_ids_2(&ranges);
            println!("{}", result);
        }
        Err(e) => {
            eprintln!("Error: {e}");
            std::process::exit(1);
        }
    }
}
