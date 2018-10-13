void *th_fib(void *arg) {
    int *n = (int *)arg;
    int *result = malloc(sizeof(int));
    *result = fib(*n);
    return (void *)result;
}
