## Smallest multiple

# 2520 is the smallest number that can be divided by each of the numbers from 1 
# to 10 without any remainder.

# What is the smallest positive number that is evenly divisible by all of the 
# numbers from 1 to 20?

def smallest_multiple(n)
  m = 1
  while true
    n.downto(1) do |i|
      if (m % i != 0)
        break
      end
      if (i==1)
        return m
      end
    end
    m += 1
  end
end

# Spent about 1 minutes do the computing.
# There must be a better solution.
puts smallest_multiple(18)



