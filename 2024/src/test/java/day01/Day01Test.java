package day01;

import org.junit.jupiter.api.Test;

import java.io.IOException;


import static org.junit.jupiter.api.Assertions.assertEquals;

class Day01Test {
    @Test
    void testSolveP1() throws IOException {
        // Arrange
        String inputFilePath = "day01/input.txt";
        int expectedSolution = 11;

        // Act
        int solution = Day01.solveP1(inputFilePath);

        // Assert
        assertEquals(expectedSolution, solution);
    }

    @Test
    void testSolveP2() throws IOException {
        // Arrange
        String inputFilePath = "day01/input.txt";
        long expectedSolution = 31;

        // Act
        long solution = Day01.solveP2(inputFilePath);

        // Assert
        assertEquals(expectedSolution, solution);
    }
}