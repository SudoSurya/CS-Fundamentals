#include <stdio.h>


int len(const int a){
    return a;
}

int main(){
    int a = 10;
    int res = len(a);
    printf("%d\n", res);
}

// Path: const.c


