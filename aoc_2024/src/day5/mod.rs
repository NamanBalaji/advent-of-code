use std::{collections::HashMap, str::FromStr};

fn parse_input(input: &str) -> (Vec<(i32, i32)>, Vec<Vec<i32>>) {
    let (ordering_str, updates_str) = input.split_once("\n\n").unwrap();

    let mut ordering: Vec<(i32, i32)> = Vec::new();
    for lines in ordering_str.lines() {
        let (a, b) = lines.split_once('|').unwrap();
        ordering.push((FromStr::from_str(a).unwrap(), FromStr::from_str(b).unwrap()));
    }

    let mut updates = Vec::new();
    for line in updates_str.lines() {
        let update: Vec<i32> = line
            .split(',')
            .map(|s| FromStr::from_str(s).unwrap())
            .collect();
        updates.push(update);
    }

    (ordering, updates)
}

pub fn print_queue(input: &str) -> i32 {
    let (rules, updates) = parse_input(input);
    let mut sum = 0;
    for update in updates {
        if check_valid(&update, &rules) {
            sum += update[update.len() / 2];
        }
    }

    sum
}

pub fn print_queue_2(input: &str) -> i32 {
    let (rules, mut updates) = parse_input(input);

    let mut sum = 0;

    for update in &mut updates {
        if !check_valid(&update, &rules) {
            sort_update(update, &rules);
            sum += update[update.len() / 2];
        }
    }

    sum
}

fn check_valid(update: &Vec<i32>, rules: &Vec<(i32, i32)>) -> bool {
    let mut idx_map = HashMap::new();
    for i in 0..update.len() {
        idx_map.insert(update[i], i);
    }

    for rule in rules {
        let idx_prev = idx_map.get(&rule.0);
        let idx_later = idx_map.get(&rule.1);

        if idx_prev.is_none() || idx_later.is_none() {
            continue;
        }

        if idx_later.unwrap() <= idx_prev.unwrap() {
            return false;
        }
    }

    return true;
}

fn sort_update(update: &mut Vec<i32>, rules: &Vec<(i32, i32)>) {
    loop {
        let mut is_sorted = true;
        for i in 0..update.len() - 1 {
            let pair = (update[i], update[i + 1]);
            if rules.contains(&pair) {
                update.swap(i, i + 1);
                is_sorted = false;
            }
        }

        if is_sorted {
            break;
        }
    }
}
