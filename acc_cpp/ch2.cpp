#include <iostream>
#include <string>
#include "../include/acc_cpp.h"

int ch2_main()
{
  std::string name;
  std::cout << "Please enter a name: ";
  std::cin >> name;

  std::string greeting = "Hello, " + name + "!";

  const int padding = 1;
  const int rows = padding * 2 + 3;

  return 0;
}
