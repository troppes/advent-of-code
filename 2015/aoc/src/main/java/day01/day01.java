package day01;

import java.io.*;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.List;

public class day01 {
    public static void main(String[] args) throws IOException {
        List<String> lines = Files.readAllLines(Path.of("input.txt"));
        String input = lines.get(0);
        int counter = 0;
        int currentFloor = 0;
        int firstTimeBasement = -1;
        for (char c : input.toCharArray()) {
            counter++;
            currentFloor = (c == '(') ? currentFloor + 1 : currentFloor - 1;
            if (currentFloor == -1 && firstTimeBasement == -1) {
                firstTimeBasement = counter;
            }
        }
        System.out.println("Part 1: " + currentFloor);
        System.out.println("Part 2: " + firstTimeBasement);
    }
}
