#include <iostream>
#include <string>
#include "../include/acc_cpp.h"

int ch1_hyphen_main()
{
  std::string name;
  std::cout << "Please enter a name: ";
  std::cin >> name;

  std::string greeting = "Hello, " + name + "!";
  const std::string spaces(greeting.size(), ' ');
  const std::string second = "| " + spaces + " |";
  const std::string first(second.size(), '-');
  std::cout << std::endl;
  std::cout << first << std::endl;
  std::cout << second << std::endl;
  std::cout << "| " << greeting << " |" << std::endl;
  std::cout << second << std::endl;
  std::cout << first << std::endl;
  return 0;
}

int ch1_asterisk_main()
{
  std::string name;
  std::cout << "Please enter a name: ";
  std::cin >> name;

  std::string greeting = "Hello, " + name + "!";
  const std::string spaces(greeting.size(), ' ');
  const std::string second = "* " + spaces + " *";
  const std::string first(second.size(), '*');
  std::cout << std::endl;
  std::cout << first << std::endl;
  std::cout << second << std::endl;
  std::cout << "* " << greeting << " *" << std::endl;
  std::cout << second << std::endl;
  std::cout << first << std::endl;
  return 0;
}
