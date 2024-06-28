#include <iostream>
#include <vector>
#include <sstream>
#include "file_utils.h"

struct Position {
    Position(int depth = 0, int horizontal = 0, int aim = 0)
        : depth_(depth), horizontal_(horizontal), aim_(aim) {}

    void addDepth(int d) { depth_ += d; }
    void addHorizontal(int h) { horizontal_ += h; }
    void addAim(int a) { aim_ += a; }
    int getDepth() const { return depth_; }
    int getHorizontal() const { return horizontal_; }
    int getAim() const { return aim_; }

private:
    int depth_;
    int horizontal_;
    int aim_;
};

void executeInstruction(Position& pos, const std::string& instruction, int magnitude, bool with_aim) {
    if (instruction == "forward") {
        pos.addHorizontal(magnitude);
        if(with_aim) {
            pos.addDepth(pos.getAim() * magnitude);
        }
    } else if (instruction == "down") {
        if(with_aim){
            pos.addAim(magnitude);
            return;
        }
        pos.addDepth(magnitude);
    } else if (instruction == "up") {
        if(with_aim) {
            pos.addAim(-magnitude);
            return;
        }
        pos.addDepth(-magnitude);
    } else {
        std::cerr << "Unknown instruction: " << instruction << std::endl;
    }
}

void parse_instruction(std::string input, std::vector<std::string>& vec) {
    std::stringstream ss(input);
    std::string token; 
    char delimiter = ' '; 
  
    while (getline(ss, token, delimiter)) { 
        vec.push_back(token); 
    } 
}

int main() {
    std::vector<std::string> input = read_lines("../Day02/input.txt");
    Position pos1;
    Position pos2;

// part 1
    for(int i = 0; i<input.size(); i++) {
        std::vector<std::string> instructions;
        parse_instruction(input[i], instructions);
        executeInstruction(pos1, instructions[0], std::stoi(instructions[1]), false);
    }

// part 2
    for(int i = 0; i<input.size(); i++) {
        std::vector<std::string> instructions;
        parse_instruction(input[i], instructions);
        executeInstruction(pos2, instructions[0], std::stoi(instructions[1]), true);
    }
    
    std::cout << "Final position: " << pos1.getHorizontal() * pos1.getDepth() << std::endl;
    std::cout << "Final position: " << pos2.getHorizontal() * pos2.getDepth() << std::endl;

    return 0;
}
