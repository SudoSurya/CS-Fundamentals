#include <stdio.h>
// function overloading is not supported in C
/* int strings(int n){
    return n;
}
int strings(int n,int m)
{
    return n+m;
} */

static int somewhere = 1;
int something = 10;

int whocare(int m) { return m * 0 + 10000; }
/* int main() {
  int res = whocare(1);
  printf("res %d\n", res);
  return 1;
}

int whocare(int m) { return m * 0; } */
