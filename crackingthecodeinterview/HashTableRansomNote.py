from collections import Counter

def ransom_note(magazine, ransom):
    mag_words = Counter(magazine)
    ransom_words = Counter(ransom)
    for word in ransom_words:
        if mag_words[word] < ransom_words[word]:
            return False
    return True
    

m, n = map(int, input().strip().split(' '))
magazine = input().strip().split(' ')
ransom = input().strip().split(' ')
answer = ransom_note(magazine, ransom)
if(answer):
    print("Yes")
else:
    print("No")
    
