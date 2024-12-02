use std::{collections::HashMap, str::FromStr};

pub fn historian_hysteria(input: &str) -> i32 {
    let (mut left, mut right) = parse_input(input);

    left.sort();
    right.sort();

    let mut distance = 0;
    for i in 0..left.len() {
        let diff = left[i] - right[i];
        distance += diff.abs();
    }

    distance
}

pub fn historian_hysteria_2(input: &str) -> i32 {
    let (left, right) = parse_input(input);

    let mut freq = HashMap::new();
    for num in right.iter() {
        if let Some(count) = freq.get(num) {
            freq.insert(num, count + 1);
        } else {
            freq.insert(num, 1);
        }
    }

    let mut distance = 0;
    for num in left.iter() {
        if let Some(count) = freq.get(num) {
            distance += num * count;
        }
    }

    distance
}

fn parse_input(input: &str) -> (Vec<i32>, Vec<i32>) {
    let mut left = Vec::new();
    let mut right = Vec::new();

    for line in input.lines() {
        if let Some(t) = line.split_once("   ") {
            let l: i32 = FromStr::from_str(t.0).unwrap();
            let r: i32 = FromStr::from_str(t.1).unwrap();

            left.push(l);
            right.push(r);
        }
    }

    (left, right)
}
