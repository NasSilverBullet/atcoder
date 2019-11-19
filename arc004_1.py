import math

# 単なる全探索の鬼
# Me
def solve(p):
    max = 0
    for a in p:
        for b in p:
            c = math.sqrt((a[0] - b[0]) ** 2 + (a[1] - b[1]) ** 2)
            if c > max:
                max = c
    return max


n = int(input())
p = []
for i in range(n):
    p.append([int(x) for x in input().split()])


print(solve(p))
