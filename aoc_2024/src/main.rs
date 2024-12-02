use std::env;
use std::fs;

mod day1;
mod day2;
mod day3;

fn main() {
    let args: Vec<String> = env::args().collect();

    if args.len() < 2 {
        eprintln!("Usage: advent_of_code <day>");
        std::process::exit(1);
    }

    let day = &args[1];
    let input_path = format!("src/day{}/input.txt", day);

    let input = match fs::read_to_string(&input_path) {
        Ok(content) => content,
        Err(_) => {
            eprintln!("Error: Could not read input file for day {}", day);
            std::process::exit(1);
        }
    };

    match day.as_str() {
        "1" => {
            println!("Day 1, Part 1: {}", day1::historian_hysteria(&input));
            println!("Day 1, Part 2: {}", day1::historian_hysteria_2(&input));
        }
        "2" => {
            println!("Day 2, Part 1: {}", day2::red_nosed_report(&input));
            println!("Day 2, Part 2: {}", day2::red_nosed_report_2(&input));
        }
        "3" => {
            println!("Day 3, Part 1: {}", day3::mull_it_over(&input));
            println!("Day 3, Part 2: {}", day3::mull_it_over_2(&input));
        }
        _ => {
            eprintln!("Error: Day {} is not implemented yet.", day);
            std::process::exit(1);
        }
    }
}
