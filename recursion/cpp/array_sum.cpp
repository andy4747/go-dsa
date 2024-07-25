#include <stdio.h>
#include "../../include/recursion.h"

// Returns sum of an array using for loop.
int sum_with_loop(int *arr, size_t n)
{
  int sum = 0;
  for (int i = 0; i < n; i++)
  {
    sum += arr[i];
  }
  return sum;
}

// Returns the sum of an array using recursion.
int sum_with_recursion(int *arr, size_t n)
{
  // base case, if size of array is 0 then return 0 as sum;
  if (n == 0)
  {
    return 0;
  }

  // recursive case, add the first element with the sum of remaining elements.
  return arr[0] + sum_with_recursion(arr + 1, n - 1);
}

int sum_main()
{
  int arr[] = {643, 54, 6, 76, 768, 4435, 677, 8768, 4354, 8797};
  size_t size = sizeof(arr) / sizeof(int);
  int sum_loop_value = sum_with_loop(arr, size);
  int sum_recursion_value = sum_with_recursion(arr, size);
  printf("%d\n", sum_loop_value);
  printf("%d\n", sum_recursion_value);
  return 0;
}
