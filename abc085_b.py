# Me
def mochi(d):
    d.sort(reverse=True)
    before = 101
    cnt = 0
    for n in d:
        if n >= before:
            continue
        before = n
        cnt += 1
    return cnt


n = int(input())
d = []
for i in range(n):
    d.append(int(input()))
print(mochi(d))

# Ohter
n = int(input())
d = []
for i in range(n):
    d.append(int(input()))
# 集合を使うことで一発
print(len(set(d)))
