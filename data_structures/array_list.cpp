#include "../include/array_list.h"
#include <stdio.h>

ArrayList::ArrayList()
{
  this->size = 10;
  this->length = 0;
  this->A = new int[this->size];
}

void ArrayList::display()
{
  printf("{ ");
  for (int i = 0; i < this->length; i++)
  {
    if ((i + 1) == length)
    { // if last length of array
      printf("%d ", this->A[i]);
    }
    else
    {
      printf("%d, ", this->A[i]);
    }
  }
  printf("}\n");
}

void ArrayList::append(int value)
{
  if (this->length < this->size)
    this->A[length++] = value;
}
