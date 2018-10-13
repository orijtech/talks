#include <stdio.h>
#include <stdlib.h>
#include <pthread.h>

void *th_fib(void *arg);

int fib(int n) {
    if (n < 2)
        return n;

    pthread_t t1, t2;
    int n_1 = n-1, n_2 = n-2;
    pthread_create(&t1, NULL, th_fib, &n_1);
    pthread_create(&t2, NULL, th_fib, &n_2);

    int *x, *y;
    pthread_join(t1, (void **)&x);
    pthread_join(t2, (void **)&y);

    int result = *x + *y;
    free(x);
    free(y);
    return result;
}
