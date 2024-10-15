//// Summation of primes
//
// The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
//
// Find the sum of all the primes below two million.

#include<stdio.h>

int isprime(int n) {
  int d;

  for (d = 2; d < n; d++)
    if (n % d == 0)
      return 0;

  return 1;
}

long long int sum_of_prime(int n) {
  long long int sum = 0;
  int i;

  for (i = 2; i < n; i++) {
    if (isprime(i)) {
      printf("%i\n",i);
      sum += i;
    }
  }
  return sum;
}

int main() {
  printf("%lld\n", sum_of_prime(2000000));
  return 0;
}

