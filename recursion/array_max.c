#include <stdio.h>

/* Tracing Tree
findMax({1, 4, 5, 6, 7}, 5)
|
|--> findMax({1, 4, 5, 6}, 4)
|    |
|    |--> findMax({1, 4, 5}, 3)
|    |    |
|    |    |--> findMax({1, 4}, 2)
|    |    |    |
|    |    |    |--> findMax({1}, 1)
|    |    |    |    |
|    |    |    |    '--> return 1
|    |    |    |
|    |    |    '--> return max(4, 1) = 4
|    |    |
|    |    '--> return max(5, 4) = 5
|    |
|    '--> return max(6, 5) = 6
|
'--> return max(7, 6) = 7
*/
int findMax(int arr[], int n)
{
  if (n == 1)
  {
    return arr[0];
  }
  return (arr[n - 1] > findMax(arr, n - 1)) ? arr[n - 1] : findMax(arr, n - 1);
}

int main()
{
  int arr[] = {1, 4, 7, 6, 5};
  int n = sizeof(arr) / sizeof(arr[0]);
  printf("The maximum value in the array is %d\n", findMax(arr, n));
  return 0;
}
