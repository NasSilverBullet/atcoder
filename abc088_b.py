# Me
def buttle(a):
    alice = 0
    bob = 0

    a.sort(reverse=True)
    for i in range(len(a)):
        if i % 2 != 0:
            continue
        alice += a[i]
        if len(a) > i+1:
            bob += a[i+1]
    return alice - bob

n = int(input())
a = [int(x) for x in input().split()]
print(buttle(a))


# Other
def buttle(a):
    div = 0
    a.sort(reverse=True)
    for i in range(len(a)):
        # -1 の乗算で奇数indexのときに引くようにする
        div += ((-1) ** i) * a[i]

    return div


n = int(input())
a = list(map(int, input().split()))
print(buttle(a))
