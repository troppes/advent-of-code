package day9;

import java.io.File;
import java.io.FileNotFoundException;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Scanner;
import java.util.stream.Collectors;

public class Part1and2 {
    public static void main(String[] args) throws FileNotFoundException {
        File file = new File("src/day9/heightmap.txt");
        Scanner reader = new Scanner(file);

        Field field = new Field();

        while (reader.hasNextLine()) {
            field.addRow(new ArrayList<>(Arrays.stream(reader.nextLine().trim().split("")).map(Integer::parseInt).collect(Collectors.toList())));
        }

        System.out.println("Risk-Level: " + field.getRiskLevels());
        System.out.println("Biggest-Basins: " + field.getBasinsLevels());
    }
}
