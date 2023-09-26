#include <ctype.h>
#include <stdio.h>
int lower(int c) {
  if (c >= 'A' && c <= 'Z')
    return c + 'a' - 'A';
  else
    return c;
}
int main() {
  int e = lower('A');

  char esc = e;
  char res = tolower(res);
  printf("%c\n", esc);

  int n = n & 0177;
  printf("%d\n", n);
}


// Path: fundamentals.c
