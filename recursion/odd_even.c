#include <stdio.h>

int odd_even(int n)
{
  if (n == 0)
  {
    return 1;
  }
  else if (n == 1)
  {
    return 0;
  }
  else
  {
    return odd_even(n - 2);
  }
}

int main()
{
  int a;
  printf("Enter a number: ");
  scanf("%d", &a);
  if (odd_even(a) == 1)
  {
    printf("The number is Even\n");
  }
  else
  {
    printf("The number is Odd\n");
  }
}
