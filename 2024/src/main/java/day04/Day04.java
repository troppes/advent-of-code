package day04;

import util.Grid;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.util.regex.Matcher;
import java.util.regex.Pattern;
import java.util.stream.Collectors;

public class Day04 {
    public static void main(String[] args) throws IOException {
        System.out.println("Day01 Part 1 Solution: " + solveP1("day04/input.txt"));
        System.out.println("Day01 Part 2 Solution: " + solveP2("day04/input.txt"));
    }

    static long solveP1(String filePath) throws IOException {

        InputStream inputStream = Day04.class.getClassLoader().getResourceAsStream(filePath);

        if (inputStream == null) {
            System.err.println("Resource file not found!");
            return -1;
        }

        Grid<Character> charGrid;
        try (BufferedReader reader = new BufferedReader(new InputStreamReader(inputStream))) {
            charGrid = Grid.create(reader, line ->
                    line.chars().mapToObj(ch -> (char)ch).toArray(Character[]::new));
        }

        var findings = XMASFinder.findXMAS(charGrid);

        return findings.size();
    }

    static long solveP2(String filePath) throws IOException {

        InputStream inputStream = Day04.class.getClassLoader().getResourceAsStream(filePath);

        if (inputStream == null) {
            System.err.println("Resource file not found!");
            return -1;
        }

        Grid<Character> charGrid;
        try (BufferedReader reader = new BufferedReader(new InputStreamReader(inputStream))) {
            charGrid = Grid.create(reader, line ->
                    line.chars().mapToObj(ch -> (char)ch).toArray(Character[]::new));
        }

        return XMASFinder.findXMasPatterns(charGrid);
    }



}
