import numpy as np
# Me
# numpy は大量のデータに最適化されているため、少量では速度が出ない
def f():
    i = int(input())
    nums = np.array(list(map(int, input().split())))
    # nums = [int(x) for x in input().split()]
    c = 0
    while True:
        if any(nums % 2 == 1):
            break
        c += 1
        nums = nums / 2
    print(c)


f()


# Other
# https://abc081.contest.atcoder.jp/submissions/1842030
def how_many_times_divisible(n):
    ans = 0
    while n % 2 == 0:
        n /= 2
        ans += 1
    return ans


n = int(input())
a = map(int, input().split())
ans = min(map(how_many_times_divisible, a))

print(ans)


def divide_all(a):
    success = True
    for i, n in enumerate(a):
        if n % 2 != 0:
            success = False
            break
        a[i] = n / 2
    if success is False:
        return (a, False)
    return (a, True)

n = int(input())
a = list(map(int, input().split()))
c = 0
while True:
    a, r = divide_all(a)
    if r is False:
        break
    c += 1


print(c)
