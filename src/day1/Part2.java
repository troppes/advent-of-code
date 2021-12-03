package day1;

import java.io.File;
import java.io.FileNotFoundException;
import java.util.ArrayList;
import java.util.List;
import java.util.Scanner;

public class Part2 {

    public static void main(String[] args) throws FileNotFoundException {
        File readingsFile = new File("src/day1/readings.txt");
        Scanner reader = new Scanner(readingsFile);
        int a, b, counter = 0;
        ArrayList<Integer> readings = new ArrayList<>();
        while(reader.hasNextLine()){
            readings.add(Integer.parseInt(reader.nextLine()));
        }
        for(int i = 0; i < readings.size()-3; i++){
            a = readings.get(i) + readings.get(i+1) + readings.get(i+2);
            b = readings.get(i + 1) + readings.get(i + 2) + readings.get(i + 3);

            if(b > a) counter++;
        }
        System.out.println(counter);
    }
}
