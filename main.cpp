#include "include/array.hpp"

int main()
{
  // Create an instance of the Array class
  Array array;

  // Test appending elements
  array.append(10);
  array.append(20);
  array.append(30);
  array.append(40);
  array.append(50);

  // Print the original array
  std::cout << "Original array: ";
  array.printArray();

  // Test linear search for an existing element
  int searchValue = 30;
  int index = array.linearSearch(searchValue);
  if (index != -1)
  {
    std::cout << "Value " << searchValue << " found at index " << index << std::endl;
  }
  else
  {
    std::cout << "Value " << searchValue << " not found in the array" << std::endl;
  }

  // Test linear search for a non-existing element
  searchValue = 100;
  index = array.linearSearch(searchValue);
  if (index != -1)
  {
    std::cout << "Value " << searchValue << " found at index " << index << std::endl;
  }
  else
  {
    std::cout << "Value " << searchValue << " not found in the array" << std::endl;
  }

  // Test inserting an element
  int newElement = 25;
  int insertIndex = 2;
  array.insertAt(insertIndex, newElement);
  std::cout << "Array after inserting " << newElement << " at index " << insertIndex << ": ";
  array.printArray();

  // Test deleting an element
  int deleteIndex = 3;
  array.deleteAt(deleteIndex);
  std::cout << "Array after deleting element at index " << deleteIndex << ": ";
  array.printArray();

  // Test deleting an element at an invalid index
  deleteIndex = 10;
  array.deleteAt(deleteIndex); // Should print an error message

  return 0;
}
