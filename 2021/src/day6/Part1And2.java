package day6;

import java.io.File;
import java.io.FileNotFoundException;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Scanner;
import java.util.stream.Collectors;

public class Part1And2 {
    public static void main(String[] args) throws FileNotFoundException {
        File fishDates = new File("src/day6/fishes.txt");
        Scanner reader = new Scanner(fishDates);

        ArrayList<Integer> fishes = Arrays.stream(reader.nextLine().split(",")).map(Integer::parseInt).collect(Collectors.toCollection(ArrayList::new));

        long[] birthCycle = new long[9];

        fishes.forEach(e -> birthCycle[e]++);


        for(int i = 0; i < 256; i++){ // Change Number to 60 for part1
            long tempFishes, oldFishes = 0;
            for(int j = birthCycle.length-1; j >= 0; j--){
                tempFishes = birthCycle[j];
                birthCycle[j] = oldFishes;
                oldFishes = tempFishes;
                if(j == 0){
                    birthCycle[6] += oldFishes;
                    birthCycle[8] += oldFishes;
                }
            }
        }

        long totalFishes = 0;
        for (long i : birthCycle) {
            totalFishes += i;
        }
        System.out.println(totalFishes);
    }
}
