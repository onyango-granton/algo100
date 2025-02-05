def isPalindrome(s):
    left, right = 0, len(s) - 1
    while left < right:
        if not(isAlphanumeric(s[left])):
            left+=1
            continue
        if not isAlphanumeric(s[right]):
            right-=1
            continue
        if not isSame(s[left], s[right]):
            return False
        left += 1
        right -= 1
    return True


def isAlphanumeric(char):
    return char >= 'a' and char <= 'z' or char >= 'A' and char <= 'Z' or char >= '0' and char <= '9' 

def isSame(char1, char2):
    return char1.lower() == char2.lower()

print(isPalindrome("0racecar0....."))