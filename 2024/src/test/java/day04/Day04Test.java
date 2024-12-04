package day04;

import day03.Day03;
import org.junit.jupiter.api.Test;

import java.io.IOException;

import static org.junit.jupiter.api.Assertions.assertEquals;

class Day04Test {
    @Test
    void testSolveP1() throws IOException {
        String inputFilePath = "day04/input.txt";
        long expectedSolution = 18;

        long solution = Day04.solveP1(inputFilePath);

        assertEquals(expectedSolution, solution);
    }

    @Test
    void testSolveP2() throws IOException {
        String inputFilePath = "day04/input.txt";
        long expectedSolution = 9;

        long solution = Day04.solveP2(inputFilePath);

        assertEquals(expectedSolution, solution);
    }
}