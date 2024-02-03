#include <stdio.h>
#include <stdlib.h>

int main(int argc, char *argv[]) {
  printf("%d",argc);
  int *arr = malloc(3 * sizeof(int));
  *(arr + 0) = 0;
  *(arr + 1) = 1;
  *(arr + 2) = 2;
  free(arr);
  FILE *f = fopen("test.txt", "w");
}
