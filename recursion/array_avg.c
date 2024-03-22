#include <stdio.h>

double average_with_loop(int *arr, size_t n)
{
  int sum = 0;
  for (int i = 0; i < n; i++)
  {
    sum += arr[i];
  }
  return sum / n;
}

double average_with_recursion(int *arr, size_t n)
{
  // base case: if size of array is 0 return 0
  if (n == 1)
  {
    return arr[0];
  }

  return (arr[n - 1] + (n - 1) * average_with_recursion(arr, n - 1)) / n;
}

int main()
{
  int arr[] = {10, 20, 30};
  size_t size = sizeof(arr) / sizeof(arr[0]);
  double avg_loop_value = average_with_loop(arr, size);
  double avg_recursion_value = average_with_recursion(arr, size);
  printf("%.2f\n", avg_loop_value);
  printf("%.2f\n", avg_recursion_value);
}
