Problem Statement: Traveling Salesperson Problem (TSP)
You are given a list of N cities and a distance matrix dist[][], where dist[i][j] represents the distance between city i and city j.

Your task is to find the shortest possible route that visits every city exactly once and returns to the starting city.

Input:
An integer N (number of cities).
A NxN matrix dist where dist[i][j] is the distance between city i and city j.
Output:
The minimum cost of visiting all cities exactly once and returning to the starting city.
Example:
Input:

csharp
Copy
Edit
N = 4  
dist = [  
    [0, 10, 15, 20],  
    [10, 0, 35, 25],  
    [15, 35, 0, 30],  
    [20, 25, 30, 0]  
]  
Output:

Copy
Edit
80
Explanation:
The optimal path is 0 → 1 → 3 → 2 → 0 with a total cost of 80.

Constraints:
1 ≤ N ≤ 10 (Exact solution using recursion/DP)
N > 10 (Approximate solutions using heuristics like Nearest Neighbor, Genetic Algorithm, etc.)