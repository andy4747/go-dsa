#include <iostream>
#include "../include/array.hpp"

Array::Array()
{
  this->size = 20;
  this->length = 0;
  this->arr = new int[size];
}

Array::~Array()
{
  delete[] arr;
}

int Array::getAt(int index)
{
  return this->arr[index];
}

void Array::append(int value)
{
  // Check if there is space available in the array
  if (this->length < this->size)
  {
    this->arr[this->length] = value; // Append the new value at the end
    this->length++;                  // Increase the length
  }
  else
  {
    std::cout << "Array is full. Cannot append more elements." << std::endl;
  }
}

void Array::insertAt(int index, int value)
{
  // Validate the index
  if (index > this->length || index < 0)
  {
    std::cout << "Invalid Index. Cannot insert element at the specified index." << std::endl;
    return;
  }

  // Check if the array has space for a new element
  if (this->length >= this->size)
  {
    std::cout << "The array is full cannot insert more elements in the array." << std::endl;
    return;
  }

  // Shift elements to the right to make space for the new element
  for (int i = this->length; i > index; i--)
  {
    this->arr[i] = this->arr[i - 1];
  }

  // Insert the new element at the specified index
  this->arr[index] = value;

  // Increment the length of the array
  this->length++;
}

int Array::linearSearch(int value)
{
  for (int i = 0; i < this->length; i++)
  {
    if (this->arr[i] == value)
      return i;
  }
  return -1;
}

void Array::deleteAt(int index)
{
  // check if the index is valid
  if (index > this->length || index < 0)
  {
    std::cout << "Invalid Index. Cannot delete element at the specified index." << std::endl;
    return;
  }

  for (int i = index; i < this->length - 1; i++)
  {
    this->arr[i] = this->arr[i + 1];
  }

  this->length--;
}

void Array::printArray()
{
  for (int i = 0; i < length; i++)
  {
    std::cout << arr[i] << " ";
  }
  std::cout << std::endl;
}
