import plotext as plt
import time

def divide(arr):
    if len(arr) < 2:
        return arr
    midPoint = len(arr) // 2

    leftArr = arr[:midPoint]
    rightArr = arr[midPoint:]

    return conquer(divide(leftArr), divide(rightArr))

def conquer(leftArr, rightArr):
    mainArr = list(range(len(leftArr)+len(rightArr)))
    mainP, rightP, leftP = 0,0,0
    while rightP < len(rightArr) and leftP < len(leftArr):
        if rightArr[rightP] <= leftArr[leftP]:
            mainArr[mainP] = rightArr[rightP]
            rightP += 1
        else:
            mainArr[mainP] = leftArr[leftP]
            leftP += 1
        mainP += 1
        visualize_data(get_coordinates(mainArr))
    while rightP < len(rightArr):
        mainArr[mainP] = rightArr[rightP]
        visualize_data(get_coordinates(mainArr))
        mainP += 1
        rightP += 1
    while leftP < len(leftArr):
        mainArr[mainP] = leftArr[leftP]
        visualize_data(get_coordinates(mainArr))
        mainP += 1
        leftP += 1
    visualize_data(get_coordinates(mainArr))
    return mainArr

def get_coordinates(arr):
    x_values = [i for i in range(len(arr))]
    print({"X":x_values, "Y":arr})
    return {"X":x_values, "Y":arr}

def visualize_data(data):
    x_value = data.get("X")
    y_value = data.get("Y")
    plt.clf()
    plt.yticks(y_value)
    plt.bar(x_value, y_value)

    time.sleep(2)

    plt.show()


print(divide([3,6,7,3,2,5,6,57,4,34]))