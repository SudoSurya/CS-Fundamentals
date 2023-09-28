#include <stdio.h>

int len(const int a) { return a; }

int main() {
  int a = 10;
  int res = len(a);
  printf("%d\n", res);
  for (int i = 0; i < 6; i++) {
    printf("%d\n", i);
  }
}

// Path: const.c
