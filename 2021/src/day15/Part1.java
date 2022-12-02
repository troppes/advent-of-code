package day15;

import java.io.File;
import java.io.FileNotFoundException;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.Scanner;
import java.util.stream.Collectors;

public class Part1 {
    public static void main(String[] args) throws FileNotFoundException {
        File file = new File("src/day15/heightMap.txt");
        Scanner reader = new Scanner(file);

        List<List<Integer>> riskLevelInput = new ArrayList<>();
        while (reader.hasNextLine()) {
            riskLevelInput.add((new ArrayList<>(Arrays.stream(reader.nextLine().trim().split("")).map(Integer::parseInt).collect(Collectors.toList()))));
        }

        Cave cave = new Cave(riskLevelInput);
        cave.calcShortestPath();

        System.out.println("Part 1: " + cave.getTotalRisk());

    }
}
