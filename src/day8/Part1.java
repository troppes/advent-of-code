package day8;

import java.io.File;
import java.io.FileNotFoundException;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Scanner;

public class Part1 {
    public static void main(String[] args) throws FileNotFoundException {
        File file = new File("src/day8/inputs.txt");
        Scanner reader = new Scanner(file);

        ArrayList<String> outputs = new ArrayList<>();

        while(reader.hasNextLine()){
            outputs.addAll(Arrays.asList(reader.nextLine().split(" \\| ")[1].split(" ")));
        }

        int counter = 0;
        // Check if 1,4,7,8
        for (String out: outputs) {
            if(out.length() == 2 || out.length() == 4 || out.length() == 3 || out.length() == 7) counter++;
        }
        System.out.println(counter);
    }
}
