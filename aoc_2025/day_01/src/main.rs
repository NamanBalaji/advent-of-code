use std::{fs, str::FromStr};

struct Rotation {
    multiplier: i8,
    times: i32,
}

impl FromStr for Rotation {
    type Err = String;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let (dir_char, rest) = s.split_at(1);
        let multiplier = match dir_char.chars().next() {
            Some('L') => -1,
            Some('R') => 1,
            _ => return Err("First char must be L or R".to_string()),
        };

        let times: i32 = rest.parse().map_err(|_| "Invalid number".to_string())?;

        Ok(Rotation { multiplier, times })
    }
}

fn main() {
    let rotations = match parse() {
        Ok(r) => r,
        Err(e) => panic!("{e}"),
    };

    let start_pos = 50;
    let zero_count = rotations
        .iter()
        .scan(start_pos, |pos, r| {
            let delta = r.multiplier as i32 * r.times;
            *pos = (*pos + delta).rem_euclid(100);

            Some(*pos)
        })
        .filter(|&p| p == 0)
        .count();

    println!("{zero_count}");
    println!("{}", part2(&rotations))
}

fn parse() -> Result<Vec<Rotation>, String> {
    let input = match fs::read_to_string("./src/input.txt") {
        Ok(i) => i,
        Err(e) => return Err(format!("Error ocurred while trying to read input {e}")),
    };

    input.lines().map(|line| line.parse::<Rotation>()).collect()
}

fn part2(rotations: &Vec<Rotation>) -> i32 {
    let mut current = 50;
    let mut times_zero = 0;
    for rotation in rotations {
        let mut times = rotation.times;
        while times > 0 {
            current += rotation.multiplier as i32;

            if current > 99 {
                current = 0;
            }

            if current < 0 {
                current = 99;
            }

            if current == 0 {
                times_zero += 1;
            }

            times -= 1;
        }
    }

    times_zero
}
