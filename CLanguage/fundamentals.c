#include <ctype.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
void reverse(char s[]) {
  int c, i, j;
  for (i = 0, j = strlen(s) - 1; i < j; i++, j--) {
    c = s[i];
    s[i] = s[j];
    s[j] = c;
  }
}
int lower(int c) {
  if (c >= 'A' && c <= 'Z')
    return c + 'a' - 'A';
  else
    return c;
}
int lowerExp(int c) { return c >= 'A' && c <= 'Z' ? c + 'a' - 'A' : c; }
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
  printf("lower - EXP%6c\n", lowerExp('A'));
  printf("lower %6c\n", lower('A'));
  // int c = getchar();
  // printf("getchar %6c\n", c);
  char c;
  /* while (c = getchar()) {
    printf("getchar %6c\n", c);
    if (c == 'q')
      break;
  } */

  char s[] = "hello";
  int len = strlen(s);
  printf("len of s %d\n", len);
  printf("size of s %ld\n", sizeof(s));
  int size = sizeof(s);
  // for (int i = 0; i < len; i++) {
  //   printf("%c", s[i]);
  // }
  // printf("\n");
  printf("before reverse - %6s\n", s);
  reverse(s);
  printf("after reverse - %6s\n", s);

  // printf("\n");
  // for (int i = 0; i < len; i++) {
  //   printf("%c", s[i]);
  // }
}
// Path: fundamentals.c
