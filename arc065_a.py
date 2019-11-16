# Me
def check(s):
    words = [
            'dream',
            'dreamer',
            'erase',
            'eraser',
            ]

    while len(s) != 0:
        flag = False
        for w in words:
            if s[-len(w):] == w:
                s = s[:-len(w)]
                flag = True
                break
        if not flag:
            return 'NO'
    return 'YES'


print(check(input()))


# Other
# これじゃだめだよな
def check(s):
    s = s.replace('eraser', '')
    s = s.replace('erase', '')
    s = s.replace('dreamer', '')
    s = s.replace('dream', '')
    if len(s) != 0:
        return 'NO'
    return 'YES'


print(check(input()))

# Other
S = input()
while S:
    for x in ['dreamer', 'eraser', 'dream', 'erase']:
        if S.endswith(x):
            S = S[:-len(x)]
            break
    # for に else 節を書ける
    # for を抜けたあとに実行される
    else:
        print('NO')
        exit()
print('YES')
