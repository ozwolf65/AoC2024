def comp(row) -> bool:
    # asc
    if row[0] < row[1]:
        return asc(row)
    elif row[0] > row[1]:
        return desc(row)
    else:
        return False


def asc(row) -> bool:
    for i in range(len(row) - 1):
        if not (1 <= row[i + 1] - row[i] <= 3):
            return False
    return True


def desc(row) -> bool:
    for i in range(len(row) - 1):
        if not (1 <= row[i] - row[i + 1] <= 3):
            return False
    return True


def p1(data):
    count = 0
    for line in data:
        row = [int(num) for num in line.split(" ")]
        if comp(row):
            count += 1
    print(count)


def comp_minus_one(row):
    for i in range(len(row)):
        minus_one = row[:i] + row[i + 1 :]
        if comp(minus_one):
            return True
    return False


def p2(data):
    count = 0
    for line in data:
        row = [int(num) for num in line.split(" ")]
        if comp(row) or comp_minus_one(row):
            count += 1

    print(count)


def main():
    with open("input.txt", "r") as f:
        data = f.readlines()
        print("Day 2 Advent of Code")
        print("-----------------")
        print("Part 1")
        p1(data)
        print("-----------------")
        print("Part 2")
        p2(data)


if __name__ == "__main__":
    main()
