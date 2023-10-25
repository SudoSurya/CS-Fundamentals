#include "stdio.h"

int main(){

    int x = 10;
    int z = 11;
    int *p = &x;
    *p = z;
    printf("%p\n", &z);
    printf("%d\n", *p);
    return 0;
}
