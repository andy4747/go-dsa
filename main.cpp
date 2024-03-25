#include <iostream>
#include <stdio.h>
#include "include/array_list.h"

using namespace std;

int main()
{
  ArrayList *a = new ArrayList();
  a->append(10);
  a->append(91);
  a->display();
}
