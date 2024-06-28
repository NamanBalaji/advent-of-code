#include <iostream>
#include <vector>
#include "file_utils.h"

int get_current_window_sum(std::vector<int>& sonar, int start_index, int window) {
        int sum = 0;
        for (int i = start_index; i<start_index + window; i++) {
            sum += sonar[i];
        }
        
        return sum;
}

int get_increases(std::vector<int>& sonar, int window) {
    int increased = 0;
    int prev = get_current_window_sum(sonar, 0, window);
    for(int i = 1; i <= sonar.size()-window; i++) {
        int curr = get_current_window_sum(sonar, i, window);
        if(curr > prev) {
            ++increased;
        }
        prev = curr;
    }
    
    return increased;
}



int main() {
    std::vector<std::string> input = read_lines("../Day01/input.txt");

    std::vector<int> sonar;
    for(int i = 0; i<input.size(); i++) {
        sonar.push_back(std::stoi(input[i]));
    }
    
    // part 1
    std::cout << get_increases(sonar, 1) << std::endl;
    
    //part 2
    std::cout << get_increases(sonar, 3) << std::endl;

    return 0;
}                                                                                                                                                           