export function printNumber(n: number) {
    if (n === 0) return;
    printNumber(n - 1);
    console.log(n);
}

export function sumofNumber(n: number): number {
    if (n === 0) return 0;
    return n + sumofNumber(n - 1);
}

export function factorial(n: number): number {
    if (n === 0) return 1;
    return n * factorial(n - 1);
}

export function fibonacci(n: number): number {
    if (n === 0 || n === 1) return n;
    return fibonacci(n - 1) + fibonacci(n - 2);
}


// Path: recursion.ts
