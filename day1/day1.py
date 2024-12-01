"""Day 1 : Advent of Code 2024"""

def generateLists(data: list[str]) -> tuple[list[int], list[int]]:
    left=[]
    right=[]
    for row in data:
        vals = row.split('   ')
        left.append(vals[0])
        right.append(vals[1])
    return left, right

def p1():
    with open('input.txt') as f:
        distance = 0
        data = f.read().splitlines()
        left, right = generateLists(data)

        left.sort()
        right.sort()
        for i in range(len(left)):
            distance += abs(int(left[i])-int(right[i]))
        print(distance)

def p2():
    with open('input.txt') as f:
        similarity = 0
        data = f.read().splitlines()
        left, right = generateLists(data)
        for i in left:
            similarity+=int(i)*right.count(i)
        print(similarity)


print("Day 1 Advent of Code")
print("-----------------")
print("Part 1")
p1()
print("-----------------")
print("Part 2")
p2()