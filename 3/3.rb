def prime_factors(n)
  i, lst = 2, []

  while (i*i <= n) do
    if ( n%i == 0)
      lst << i
      n /= i
    else
      i += 1
    end
  end

  lst<<n
end

puts (prime_factors(600851475143)).max

