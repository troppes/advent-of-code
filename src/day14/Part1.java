package day14;

import java.io.File;
import java.io.FileNotFoundException;
import java.util.Arrays;
import java.util.Scanner;
import java.util.stream.Collectors;

public class Part1 {

    public static void main(String[] args) throws FileNotFoundException {
        File file = new File("src/day14/rules.txt");
        Scanner reader = new Scanner(file);

        PolymerGenerator gen = new PolymerGenerator(Arrays.stream(reader.nextLine().split("")).collect(Collectors.toList()));
        // Eliminate first empty line
        reader.nextLine();

        while (reader.hasNextLine()) {
            String[] parts = reader.nextLine().replaceAll("\\s", "").split("->");
            gen.addElem(parts[0], parts[1]);
        }

        for (int i = 0; i < 40; i++) {
            gen.processPolymer();
            if(i == 9) System.out.println("Part 1: "+gen.getCalculation());
        }
        System.out.println("Part 2: "+gen.getCalculation());
    }
}
