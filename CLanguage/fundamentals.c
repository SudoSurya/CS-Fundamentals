#include <ctype.h>
#include <stdio.h>
int lower(int c) {
  if (c >= 'A' && c <= 'Z')
    return c + 'a' - 'A';
  else
    return c;
}
int power(int base, int n) {
  int p;
  for (p = 1; n > 0; --n)
    p = p * base;
  return p;
}
int main() {
  int e = lower('A');

  char esc = e;
  char res = tolower(res);

  int n = n & 0177;
  int a = 10;
  // printf("%d+%d", a, power(2, 3));

  int x = 10;
  int y = 20;

  int z = (x > y) ? x : y;
  // printf("biggest val is %d\n", z);
  // printf("biggest val is %d\n", z);
  // printf("biggest val is %d\n", z);
  for (int i = 1; i <= 100; i++) {
    printf("%5d%c", i, (i % 10 == 0 || i == n - 1) ? '\n' : ' ');
  }
}

// Path: fundamentals.c
