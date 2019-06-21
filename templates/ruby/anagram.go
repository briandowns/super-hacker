package ruby

const Anagram = `
# splashinn

# return if string a and b are anagrams
def anagram?(a, b)
  char_count = proc {|hsh, c|
    hash[c] ||= 0
    hsh[c] += 1
    hsh
  }
  if a.each_char.reduce({}, &char_count) == b.each_char.reduce({}, &char_count)
      true
    else
      false
    end
end


def anagram_detection(parent, child)
  x = parent.length - child.length
  (0...x).reduce(0) { |c, i| anagram?(parent[i, child.length], child) ? c + 1 : c }
end
`
