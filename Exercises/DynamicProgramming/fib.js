function fib(n, memo = {}) {
    if (n in memo) return memo[n];
    if (n <= 2) return 1;
    memo[n] = fib(n - 1, memo) + fib(n - 2, memo);
    return memo[n];
}
function fibwithoutmemo(n) {
    if (n <= 2) return 1;
    return fibwithoutmemo(n - 1) + fibwithoutmemo(n - 2);
}

// Path: fib.js
console.log(fib(6)); // 8
console.log(fib(7)); // 13
console.log(fib(8)); // 21
console.log(fib(50)); // 12586269025

console.log(fibwithoutmemo(6)); // 8
console.log(fibwithoutmemo(7)); // 13
console.log(fibwithoutmemo(8)); // 21
console.log(fibwithoutmemo(50)); // 12586269025

