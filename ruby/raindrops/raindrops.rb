module BookKeeping
  VERSION = 3
end

class Raindrops
  def self.convert(input)
    factor(input)
    result = ''
    result << 'Pling' if @factors.include?(3)
    result << 'Plang' if @factors.include?(5)
    result << 'Plong' if @factors.include?(7)
    result = input.to_s if result.empty?
    result
  end

  private

  def self.factor(input)
    @factors = []
    (1..input).count do |i|
      @factors << i if (input % i).zero?
    end
    @factors
  end
end
