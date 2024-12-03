package day03;

import day02.Day02;
import org.junit.jupiter.api.Test;

import java.io.IOException;

import static org.junit.jupiter.api.Assertions.assertEquals;

class Day03Test {
    @Test
    void testSolveP1() throws IOException {
        String inputFilePath = "day03/input.txt";
        long expectedSolution = 161;

        long solution = Day03.solveP1(inputFilePath);

        assertEquals(expectedSolution, solution);
    }

    @Test
    void testSolveP2() throws IOException {
        String inputFilePath = "day03/input.txt";
        long expectedSolution = 48;

        long solution = Day03.solveP2(inputFilePath);

        assertEquals(expectedSolution, solution);
    }
}