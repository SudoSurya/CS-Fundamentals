#include <stdio.h>
#include <sys/ucontext.h>
// function overloading is not supported in C
/* int strings(int n){
    return n;
}
int strings(int n,int m)
{
    return n+m;
} */

int whocare(int m);
int main() {
  int res = whocare(1);
  printf("res %d\n", res);
  return 1;
}

int whocare(int m) { return m * 0; }
