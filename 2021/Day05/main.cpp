#include <iostream>
#include <string>
#include <vector>
#include <array>
#include "file_utils.h"

class Point {
private:
    int x_;
    int y_;

public:
    Point(int x = 0, int y = 0) : x_(x), y_(y) {} 
    Point(const std::string& point_str) {
        std::string x_str = point_str.substr(0, point_str.find(","));
        std::string y_str = point_str.substr(point_str.find(",") + 1);
        x_ = std::stoi(x_str);
        y_ = std::stoi(y_str);
    }

    int getX() const { return x_; }
    int getY() const { return y_; }
};

class Line {
private:
    Point a_;
    Point b_;

public:
    Line(const Point& a, const Point& b) : a_(a), b_(b) {}
    Line(const std::string& line_str) {
        std::string point_a_str = line_str.substr(0, line_str.find(" -> "));
        std::string point_b_str = line_str.substr(line_str.find(" -> ") + 4);
        a_ = Point(point_a_str);
        b_ = Point(point_b_str);
    }

    const Point& getA() const { return a_; }
    const Point& getB() const { return b_; }
};

std::vector<Line> parse_input(const std::vector<std::string>& input) {
    std::vector<Line> lines;
    for (const auto& line_str : input) {
        Line l = Line(line_str);
        lines.push_back(l);
    }
    return lines;
}

int part1(std::vector<Line>& lines, std::vector<std::vector<int>>& points) {
    for (const auto& line : lines) {
        if(line.getA().getX() == line.getB().getX()) {
            int x = line.getA().getX();
            int start_y = std::min(line.getA().getY(), line.getB().getY());
            int end_y = std::max(line.getA().getY(), line.getB().getY());
            for (int y = start_y; y <= end_y; ++y) {
                points[y][x]++;
            }
        } else if (line.getA().getY() == line.getB().getY()) {
            int y = line.getA().getY();
            int start_x = std::min(line.getA().getX(), line.getB().getX());
            int end_x = std::max(line.getA().getX(), line.getB().getX());
            for (int x = start_x; x <= end_x; ++x) {
                points[y][x]++;
            }
        }
    }

    int count = 0;
    for(const auto& row : points){
        for(int val : row) {
            if(val > 1) {
                count++;
            }
        }
    }

    return count;
}

int part2(std::vector<Line>& lines, std::vector<std::vector<int>>& points) {
    for (const auto& line : lines) {
        int x1 = line.getA().getX();
        int y1 = line.getA().getY();
        int x2 = line.getB().getX();
        int y2 = line.getB().getY();
        
        if (abs(x1 - x2) == abs(y1 - y2)) {
            int dx = (x2 > x1) ? 1 : -1;
            int dy = (y2 > y1) ? 1 : -1;
            int x = x1;
            int y = y1;
            while (x != x2 + dx && y != y2 + dy) {
                points[y][x]++;
                x += dx;
                y += dy;
            }
        }
    }

    int count = 0;
    for (const auto& row : points) {
        for (int val : row) {
            if (val > 1) {
                count++;
            }
        }
    }

    return count;
}

int main() {
    std::vector<std::string> input = read_lines("../Day05/input.txt");
    auto lines = parse_input(input);

    int max_x = 0;
    int max_y = 0;
    for (const auto& line : lines) {
        if(line.getA().getX() > max_x) {
            max_x = line.getA().getX();
        }
        if(line.getB().getX() > max_x) {
            max_x = line.getB().getX();
        }

        if(line.getA().getY() > max_y) {
            max_y = line.getA().getY();
        }
        if(line.getB().getY() > max_y) {
            max_y = line.getB().getY();
        }        
    }

    max_x += 1;
    max_y += 1;

    std::vector<std::vector<int>> points(max_y, std::vector<int>(max_x, 0));

    int count1 = part1(lines, points);
    int count2 = part2(lines, points);

    std::cout << count1 << std::endl;
    std::cout << count2 << std::endl;
    
    return 0;
}
