testcases = int(input())


def fn(input: str, n: int, k: int):
    count = [0] * 26
    oddCount = 0
    isFinalEven = (n - k) % 2 == 0

    for ch in input:
        idx = ord(ch) - ord("a")
        count[idx] += 1

    for idx, chCount in enumerate(count):
        if k > 0 and chCount % 2 == 1:
            count[idx] -= 1
            k -= 1
        if count[idx] % 2 == 1:
            oddCount += 1

    if isFinalEven:
        return oddCount == 0 and k % 2 == 0

    for idx, chCount in enumerate(count):
        if chCount % 2 == 0 and k > 1:
            diff = min(chCount, k)
            if diff % 2 == 1:
                diff -= 1
            count[idx] -= diff
            k -= diff

    return (oddCount == 1 and k == 0) or (oddCount == 0 and k == 1)


for i in range(testcases):
    n, k = list(map(lambda x: int(x), input().split(" ")))
    line = input()
    res = fn(line, n, k)
    if res:
        print("YES")
    else:
        print("NO")
