#ifndef ARRAY_HPP
#define ARRAY_HPP

#include <iostream>

class Array
{
private:
  int *arr;   // Pointer to the array
  int size;   // Maximum size of the array
  int length; // Current length of the array

public:
  Array();
  ~Array();

  int getAt(int index);
  void append(int value);
  void insertAt(int index, int value);
  int linearSearch(int value);
  void deleteAt(int index);
  void printArray();
};

#endif // ARRAY_HPP
