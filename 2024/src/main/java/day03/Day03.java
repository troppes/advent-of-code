package day03;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.regex.Matcher;
import java.util.regex.Pattern;
import java.util.stream.Collectors;

public class Day03 {
    public static void main(String[] args) throws IOException {
        System.out.println("Day01 Part 1 Solution: " + solveP1("day03/input.txt"));
        System.out.println("Day01 Part 2 Solution: " + solveP2("day03/input.txt"));
    }

    static long solveP1(String filePath) throws IOException {

        InputStream inputStream = Day03.class.getClassLoader().getResourceAsStream(filePath);

        if (inputStream == null) {
            System.err.println("Resource file not found!");
            return -1;
        }

        long solution = 0;
        try (BufferedReader reader = new BufferedReader(new InputStreamReader(inputStream))) {
            String lines = reader.lines().collect(Collectors.joining("\n"));
            Pattern pattern = Pattern.compile("mul\\((\\d{1,3}),(\\d{1,3})\\)");
            Matcher matcher = pattern.matcher(lines);

            while (matcher.find()) {
                long a = Long.parseLong(matcher.group(1));
                long b = Long.parseLong(matcher.group(2));
                solution += a * b;
            }
        }
        return solution;
    }

    static long solveP2(String filePath) throws IOException {

        InputStream inputStream = Day03.class.getClassLoader().getResourceAsStream(filePath);

        if (inputStream == null) {
            System.err.println("Resource file not found!");
            return -1;
        }

        long solution = 0;
        try (BufferedReader reader = new BufferedReader(new InputStreamReader(inputStream))) {
            String lines = reader.lines().collect(Collectors.joining("\n"));
            String[] parts = lines.split("(?=((?:do|don't)\\(\\)))");

            for (String part : parts) {
                if (!part.startsWith("don't")) {
                    Pattern mulPattern = Pattern.compile("mul\\((\\d{1,3}),(\\d{1,3})\\)");
                    Matcher mulMatcher = mulPattern.matcher(part);
                    while (mulMatcher.find()) {
                        long a = Integer.parseInt(mulMatcher.group(1));
                        long b = Integer.parseInt(mulMatcher.group(2));
                        solution += a * b;
                    }
                }
            }
        }
        return solution;
    }

}
