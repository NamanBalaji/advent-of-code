use std::collections::HashSet;

struct SimulateResult {
    visited_positions: HashSet<(isize, isize)>,
    loop_detected: bool,
}

fn parse_input(input: &str) -> Vec<Vec<char>> {
    let mut grid = Vec::new();

    for line in input.lines() {
        let row: Vec<char> = line.chars().collect();
        grid.push(row);
    }

    grid
}

pub fn guard_gallivant(input: &str) -> i32 {
    let grid = parse_input(input);

    let rows = grid.len() as isize;
    let cols = grid[0].len() as isize;

    let mut r = 0;
    let mut c = 0;

    for row in 0..rows {
        let mut found = false;
        for col in 0..cols {
            if grid[row as usize][col as usize] == '^' {
                r = row as isize;
                c = col as isize;
                found = true;

                break;
            }
        }
        if found {
            break;
        }
    }

    let res = simulate(&grid, r, c);

    res.visited_positions.len() as i32
}

pub fn guard_gallivant_2(input: &str) -> i32 {
    let grid = parse_input(input);

    let rows = grid.len() as isize;
    let cols = grid[0].len() as isize;

    let mut start_r = 0;
    let mut start_c = 0;

    for row in 0..rows {
        let mut found = false;
        for col in 0..cols {
            if grid[row as usize][col as usize] == '^' {
                start_r = row as isize;
                start_c = col as isize;
                found = true;

                break;
            }
        }
        if found {
            break;
        }
    }

    let mut valid_obstruction_count = 0;
    for r in 0..rows {
        for c in 0..cols {
            if grid[r as usize][c as usize] == '#' || (r == start_r && c == start_c) {
                continue;
            }

            let mut grid_copy = grid.clone();
            grid_copy[r as usize][c as usize] = '#';

            let result = simulate(&grid_copy, start_r, start_c);
            if result.loop_detected {
                valid_obstruction_count += 1;
            }
        }
    }

    valid_obstruction_count
}

fn simulate(grid: &Vec<Vec<char>>, mut r: isize, mut c: isize) -> SimulateResult {
    let rows = grid.len() as isize;
    let cols = grid[0].len() as isize;

    let mut dx = -1;
    let mut dy = 0;

    let mut visited_positions = HashSet::new();
    let mut seen_states = HashSet::new();

    loop {
        visited_positions.insert((r, c));
        seen_states.insert((r, c, dx, dy));

        let next_r = r + dx;
        let next_c = c + dy;
        if next_r < 0 || next_r >= rows || next_c < 0 || next_c >= cols {
            return SimulateResult {
                visited_positions,
                loop_detected: false,
            };
        }
        if grid[next_r as usize][next_c as usize] == '#' {
            let temp = dx;
            dx = dy;
            dy = -1 * temp;
        } else {
            r += dx;
            c += dy;
        }

        if seen_states.contains(&(r, c, dx, dy)) {
            return SimulateResult {
                visited_positions,
                loop_detected: true,
            };
        }
    }
}
