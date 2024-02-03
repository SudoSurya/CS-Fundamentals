#include <stdio.h>

typedef char *string;

int main(void) {
  int n = 50;
  int *p = &n;
  printf("%p\n", p);
  string s = "Hi!";
  printf("%p\n", &s[0]);
  printf("%c\n", *s);
  printf("%c\n", *(s+1));
  printf("%c\n", *(s+2));
  printf("%c\n", *(s+3));
  printf("%c\n", *(s+4));
}
