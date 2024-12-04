pub fn ceres_search(input: &str) -> i32 {
    let board = parse_input(input);
    let mut total = 0;

    for i in 0..board.len() {
        for j in 0..board[0].len() {
            if board[i][j] != 'X' {
                continue;
            }

            total += get_xmas(&board, i, j);
        }
    }

    total
}

pub fn ceres_search_2(input: &str) -> i32 {
    let board = parse_input(input);
    let mut total = 0;

    for i in 0..board.len() {
        for j in 0..board[0].len() {
            if board[i][j] != 'A' {
                continue;
            }

            if check_x_mas(&board, i, j) {
                total += 1;
            }
        }
    }

    total
}

fn get_xmas(board: &Vec<Vec<char>>, r: usize, c: usize) -> i32 {
    let chars = ['M', 'A', 'S'];
    let mut count = 0;
    let directions = [
        (-1, 0),
        (1, 0),
        (0, -1),
        (0, 1),
        (-1, -1),
        (-1, 1),
        (1, -1),
        (1, 1),
    ];

    for &(dx, dy) in &directions {
        let mut x = r as isize;
        let mut y = c as isize;
        let mut matched = true;

        for i in 0..chars.len() {
            x += dx;
            y += dy;
            if x < 0 || x >= board.len() as isize || y < 0 || y >= board[0].len() as isize {
                matched = false;
                break;
            }
            if board[x as usize][y as usize] != chars[i] {
                matched = false;
                break;
            }
        }
        if matched {
            count += 1;
        }
    }

    count
}

fn check_x_mas(board: &Vec<Vec<char>>, r: usize, c: usize) -> bool {
    if r == 0 || c == 0 || r + 1 >= board.len() || c + 1 >= board[0].len() {
        return false;
    }

    let dig1_valid = board[r - 1][c - 1] != 'X'
        && board[r - 1][c - 1] != 'A'
        && board[r + 1][c + 1] != 'X'
        && board[r + 1][c + 1] != 'A'
        && board[r + 1][c + 1] != board[r - 1][c - 1];
    let dig2_valid = board[r + 1][c - 1] != 'X'
        && board[r + 1][c - 1] != 'A'
        && board[r - 1][c + 1] != 'X'
        && board[r - 1][c + 1] != 'A'
        && board[r + 1][c - 1] != board[r - 1][c + 1];

    dig1_valid && dig2_valid
}

fn parse_input(input: &str) -> Vec<Vec<char>> {
    let mut board = Vec::new();
    for line in input.lines() {
        let line_vec: Vec<char> = line.chars().collect();
        board.push(line_vec)
    }

    board
}
