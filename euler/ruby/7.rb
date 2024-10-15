## 10001st prime

# By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we
# can see that the 6th prime is 13.

# What is the 10 001st prime number?

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

def prime(n)
  i,j = 0,0
  while j < n
    i +=1
    if is_prime?(i)
      j += 1
      puts "#{j}: #{i}"
    end
  end
end

prime(10001)
