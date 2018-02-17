module BookKeeping
  VERSION = 1
end

class Sieve
  def initialize(limit)
    @limit = limit
  end

  def primes
    numbers = (2..@limit).to_a
    numbers.each do |i|
      numbers.reject! { |num| num % i == 0 }
      numbers.insert(0, i)
    end
    numbers.reverse
  end
end
