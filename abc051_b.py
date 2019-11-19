# Me
def resolve(k, s):
    c = 0
    for x in range(0, k+1):
        for y in range(0, k+1):
            if 0 <= s - (x + y) <= k:
                c += 1
    return c


k, s = map(int, input().split())
print(resolve(k, s))


# Other
# 理解できんかったけど一応メモ
K, S = list(map(int, input().split()))

num = 0
for x in range(K + 1):
    rest = S - x
    if 0 <= rest and rest <= K:
        num += rest + 1
    elif K < rest and rest <= 2 * K:
        num += 2 * K + 1 - rest
print(num)
K, S = list(map(int, input().split()))

num = 0
for x in range(K + 1):
    rest = S - x
    if 0 <= rest and rest <= K:
        num += rest + 1
    elif K < rest and rest <= 2 * K:
        num += 2 * K + 1 - rest
print(num)
