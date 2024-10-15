## Summation of primes

# The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

# Find the sum of all the primes below two million.

def is_prime?(n)
  if n < 2
    return false
  elsif n >= 2
    (2..(n-1)).each do |i|
      if n % i == 0
        return false
      end
    end
  end
  return true
end

def sum_of_prime(n)
  sum = 0
  (2...n).each do |i|
    if is_prime?(i)
      puts i
      sum += i
    end
  end
  sum
end

puts sum_of_prime(200_0000)
