#include <stdio.h>
#include <string.h>
#include <uchar.h>

#define MAX 100
#define MIN "Hello World!"

enum boolean { NO, YES };

int main() {
  int a;
  a = MAX;
  printf("%d\n", a);
  char s[] = MIN;
  int len = strlen(s);

  printf("%d\n", len);
}

int strlen1(char s[]) {
  int i;
  while (s[i] != '\0')
    ++i;
  return i;
}
