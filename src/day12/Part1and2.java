package day12;

import java.io.File;
import java.io.FileNotFoundException;
import java.util.Scanner;

public class Part1and2 {
    public static void main(String[] args) throws FileNotFoundException {
        File file = new File("src/day12/connections.txt");
        Scanner reader = new Scanner(file);

        CaveSystem caveSystem = new CaveSystem();

        while (reader.hasNextLine()) {
            String[] line = reader.nextLine().split("-");
            caveSystem.addCavesAndConnection(line[0],line[1]);
        }

        System.out.println("Part 1: "+caveSystem.findPaths());
        caveSystem.activatePart2();
        System.out.println("Part 2: "+caveSystem.findPaths());
    }
}
