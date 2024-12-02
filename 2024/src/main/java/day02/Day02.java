package day02;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.stream.Collectors;

public class Day02 {
    public static void main(String[] args) throws IOException {
        System.out.println("Day01 Part 1 Solution: " + solveP1("day02/input.txt"));
        System.out.println("Day01 Part 2 Solution: " + solveP2("day02/input.txt"));

    }

    static int solveP1(String filePath) throws IOException {

        InputStream inputStream = Day02.class.getClassLoader().getResourceAsStream(filePath);

        if (inputStream == null) {
            System.err.println("Resource file not found!");
            return -1;
        }

        int solution = 0;
        try (BufferedReader reader = new BufferedReader(new InputStreamReader(inputStream))) {
            String line;
            while ((line = reader.readLine()) != null) {
                List<Integer> numbers = Arrays.stream(line.split(" "))
                        .map(Integer::parseInt)
                        .collect(Collectors.toList());
                if (rateSafety(numbers)) solution++;
            }
        }
        return solution;
    }

    static long solveP2(String filePath) throws IOException {

        InputStream inputStream = Day02.class.getClassLoader().getResourceAsStream(filePath);

        if (inputStream == null) {
            System.err.println("Resource file not found!");
            return -1;
        }

        int solution = 0;
        try (BufferedReader reader = new BufferedReader(new InputStreamReader(inputStream))) {
            String line;
            while ((line = reader.readLine()) != null) {
                List<Integer> numbers = Arrays.stream(line.split(" "))
                        .map(Integer::parseInt)
                        .collect(Collectors.toList());

                if (rateSafety(numbers)) {
                    solution++;
                    continue;
                }

                for (int i = 0; i < numbers.size(); i++) {
                    List<Integer> currNumbers = new ArrayList<>(numbers);
                    currNumbers.remove(i);

                    if (rateSafety(currNumbers)) {
                        solution++;
                        break;
                    }
                }
            }
        }
        return solution;
    }

    private static boolean rateSafety(List<Integer> numbers) {

        boolean isIncreasing = true;
        boolean isDecreasing = true;

        for (int i = 0; i < numbers.size() - 1; i++) {
            int diff = numbers.get(i + 1) - numbers.get(i);

            if (Math.abs(diff) < 1 || Math.abs(diff) > 3) return false;

            if (diff < 0) isIncreasing = false;
            if (diff > 0) isDecreasing = false;
        }

        return isIncreasing || isDecreasing;
    }
}
