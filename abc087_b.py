# Me
def count(a, b, c, x):
    A = 10
    B = 2

    x /= 50

    count = 0
    for an in range(0, a+1):
        for bn in range(0, b+1):
            if 0 <= (x - (an * A + bn * B)) <= c:
                count += 1

    return count


a = int(input())
b = int(input())
c = int(input())
x = int(input())
print(count(a, b, c, x))
