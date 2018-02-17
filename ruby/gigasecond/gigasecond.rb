module BookKeeping
  VERSION = 6
end

class Gigasecond
  GIGASECOND = 1_000_000_000


  def self.from(time)
    Time.at(time.to_i + GIGASECOND).utc
  end
end
