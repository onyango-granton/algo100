def decoratorFunc(f):
    def wrapper(num, odd, even):
        result = f(num, odd, even)
        if result:
            even, odd = result

            for x in odd:
                print(x, "Odd")
            
            for y in even:
                print(y, "even")

    return wrapper

@decoratorFunc
def oddeven(num, odd, even): 
    if num > 100:
        return  
    if num % 2 == 0:
        even.append(num)
    else:
        odd.append(num)
    if num == 100:
        return even, odd
    
    return oddeven(num+1, odd, even)

oddeven(0, [], [])