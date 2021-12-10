package day10;

import java.io.File;
import java.io.FileNotFoundException;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Scanner;
import java.util.Stack;
import java.util.stream.Collectors;

public class Part1and2 {
    public static void main(String[] args) throws FileNotFoundException {

        File file = new File("src/day10/navigation.txt");
        Scanner reader = new Scanner(file);

        int scoreCorrupt = 0;
        ArrayList<Long> scoresIncomplete = new ArrayList<>();
        while (reader.hasNextLine()) {
            ArrayList<Character> line = Arrays.stream(reader.nextLine().trim().split("")).map(e -> e.charAt(0)).collect(Collectors.toCollection(ArrayList::new));
            scoreCorrupt += findCorrupted(line);
            scoresIncomplete.add(findIncomplete(line));
        }
        scoresIncomplete = scoresIncomplete.stream().sorted().filter(e -> e != 0).collect(Collectors.toCollection(ArrayList::new));

        System.out.println("Part 1: "+scoreCorrupt);
        System.out.println("Part 2: "+scoresIncomplete.get(scoresIncomplete.size()/2));
    }


    static int findCorrupted(ArrayList<Character> list) {
        // 0 = ) | 1 = ] | 2 = } | 3 = >
        int[] corruptedFound = new int[4];

        Stack<Character> currentOpen = new Stack<>();
        for (Character c : list) {
            if (c == '[' || c == '{' || c == '<') {
                currentOpen.push((char) (c+2));
            } else if (c == '('){
                currentOpen.push((char) (c+1));
            } else {
                if (!currentOpen.empty()) {
                    if (c != currentOpen.pop()) {
                       switch (c){
                           case ')': corruptedFound[0]++; break;
                           case ']': corruptedFound[1]++; break;
                           case '}': corruptedFound[2]++; break;
                           case '>': corruptedFound[3]++; break;
                       }
                    }
                }
            }
        }
        return ((corruptedFound[0]*3)+(corruptedFound[1]*57)+(corruptedFound[2]*1197)+(corruptedFound[3]*25137));
    }

    static long findIncomplete(ArrayList<Character> list) {

        Stack<Character> currentOpen = new Stack<>();
        for (Character c : list) {
            if (c == '[' || c == '{' || c == '<') {
                currentOpen.push((char) (c+2));
            } else if (c == '('){
                currentOpen.push((char) (c+1));
            } else {
                if (!currentOpen.empty()) {
                    if (c != currentOpen.pop()) {
                        currentOpen.clear();
                        break;
                    }
                }
            }
        }
        long score = 0;
        while (!currentOpen.empty()){
            score = score * 5;
            switch (currentOpen.pop()){
                case ')': score += 1; break;
                case ']': score += 2; break;
                case '}': score += 3; break;
                case '>': score += 4; break;
            }
        }
        return score;
    }
}
