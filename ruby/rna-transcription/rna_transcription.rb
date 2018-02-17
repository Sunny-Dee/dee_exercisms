module BookKeeping
  VERSION = 4
end

class Complement
  def self.of_dna(dna)
    transcription = dna.dup
    (0...transcription.length).each do |i|
      transcription[i] = case transcription[i]
                         when 'G' then 'C'
                         when 'C' then 'G'
                         when 'T' then 'A'
                         when 'A' then 'U'
                         else return ''
                         end
    end
    transcription
  end
end