package day02;

import org.junit.jupiter.api.Test;

import java.io.IOException;

import static org.junit.jupiter.api.Assertions.assertEquals;

class Day02Test {
    @Test
    void testSolveP1() throws IOException {
        String inputFilePath = "day02/input.txt";
        int expectedSolution = 2;

        int solution = Day02.solveP1(inputFilePath);

        assertEquals(expectedSolution, solution);
    }

    @Test
    void testSolveP2() throws IOException {
        String inputFilePath = "day02/input.txt";
        long expectedSolution = 4;

        long solution = Day02.solveP2(inputFilePath);

        assertEquals(expectedSolution, solution);
    }
}