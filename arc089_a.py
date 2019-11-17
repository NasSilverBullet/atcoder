from sys import stdin


# Me
class Check():
    def __init__(self):
        self.now_t, self.now_x, self.now_y = 0, 0, 0

    def do(self, t, x, y):
        distance = ((x - self.now_x) + (y - self.now_y))
        duration = t - self.now_t
        if duration < distance:
            return False

        if (distance - duration) % 2 != 0:
            return False

        self.now_t, self.now_x, self.now_y = t, x, y
        return True


N = int(input())
c = Check()
for i in range(N):
    a = [int(x) for x in input().split()]
    if not c.do(a[0], a[1], a[2]):
        print('No')
        exit()
print('Yes')


# Other
input = stdin.readline
lines = stdin.readlines

N = int(input())

txy = ((map(int, line.split())) for line in lines())


def main():
    pt = 0
    px = 0
    py = 0

    flg = False

    for t, x, y in txy:
        dt = t - pt - abs(x - px) - abs(y - py)
        # 余った時間が奇数 戻れない
        if dt < 0 or dt % 2 == 1:
            break

        pt = t
        px = x
        py = y
    else:
        # break しなかった 完走
        flg = True

    print("Yes" if flg else "No")
    return


main()
