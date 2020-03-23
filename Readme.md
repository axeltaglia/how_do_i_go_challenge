How do I "go"?
==============================

HakerRank Link: https://www.hackerrank.com/leaderboard

### Problem

Great Goofy Gorgeous Nivash is a go-getter with a gift of the gab. Since he cant take gigantic steps, he always wants to gopher(go for) the shortest path for "going" between 2 cities.

Help him find the length of such a path.

### Input:
First line of the input has a single integer T denoting the number of test cases.

Each test case begins with an integer N denoting the number cities.

This is followed by an NxN binary matrix in which if the ith row and jth column denotes if city i and city j are connected by a birectional road.

The next line has 2 integers X and Y denoting the start and end cities.

### Output:
For each test case, print "Case #" followed by

"no"(quotes for clarity) if X and Y are not connected.

"yes L" where L is the length of the shortest path from X to Y.

### Constraints:
1 <= T <= 100

1 <= N <= 100

The nodes are numbered from 0 to N-1

### Sample Test Case:
Sample Input:
```
2 
3 
0 1 1 
1 0 0 
1 0 0
1 2
2
1 0
0 0
0 1
```

Sample Output:
```
Case #1: yes 2
Case #2: no
```