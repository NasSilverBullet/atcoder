def get_pattern(n, y):
    X = 10
    Y = 5
    Z = 1

    y //= 1000
    for nx in range(n+1):
        m = n - nx
        for ny in range(m+1):
            nz = m - ny
            sum = nx * X + ny * Y + nz * Z
            if y == sum:
                return nx, ny, nz
    return -1, -1, -1


n, y = [int(x) for x in input().split()]
a, b, c = get_pattern(n, y)
print(a, b, c)

# Other
# https://atcoder.jp/contests/abc085/submissions/5565312
# TODO 理解できなかったので今度読む
