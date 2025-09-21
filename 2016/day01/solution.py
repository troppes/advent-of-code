def read_file(filename):
    with open(filename, 'r') as f:
        data = f.read().strip()

    instructions = []
    for instruction in data.split(', '):
        direction = instruction[0]
        steps = int(instruction[1:])
        instructions.append((direction, steps))
    return instructions

def solve_part1(instructions):
    x, y = 0, 0
    direction = 0 
    for nextDirection, steps in instructions:
        if nextDirection == 'R':
            direction = (direction + 1) % 4 
        else: 
            direction = (direction - 1) % 4

        if direction == 0:    # North
            y += steps
        elif direction == 1:  # East  
            x += steps
        elif direction == 2:  # South
            y -= steps
        elif direction == 3:  # West
            x -= steps
    return abs(x) + abs(y)

def solve_part2(instructions):
    x, y = 0, 0
    direction = 0
    visited = set()
    for nextDirection, steps in instructions:
        if nextDirection == 'R':
            direction = (direction + 1) % 4 
        else: 
            direction = (direction - 1) % 4

        for _ in range(steps):
            if direction == 0:    # North
                y += 1
            elif direction == 1:  # East  
                x += 1
            elif direction == 2:  # South
                y -= 1
            elif direction == 3:  # West
                x -= 1
            if (x, y) in visited:
                return abs(x) + abs(y)
            else:
                visited.add((x, y))

print("Solution Part1:" + str(solve_part1(read_file('input.txt'))))
print("Solution Part2:" + str(solve_part2(read_file('input.txt'))))