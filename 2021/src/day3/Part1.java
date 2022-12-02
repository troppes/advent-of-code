package day3;

import java.io.File;
import java.io.FileNotFoundException;
import java.util.Scanner;

public class Part1 {
    public static void main(String[] args) throws FileNotFoundException {
        File readings = new File("src/day3/diagnosticReport.txt");
        Scanner reader = new Scanner(readings);
        int[] ones = new int[12];
        int lines = 0;

        while(reader.hasNextLine()){
            lines++;
            String line = reader.nextLine();
            for(int i = 0, n = line.length() ; i < n ; i++) {
                if(line.charAt(i) == '1') ones[i]++;
            }
        }
        StringBuilder binaryStringGamma = new StringBuilder();
        StringBuilder binaryStringEpsilon = new StringBuilder();
        for (int i : ones) {
            if(i > lines/2){
                binaryStringGamma.append(1);
                binaryStringEpsilon.append(0);
            }else {
                binaryStringGamma.append(0);
                binaryStringEpsilon.append(1);
            }
        }
        System.out.println(Integer.parseInt(binaryStringGamma.toString(),2) * Integer.parseInt(binaryStringEpsilon.toString(),2));
    }
}
