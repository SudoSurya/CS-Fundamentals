#include "stdio.h"
#include <string.h>

int strlen1(char *s) {
  int n;
  for (n = 0; *s != '\0'; s++) {
    n++;
  }
  return n;
}

int pass_by_value(int *x) {
  *x = 10;
  return *x;
}
void ere() {
    int arr[5] = {1, 2, 3, 4, 5};
    for (int i = 0; arr[i] != '\0'; i++) {
        printf("arr[%d] = %d\n", i, arr[i]);
    }
}


int main() {
  int arr[5];
  int *pa = &arr[0];
  *pa = 0;
  printf("arr[0] = %d\n", arr[0]);
  int n = sizeof(arr) / sizeof(arr[0]);
  printf("n = %d\n", n);
  for (int i = 0; i < sizeof(arr) / sizeof(arr[0]); i++) {
    printf("arr[%d] = %d\n", i, arr[i]);
  }
  int nums = 10111;
  int ref = pass_by_value(&nums);
  printf("nums = %d\n", nums);
  printf("ref = %d\n", ref);
  ere();
}

