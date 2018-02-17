module BookKeeping
  VERSION = 3
end

class Hamming
  ERROR_MSG = 'DNA strands need to be the same length!'.freeze

  def self.compute(orig_strand, new_strand)
    validate(orig_strand, new_strand)
    orig_strand.length.times.count {|i| orig_strand[i] != new_strand[i]}
  end

  def self.validate(orig_strand, new_strand)
    return 0 if orig_strand.empty? || new_strand.empty?
    raise ArgumentError, ERROR_MSG if orig_strand.length != new_strand.length
  end
end
