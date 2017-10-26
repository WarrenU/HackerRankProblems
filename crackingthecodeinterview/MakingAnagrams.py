from collections import Counter

def number_needed(a, b):
    num_deletions = 0
    letters_in_a = Counter(a)
    letters_in_b = Counter(b)
    letters_in_a.subtract(letters_in_b)
    return sum(abs(l) for l in letters_in_a.values())
            
a = input().strip()
b = input().strip()

print(number_needed(a, b))
