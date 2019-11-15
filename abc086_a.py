# Me
def f():
    a, b = input().split()
    a, b = int(a), int(b)
    rem = int(a) * int(b) % 2
    if rem == 1:
        print('Odd')
        return
    print('Even')


f()


# Other
def f():
    # map(function, args(list))
    a, b = map(int, input().split())
    result = a * b
    if result % 2 == 1:
        print('Odd')
    else:
        print('Even')


f()
