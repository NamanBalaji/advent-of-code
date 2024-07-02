#include "file_utils.h"
#include <string>
#include <vector>
#include <iostream>
#include <array>

long long simulate_days(int days, std::vector<int> interval_timers) {
    std::array<long long, 9> fish_timers = {0};
    for (int timer : interval_timers) {
        fish_timers[timer]++;
    }

    // Simulate each day
    for (int day = 0; day < days; ++day) {
        long long new_fish = fish_timers[0];
        for (int i = 0; i < 8; ++i) {
            fish_timers[i] = fish_timers[i + 1];
        }
        fish_timers[6] += new_fish;
        fish_timers[8] = new_fish;
    }

    // Sum up the total number of fish
    long long total_fish = 0;
    for (long long count : fish_timers) {
        total_fish += count;
    }

    return total_fish;
}

int main() {
    std::vector<std::string> input = read_lines("../Day06/input.txt");
    std::vector<int> interval_timers = get_nums_vector(input[0]);

    std::cout << simulate_days(80, interval_timers) << std::endl;
    std::cout << simulate_days(256, interval_timers) << std::endl;

    return 0;
}
