#include "stdio.h"
#include <stdlib.h>

void swap(int *a, int *b) {
  int temp = *a;
  *a = *b;
  *b = temp;
}

int main() {
  int x = 10;
  int z = 11;
  int *p = &x;
  *p = z;
  printf("%p\n", &z);
  printf("%d\n", *p);
  int i = atoi("jjjk");
  int j = atof("1.2");
  printf("%d atoi\n", i);

  printf("%c\n", j);
  int arr[2];
  int *pa = &arr[0];
  int *pb = &arr[1];
  printf("%d\n", *(pa));
  printf("%d\n", *(pa+1));
  return 0;
}
