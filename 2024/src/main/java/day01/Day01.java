package day01;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;
import java.util.List;

public class Day01 {
    public static void main(String[] args) throws IOException {
        System.out.println("Day01 Part 1 Solution: " + solveP1("day01/input.txt"));
        System.out.println("Day01 Part 2 Solution: " + solveP2("day01/input.txt"));

    }

    static int solveP1(String filePath) throws IOException {

        InputStream inputStream = Day01.class.getClassLoader().getResourceAsStream(filePath);

        if (inputStream == null) {
            System.err.println("Resource file not found!");
            return -1;
        }

        List<Integer> l1 = new ArrayList<>();
        List<Integer> l2 = new ArrayList<>();
        try (BufferedReader reader = new BufferedReader(new InputStreamReader(inputStream))) {
            String line;
            while ((line = reader.readLine()) != null) {
                int[] numbers = Arrays.stream(line.split(" {3}")).mapToInt(Integer::parseInt).toArray();
                l1.add(numbers[0]);
                l2.add(numbers[1]);
            }
        }

        Collections.sort(l1);
        Collections.sort(l2);


        int solution = 0;
        for (int i = 0; i < l1.size(); i++) {
            solution += Math.abs(l1.get(i) - l2.get(i));
        }
        return solution;
    }

    static long solveP2(String filePath) throws IOException {

        InputStream inputStream = Day01.class.getClassLoader().getResourceAsStream(filePath);

        if (inputStream == null) {
            System.err.println("Resource file not found!");
            return -1;
        }

        List<Integer> l1 = new ArrayList<>();
        List<Integer> l2 = new ArrayList<>();
        try (BufferedReader reader = new BufferedReader(new InputStreamReader(inputStream))) {
            String line;
            while ((line = reader.readLine()) != null) {
                int[] numbers = Arrays.stream(line.split(" {3}")).mapToInt(Integer::parseInt).toArray();
                l1.add(numbers[0]);
                l2.add(numbers[1]);
            }
        }

        long solution = 0;
        for (long currentNumber : l1) {
            long noOfTimesFound = l2.stream().filter(no -> no == currentNumber).count();
            solution += currentNumber * noOfTimesFound;
        }
        return solution;
    }
}
