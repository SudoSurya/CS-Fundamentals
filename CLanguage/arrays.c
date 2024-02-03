#include "stdio.h"

char *month_name(int n) {
  static char *name[] = {"Illegal month", "January",   "February", "March",
                         "April",         "May",       "June",     "July",
                         "August",        "September", "October",  "November",
                         "December"};
  return (n < 1 || n > 12) ? name[0] : name[n];
}

int strlen1(char *s) {
  int n;
  for (n = 0; *s != '\0'; s++) {
    n++;
  }
  return n;
}

int *pass_by_value(int *x) {
  *x = 10;
  return x;
}

void ere() {
  int arr[5] = {1, 2, 3, 4, 5};
  for (int i = 0; arr[i] != '\0'; i++) {
    printf("arr[%d] = %d\n", i, arr[i]);
  }
}

static char daytab[2][13] = {
    {0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31},
    {0, 31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31},
};

int day_of_year(int year, int month, int day) {
  int i, leap;
  leap = year % 4 == 0 && year % 100 != 0 || year % 400 == 0;
  for (i = 1; i < month; i++) {
    day += daytab[leap][i];
  }
  return day;
}

int main(int argc, char *argv[]) {
  int arr[5];
  int *pa = &arr[0];
  *pa = 0;
  printf("arr[0] = %d\n", arr[0]);
  int n = sizeof(arr) / sizeof(arr[0]);
  printf("n = %d\n", n);
  for (int i = 0; i < sizeof(arr) / sizeof(arr[0]); i++) {
    printf("arr[%d] = %d\n", i, arr[i]);
  }
  int nums = 10111;
  int *ref = pass_by_value(&nums);
  printf("nums = %d\n", nums);
  printf("ref = %d\n", *ref);
  ere();
  printf("----------------------------------------\n");
  while (--argc > 0) {
    printf("%s%s", *++argv, (argc > 1) ? " " : "");
  }
  printf("----------------------------------------\n");
  int aser[5] = {1, 2, 3, 4, 5};
  int *p = &aser[0];
  int les = 5;
  while (--les >= 0) {
    printf("%d%s", *++p, (les > 1) ? " " : "");
  }
}
