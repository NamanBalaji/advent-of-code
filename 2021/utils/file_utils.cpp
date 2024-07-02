#include "file_utils.h"
#include <fstream>
#include <stdexcept>

std::vector<std::string> read_lines(const std::string &filename) {
    std::ifstream file(filename);
    if (!file.is_open()) {
        throw std::runtime_error("Could not open file: " + filename);
    }
    std::vector<std::string> lines;
    std::string line;
    while (std::getline(file, line)) {
        lines.push_back(line);
    }
    file.close();
    return lines;
}

std::vector<int> get_nums_vector(const std::string& nums_string) {
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