use std::str::FromStr;

pub fn red_nosed_report(input: &str) -> i32 {
    let reports = parse_input(input);
    let mut total_safe = 0;

    for report in reports {
        if is_safe(&report, 0) {
            total_safe += 1;
        }
    }

    total_safe
}

pub fn red_nosed_report_2(input: &str) -> i32 {
    let reports = parse_input(input);
    let mut total_safe = 0;

    for report in reports {
        if is_safe(&report, 1) {
            total_safe += 1;
        }
    }

    total_safe
}

fn is_safe(report: &Vec<i32>, allowed: i32) -> bool {
    let ordering = report[0] > report[1];
    let mut prev = report[0];
    let mut total = 0;

    for i in 1..report.len() {
        if ordering != (prev > report[i]) {
            if total == allowed {
                return false;
            }
            total += 1;
            prev = report[i];

            continue;
        }

        let diff = prev - report[i];
        if diff.abs() < 1 || diff.abs() > 3 {
            if total == allowed {
                return false;
            }
            total += 1;
            prev = report[i];

            continue;
        }

        prev = report[i];
    }

    true
}

fn parse_input(input: &str) -> Vec<Vec<i32>> {
    let mut reports = Vec::new();
    for line in input.lines() {
        let report = line
            .split_whitespace()
            .map(|s| FromStr::from_str(s).unwrap())
            .collect();
        reports.push(report);
    }

    reports
}
