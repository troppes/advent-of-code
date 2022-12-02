package day7;

import java.io.File;
import java.io.FileNotFoundException;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;
import java.util.Scanner;
import java.util.stream.Collectors;

public class Part2 {
    public static void main(String[] args) throws FileNotFoundException {
        File file = new File("src/day7/submarines.txt");
        Scanner reader = new Scanner(file);

        ArrayList<Integer> submarines = Arrays.stream(reader.nextLine().split(",")).map(Integer::parseInt).collect(Collectors.toCollection(ArrayList::new));

        int leastFuel = Integer.MAX_VALUE;
        for(int i = 0; i < Collections.max(submarines); i++){
            int currentRunFuel = 0;
            for (Integer submarine: submarines) {
                int currentCost = 0;
                for(int j = 0; j <= Math.abs(submarine-i); j++){
                    currentCost += j;
                }
                currentRunFuel += currentCost;
            }
            if(leastFuel > currentRunFuel) leastFuel = currentRunFuel;
        }
        System.out.println(leastFuel);
    }
}
