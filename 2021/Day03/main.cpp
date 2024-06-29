#include "file_utils.h"
#include <iostream>
#include <vector>
#include <string>
#include <cmath>
#include <algorithm>

int binary_to_decimal(const std::vector<int>& binary_num) {
    int decimal_value = 0;
    int size = binary_num.size();

    for (int i = 0; i < size; ++i) {
        int bit_value = binary_num[size - 1 - i];
        decimal_value += bit_value * std::pow(2, i);
    }

    return decimal_value;
}

std::vector<int> decimal_to_binary(int decimal_num) {
    std::vector<int> binary_num;
    if (decimal_num == 0) {
        binary_num.push_back(0);
    } else {
        while (decimal_num > 0) {
            binary_num.push_back(decimal_num % 2);
            decimal_num /= 2;
        }
        std::reverse(binary_num.begin(), binary_num.end());
    }
    return binary_num;
}

std::vector<int> binary_string_to_vector(const std::string& binary_string) {
    std::vector<int> binary_num;
    for (char bit : binary_string) {
        binary_num.push_back(bit - '0');
    }
    return binary_num;
}

int get_most_common_bit(const std::vector<std::string>& numbers, int position) {
    int count = 0;
    for (const auto& num : numbers) {
        if (num[position] == '1') {
            count++;
        }
    }
    return (count >= numbers.size() - count) ? 1 : 0;
}

int get_least_common_bit(const std::vector<std::string>& numbers, int position) {
    int count = 0;
    for (const auto& num : numbers) {
        if (num[position] == '1') {
            count++;
        }
    }
    return (count < numbers.size() - count) ? 1 : 0;
}

std::vector<std::string> filter_numbers(const std::vector<std::string>& numbers, int position, bool most_common) {
    std::vector<std::string> filtered_numbers;
    int bit_criteria = most_common ? get_most_common_bit(numbers, position) : get_least_common_bit(numbers, position);

    for (const auto& num : numbers) {
        if ((num[position] - '0') == bit_criteria) {
            filtered_numbers.push_back(num);
        }
    }
    return filtered_numbers;
}

int find_rating(std::vector<std::string> numbers, bool most_common) {
    int bit_length = numbers[0].size();
    for (int i = 0; i < bit_length; ++i) {
        numbers = filter_numbers(numbers, i, most_common);
        if (numbers.size() == 1) {
            break;
        }
    }
    return binary_to_decimal(binary_string_to_vector(numbers[0]));
}

std::vector<int> get_ones_count(const std::vector<std::string>& binary_numbers) {
    int bit_length = binary_numbers[0].size();
    std::vector<int> ones(bit_length, 0);

    for (const auto& binary_num : binary_numbers) {
        for (int i = 0; i < bit_length; ++i) {
            if (binary_num[i] == '1') {
                ++ones[i];
            }
        }
    }

    return ones;
}

int get_gamma_rate(const std::vector<std::string>& binary_numbers) {
    std::vector<int> ones = get_ones_count(binary_numbers);
    std::vector<int> binary_gamma;

    int total_nums = binary_numbers.size();

    for (int i = 0; i < ones.size(); ++i) {
        if (ones[i] > total_nums / 2) {
            binary_gamma.push_back(1);
        } else {
            binary_gamma.push_back(0);
        }
    }

    return binary_to_decimal(binary_gamma);
}

std::vector<int> ones_complement(const std::vector<int>& binary_num) {
    std::vector<int> complement;
    for (int bit : binary_num) {
        complement.push_back(bit == 0 ? 1 : 0);
    }
    return complement;
}

int get_epsilon_rate(int gamma_rate, int bit_length) {
    std::vector<int> binary_gamma = decimal_to_binary(gamma_rate);

    std::vector<int> binary_epsilon = ones_complement(binary_gamma);
    return binary_to_decimal(binary_epsilon);
}

int main() {
    std::vector<std::string> input = read_lines("../Day03/input.txt");

    int gamma = get_gamma_rate(input);
    int bit_length = input[0].size();
    int epsilon = get_epsilon_rate(gamma, bit_length);

    std::cout << "Gamma: " << gamma << std::endl;
    std::cout << "Epsilon: " << epsilon << std::endl;
    std::cout << "Power consumption: " << gamma * epsilon << std::endl;

    int oxygen_generator_rating = find_rating(input, true);
    int co2_scrubber_rating = find_rating(input, false);

    std::cout << "Oxygen Generator Rating: " << oxygen_generator_rating << std::endl;
    std::cout << "CO2 Scrubber Rating: " << co2_scrubber_rating << std::endl;
    std::cout << "Life Support Rating: " << oxygen_generator_rating * co2_scrubber_rating << std::endl;

    return 0;
}
