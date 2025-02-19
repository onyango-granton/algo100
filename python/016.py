def findFirstNonRepeatingChar(word):
    char_Dict = {}
    index = len(word)
    
    for i, char in enumerate(word):
        char_Dict.setdefault(char, []).append(i)

    for key in char_Dict:
        if len(char_Dict[key]) == 1:
            if char_Dict[key][0] < index:
                index = char_Dict[key][0]

    if index >= 0 and index < len(word):
        return index
    else:
        return -1


print(findFirstNonRepeatingChar("ww"))