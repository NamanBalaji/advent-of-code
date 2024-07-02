#include "file_utils.h"
#include <string>
#include <vector>
#include <iostream>

int get_fuel(int diff) {
    int fuel = 0;
    for(int i =1; i<=diff; ++i) {
        fuel += i;
    }
    return fuel;
}

int main () {
    std::vector<std::string> input = read_lines("../Day07/input.txt");
    std::vector<int> positions = get_nums_vector(input[0]);

// part 1
    int min = 1000000;
    for (int i=0; i<positions.size(); i++) {
        int fuel = 0;
        for(int j=0; j<positions.size(); j++) {
            fuel += std::abs(positions[i] - positions[j]);
        }
        if(fuel < min) {
            min = fuel;
        }
    }

    std::cout<<min<<std::endl;

// part 2
    int min_2 = 100000000;
    int max_pos = 0;
    for (int i=0; i<positions.size(); i++){
        if(positions[i] > max_pos) {
            max_pos = positions[i];
        }
    }
    for (int i=0; i<=max_pos; i++) {
        int fuel = 0;
        for(int j=0; j<positions.size(); j++) {
            fuel += get_fuel(std::abs(i - positions[j]));
        }
        if(fuel < min_2) {
            min_2 = fuel;
        }
    }
    
    std::cout<<min_2<<std::endl;

    return 0;
}