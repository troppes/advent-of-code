package day11;

import java.io.File;
import java.io.FileNotFoundException;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Scanner;
import java.util.stream.Collectors;

public class Part1and2 {
    public static void main(String[] args) throws FileNotFoundException {
        File file = new File("src/day11/octopi.txt");
        Scanner reader = new Scanner(file);

        OctoField field = new OctoField();

        while (reader.hasNextLine()) {
            field.addRow(new ArrayList<>(Arrays.stream(reader.nextLine().trim().split("")).map(Integer::parseInt).collect(Collectors.toList())));
        }

        int counter = 0;
        while (!field.checkFlashes()) {
            if (counter == 100) System.out.println("Part 1: " + field.getFlashes());
            field.raiseLevels();
            counter++;
        }
        System.out.println("Part 2: " + counter);
    }
}
