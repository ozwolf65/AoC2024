with open("input.txt", "r") as f:
    data = f.readlines()
    grid = []
    for row in data:
        chars = [char for char in row if char != "\n"]
        grid.append(chars)

    def matchp1(i, j):
        matches = 0
        # Right
        if (
            j < len(grid[i]) - 3
            and grid[i][j] == "X"
            and grid[i][j + 1] == "M"
            and grid[i][j + 2] == "A"
            and grid[i][j + 3] == "S"
        ):
            matches += 1

        # Left
        if (
            j > 2
            and grid[i][j] == "X"
            and grid[i][j - 1] == "M"
            and grid[i][j - 2] == "A"
            and grid[i][j - 3] == "S"
        ):
            matches += 1

        # Down
        if (
            i < len(grid) - 3
            and grid[i][j] == "X"
            and grid[i + 1][j] == "M"
            and grid[i + 2][j] == "A"
            and grid[i + 3][j] == "S"
        ):
            matches += 1

        # Up
        if (
            i > 2
            and grid[i][j] == "X"
            and grid[i - 1][j] == "M"
            and grid[i - 2][j] == "A"
            and grid[i - 3][j] == "S"
        ):
            matches += 1

        # Down-Right
        if (
            j < len(grid[i]) - 3
            and i < len(grid) - 3
            and grid[i][j] == "X"
            and grid[i + 1][j + 1] == "M"
            and grid[i + 2][j + 2] == "A"
            and grid[i + 3][j + 3] == "S"
        ):
            matches += 1

        # Down-Left
        if (
            j > 2
            and i < len(grid) - 3
            and grid[i][j] == "X"
            and grid[i + 1][j - 1] == "M"
            and grid[i + 2][j - 2] == "A"
            and grid[i + 3][j - 3] == "S"
        ):
            matches += 1

        # Up-Right
        if (
            i > 2
            and j < len(grid[i]) - 3
            and grid[i][j] == "X"
            and grid[i - 1][j + 1] == "M"
            and grid[i - 2][j + 2] == "A"
            and grid[i - 3][j + 3] == "S"
        ):
            matches += 1

        # Up-Left
        if (
            i > 2
            and j > 2
            and grid[i][j] == "X"
            and grid[i - 1][j - 1] == "M"
            and grid[i - 2][j - 2] == "A"
            and grid[i - 3][j - 3] == "S"
        ):
            matches += 1
        return matches

    count = 0
    for i in range(len(grid)):
        for j in range(len(grid[i])):
            count += matchp1(i, j)

    print(count)

    def matchp2(i, j):
        if i < 1 or j < 1 or i >= len(grid) - 1 or j >= len(grid[i]) - 1:
            return 0
        centre = grid[i][j]
        top_left = grid[i - 1][j - 1]
        top_right = grid[i - 1][j + 1]
        bottom_left = grid[i + 1][j - 1]
        bottom_right = grid[i + 1][j + 1]
        good = ["M", "S"]
        if centre != "A":
            return 0
        # if not all M and S
        if (
            top_left not in good
            or top_right not in good
            or bottom_left not in good
            or bottom_right not in good
        ):
            return 0
        # If not spelling out MAS
        if top_left == bottom_right or top_right == bottom_left:
            return 0

        return 1

    count = 0
    for i in range(len(grid)):
        for j in range(len(grid[i])):
            count += matchp2(i, j)

    print(count)
