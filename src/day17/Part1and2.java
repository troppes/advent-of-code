package day17;

import java.io.File;
import java.io.FileNotFoundException;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.Scanner;
import java.util.stream.Collectors;

public class Part1and2 {
    public static void main(String[] args) throws FileNotFoundException {
        File file = new File("src/day17/target.txt");
        Scanner reader = new Scanner(file);

        String[] line = reader.nextLine().split(" ");

        List<Integer> xRange = new ArrayList<>(), yRange = new ArrayList<>();
        for (String s : line) {
            if (s.startsWith("x=")) xRange = Arrays.stream(s.substring(2, s.length() - 1).split("\\.\\.")).map(Integer::parseInt).collect(Collectors.toList());
            if (s.startsWith("y=")) yRange = Arrays.stream(s.substring(2).split("\\.\\.")).map(Integer::parseInt).collect(Collectors.toList());
        }

        Launcher launcher = new Launcher(xRange, yRange);

        System.out.println("Part 1: "+launcher.getMaxY());
        System.out.println("Part 2: "+launcher.getSuccessfulThrows());
    }
}
