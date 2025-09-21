"""pytest test.py -v"""

import pytest
from solution import solve_part1, solve_part2


def parse_instructions(input_str):
    """Helper function to parse instruction strings for testing"""
    instructions = []
    for instruction in input_str.split(', '):
        direction = instruction[0]
        steps = int(instruction[1:])
        instructions.append((direction, steps))
    return instructions


@pytest.mark.parametrize("input_str,expected", [
    ("R2, L3", 5),
    ("R2, R2, R2", 2), 
    ("R5, L5, R5, R3", 12),
])
def test_examples(input_str, expected):
    """Test the three examples you provided"""
    instructions = parse_instructions(input_str)
    result = solve_part1(instructions)
    assert result == expected, f"Input '{input_str}' should give {expected}, got {result}"


def test_r2_l3():
    """R2, L3 leaves you 2 blocks East and 3 blocks North, or 5 blocks away"""
    instructions = [('R', 2), ('L', 3)]
    result = solve_part1(instructions)
    assert result == 5


def test_r2_r2_r2():
    """R2, R2, R2 leaves you 2 blocks due South, which is 2 blocks away"""
    instructions = [('R', 2), ('R', 2), ('R', 2)]
    result = solve_part1(instructions)
    assert result == 2


def test_r5_l5_r5_r3():
    """R5, L5, R5, R3 leaves you 12 blocks away"""
    instructions = [('R', 5), ('L', 5), ('R', 5), ('R', 3)]
    result = solve_part1(instructions)
    assert result == 12

def test_p2_example():
    """R8, R4, R4, R8 first location visited twice is 4 blocks away"""
    instructions = [('R', 8), ('R', 4), ('R', 4), ('R', 8)]
    result = solve_part2(instructions)
    assert result == 4