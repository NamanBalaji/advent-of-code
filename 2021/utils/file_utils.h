#ifndef FILE_UTILS_H
#define FILE_UTILS_H

#include <string>
#include <vector>

std::vector<std::string> read_lines(const std::string &filename);
std::vector<int> get_nums_vector(const std::string& nums_string);

#endif // FILE_UTILS_H