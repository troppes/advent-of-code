package day5;

import java.io.File;
import java.io.FileNotFoundException;
import java.util.Arrays;
import java.util.List;
import java.util.Scanner;
import java.util.stream.Collectors;

public class Part1and2 {

    public static void main(String[] args) throws FileNotFoundException {
        File vents = new File("src/day5/vents.txt");
        Scanner reader = new Scanner(vents);

        Grid grid = new Grid(1000);

        while(reader.hasNextLine()){
            String line = reader.nextLine().replaceAll("\\s","");

            String[] parts = line.split("->");
            List<Integer> start = Arrays.stream(parts[0].split(",")).map(Integer::parseInt).collect(Collectors.toList());
            List<Integer> end = Arrays.stream(parts[1].split(",")).map(Integer::parseInt).collect(Collectors.toList());

            grid.addVentLine(start, end, false); // For Part2 switch to false
        }
        System.out.println(grid.getDoubleAndHigherVents());
    }
}
