## Special Pythagorean triplet

# Pythagorean triplet is a set of three natural numbers, a < b < c,
# for which,

# a^2 + b^2 = c^2
# For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.

# There exists exactly one Pythagorean triplet for which a + b + c =
# 1000.  Find the product abc.

# a + b + c = 1000
# a*a + b*b  = c*c

def triplet(n)
  (1..n).each do |a|
    (1..n).each do |b|
      (1..n).each do |c|
        if (a<b && b<c)
          if (a+b+c==n)
            if ((a**2 + b**2) == c**2)
              return a*b*c
            end
          end
        end
      end
    end
  end
end

puts triplet(12)
puts triplet(1000)
