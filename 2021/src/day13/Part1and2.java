package day13;

import java.io.File;
import java.io.FileNotFoundException;
import java.util.*;
import java.util.stream.Collectors;

public class Part1and2 {
    public static void main(String[] args) throws FileNotFoundException {
        File file = new File("src/day13/paper.txt");
        Scanner reader = new Scanner(file);

        ArrayList<Integer> x = new ArrayList<>(), y = new ArrayList<>();
        // Get Dimensions for Array
        while (reader.hasNextLine()) { // Coordinates
            String line = reader.nextLine();
            if(line.equals("")) break;
            List<Integer> parts = Arrays.stream(line.split(",")).map(Integer::parseInt).collect(Collectors.toList());
            x.add(parts.get(0));
            y.add(parts.get(1));
        }

        Paper paper = new Paper(Collections.max(x)+1, Collections.max(y)+1);

        for(int i = 0; i < x.size(); i++){
            paper.addCoordinate(x.get(i), y.get(i));
        }

        boolean first = true;
        while (reader.hasNextLine()){ // Folds
            String[] parts = reader.nextLine().split("=");
            paper.fold(Integer.parseInt(parts[1]),(parts[0].equals("fold along y")));
            if(first){
                first = false;
                System.out.println("Part 1: "+paper.countDots());
            }
        }
        System.out.println("Part 2 Code:");
        paper.print();
    }
}
