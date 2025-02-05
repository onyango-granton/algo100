def isPalindrome(s):
    left, right = 0, len(s) - 1
    while left < right:
        if not s[left].isalnum():
            left+=1
            continue
        if not s[right].isalnum():
            right-=1
            continue
        if s[left].lower() != s[right].lower():
            return False
        left += 1
        right -= 1
    return True


def isAlphanumeric(char):
    return char >= 'a' and char <= 'z' or char >= 'A' and char <= 'Z' or char >= '0' and char <= '9' 

def isSame(char1, char2):
    return char1.lower() == char2.lower()

print(isPalindrome("0race###Car0....."))