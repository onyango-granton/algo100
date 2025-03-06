def tower_of_hanoi(n, source, auxiliary, destination):
    if n == 0:
        return
    
    # auxiliary somewhat of the waiting,bay
    tower_of_hanoi(n-1, source, destination, auxiliary)
    print("Move", n, "from ",source, "to" ,destination)
    tower_of_hanoi(n-1, auxiliary, source, destination)

tower_of_hanoi(3,"A","B","C")