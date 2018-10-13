int fib(int n) {
    if (n < 2)
        return n;

    int x = spawn fib(n-1);
    int y = spawn fib(n-2);

    sync;

    return x + y;
}
