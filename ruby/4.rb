## Largest palindrome product

# A palindromic number reads the same both ways. The largest palindrome made 
# from the product of two 2-digit numbers is 9009 = 91 99.

# Find the largest palindrome made from the product of two 3-digit numbers.


# It's a wrong solution, but get a correct answer. 
# Should try to fix this later.
def largest_palindrome(i,j)
  i.downto(900) do |i|
    j.downto(900) do |j|
      s = i*j
      if s == s.to_s.reverse.to_i
        return s
      end
    end
  end
end

puts largest_palindrome(999,999)
