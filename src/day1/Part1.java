package day1;

import java.io.File;
import java.io.FileNotFoundException;
import java.util.Scanner;

public class Part1 {

    public static void main(String[] args) throws FileNotFoundException {
        File readings = new File("src/day1/readings.txt");
        Scanner reader = new Scanner(readings);
        int counter = 0;
        int lastNumber = Integer.MAX_VALUE;

        while(reader.hasNextLine()){
            int reading = Integer.parseInt(reader.nextLine());
            if(reading > lastNumber) counter++;
            lastNumber = reading;
        }
        System.out.println(counter);

    }
}
