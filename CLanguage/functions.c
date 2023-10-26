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

int whocare(int *m) {
    *m = *m + 100000;
    return *m;
}
int main() {
  int a = 10;
  int res = whocare(&a);
  printf("res %d\n", res);
  printf("a %d\n", a);
  return 1;
}
