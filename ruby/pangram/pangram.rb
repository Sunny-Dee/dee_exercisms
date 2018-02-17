module BookKeeping
  VERSION = 5
end

class Pangram
  def self.pangram?(phrase)
    ('a'..'z').all? do |letter|
      phrase.downcase.include?(letter)
    end
  end
end

