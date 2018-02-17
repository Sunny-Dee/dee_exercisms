module BookKeeping
  VERSION = 4
end

class Squares
  def initialize(number)
    @range = 1..number
  end

  def square_of_sum
    @range.reduce(:+)**2
  end

  def sum_of_squares
    @range.reduce { |sum, num| sum + num**2 }
  end

  def difference
    square_of_sum - sum_of_squares
  end
end
