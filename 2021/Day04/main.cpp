#include <iostream>
#include <vector>
#include <string>
#include "file_utils.h"

// Function to parse the numbers from the input string
std::vector<int> get_numbers(const std::string& nums_string) {
    std::vector<int> nums;
    std::string num_string;
    for (char ch : nums_string) {
        if (ch != ',') {
            num_string += ch;
        } else {
            nums.push_back(std::stoi(num_string));
            num_string = "";
        }
    }
    if (!num_string.empty()) {
        nums.push_back(std::stoi(num_string));
    }
    return nums;
}

// Function to parse the boards from the input lines
std::vector<std::vector<std::vector<int>>> get_boards(const std::vector<std::string>& input) {
    std::vector<std::vector<std::vector<int>>> boards;
    int i = 2;

    while (i < input.size()) {
        std::vector<std::vector<int>> board;
        for (int j = i; j < i + 5; j++) {
            std::vector<int> row;
            std::string num_string;
            for (char ch : input[j]) {
                if (ch != ' ') {
                    num_string += ch;
                } else if (!num_string.empty()) {
                    row.push_back(std::stoi(num_string));
                    num_string = "";
                }
            }
            if (!num_string.empty()) {
                row.push_back(std::stoi(num_string));
            }
            board.push_back(row);
        }
        boards.push_back(board);
        i += 6;
    }

    return boards;
}


void find_and_mark(int num, std::vector<std::vector<int>>& grid, std::vector<std::vector<bool>>& mark) {
    for (int i = 0; i < grid.size(); i++) {
        for (int j = 0; j < grid[i].size(); j++) {
            if (grid[i][j] == num) {
                mark[i][j] = true;
            }
        }
    }
}


bool check_win(const std::vector<std::vector<bool>>& mark) {
    for (int i = 0; i < mark.size(); i++) {
        bool won = true;
        for (int j = 0; j < mark[i].size(); j++) {
            if (!mark[i][j]) {
                won = false;
                break;
            }
        }
        if (won) {
            return true;
        }
    }

    for (int i = 0; i < mark[0].size(); i++) {
        bool won = true;
        for (int j = 0; j < mark.size(); j++) {
            if (!mark[j][i]) {
                won = false;
                break;
            }
        }
        if (won) {
            return true;
        }
    }

    return false;
}

std::pair<int, int> play_game(const std::vector<int>& nums, std::vector<std::vector<int>>& grid, std::vector<std::vector<bool>>& mark) {
    for (int draw = 0; draw < nums.size(); ++draw) {
        int num = nums[draw];
        find_and_mark(num, grid, mark);
        if (check_win(mark)) {
            return {draw + 1, num};
        }
    }

    return {nums.size(), -1};
}

int calculate_score(const std::vector<std::vector<int>>& grid, const std::vector<std::vector<bool>>& mark) {
    int board_score = 0;
    for (int i = 0; i < mark.size(); i++) {
        for (int j = 0; j < mark[i].size(); j++) {
            if (!mark[i][j]) {
                board_score += grid[i][j];
            }
        }
    }

    return board_score;
}

int main() {
    std::vector<std::string> input = read_lines("../Day04/input.txt");

    std::vector<int> nums = get_numbers(input[0]);
    std::vector<std::vector<std::vector<int>>> boards = get_boards(input);

    int min_turns = 100000;
    int min_idx = -1;
    int last_num_min = -1;

    int max_turns = -1;
    int max_idx = -1;
    int last_num_max = 1;

    for (int i = 0; i < boards.size(); i++) {
        std::vector<std::vector<bool>> mark(5, std::vector<bool>(5, false));
        auto [turns, num] = play_game(nums, boards[i], mark);
        if (turns < min_turns) {
            min_idx = i;
            min_turns = turns;
            last_num_min = num;
        }

        if(turns> max_turns) {
            max_idx = i;
            max_turns = turns;
            last_num_max = num;
        }
    }

// part 1
    if (min_idx != -1) {
        std::vector<std::vector<bool>> mark(5, std::vector<bool>(5, false));
        play_game(nums, boards[min_idx], mark);
        int score = calculate_score(boards[min_idx], mark);
        std::cout << last_num_min * score << std::endl;
    }

// part 2
    if (max_idx != -1) {
        std::vector<std::vector<bool>> mark(5, std::vector<bool>(5, false));
        play_game(nums, boards[max_idx], mark);
        int score = calculate_score(boards[max_idx], mark);
        std::cout << last_num_max * score << std::endl;
    }

    return 0;
}
