# Me
def check(x, a, b):
    sum = 0
    while x != 0:
        sum += x % 10
        x //= 10
    if a <= sum <= b:
        return True
    return False


n, a, b = list(map(int, input().split()))
sum = 0
for i in range(1, n+1):
    if check(i, a, b):
        sum += i
print(sum)
