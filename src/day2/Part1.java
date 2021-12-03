package day2;

import java.io.File;
import java.io.FileNotFoundException;
import java.util.Scanner;

public class Part1 {
    public static void main(String[] args) throws FileNotFoundException {
        File inputs = new File("src/day2/inputs.txt");
        Scanner reader = new Scanner(inputs);
        int depth = 0, position = 0;

        while (reader.hasNextLine()) {
            String input = reader.nextLine();
            String[] splitInput = input.split(" ");
            switch (splitInput[0]){
                case "forward":
                    position += Integer.parseInt(splitInput[1]);
                    break;
                case "down":
                    depth += Integer.parseInt(splitInput[1]);
                    break;
                case "up":
                    depth -= Integer.parseInt(splitInput[1]);
                    break;
                default:
                    throw new UnknownError("Input not known: "+splitInput[0]);
            }
        }
        System.out.println(depth*position);
    }
}