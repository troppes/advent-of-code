package day15;

import java.io.File;
import java.io.FileNotFoundException;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.Scanner;
import java.util.stream.Collectors;

public class Part2 {
    public static void main(String[] args) throws FileNotFoundException {
        File file = new File("src/day15/heightMap.txt");
        Scanner reader = new Scanner(file);


        List<List<Integer>> riskLevelInput = new ArrayList<>();
        while (reader.hasNextLine()) {
            riskLevelInput.add((new ArrayList<>(Arrays.stream(reader.nextLine().trim().split("")).map(Integer::parseInt).collect(Collectors.toList()))));
        }

        // Extending Array Y
        List<List<Integer>> newRiskLevelInputY = new ArrayList<>(riskLevelInput);

        for (int i = 1; i < 5; i++) {
            for (List<Integer> current : riskLevelInput) {
                int finalI = i;
                newRiskLevelInputY.add(current.stream().map(e -> {
                    int newE = e + finalI;
                    return newE > 9 ? newE - 9 : newE;
                }).collect(Collectors.toList()));
            }
        }

        // Extending Array X
        List<List<Integer>> newRiskLevelInputX = new ArrayList<>();

        for (List<Integer> current : newRiskLevelInputY) {
            List<Integer> newCurrent = new ArrayList<>(current);
            for (int i = 1; i < 5; i++) {
                int finalI = i;
                newCurrent.addAll(current.stream().map(e -> {
                    int newE = e + finalI;
                    return newE > 9 ? newE - 9 : newE;
                }).collect(Collectors.toList()));
            }
            newRiskLevelInputX.add(newCurrent);
        }

        Cave cave = new Cave(newRiskLevelInputX);
        cave.calcShortestPath();

        System.out.println("Part 2: " + cave.getTotalRisk());

    }
}
