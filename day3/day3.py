import re


def main(p2: bool):
    with open("input.txt", "r") as f:
        data = f.read()
        enabled = True
        disabledStrings = re.findall("(?s)don't\(\).*?do\(\)", data)
        if p2:
            for string in disabledStrings:
                data = data.replace(string, "")
        # for string in disabledStrings:
        # data = data.replace(string, "removed")
        matches: list[str] = re.findall("mul\([0-9]+,[0-9]+\)", data)
        total = 0
        for match in matches:
            lower = match.find("(")
            higher = match.find(")")
            nums = match[lower + 1 : higher]
            if nums == "":
                if match.find(")") != -1:
                    nums = match[: match.find(")")]
            left, right = nums.split(",")[0], nums.split(",")[1]
            total += int(left) * int(right)
        print(total)


if __name__ == "__main__":
    print("Advent of Code 2024 Day 3")
    print("Part 1")
    main(False)
    print("------------------\nPart 2")
    main(True)
