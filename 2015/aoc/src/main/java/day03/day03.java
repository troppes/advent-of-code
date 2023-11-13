package day03;

import java.awt.*;
import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.*;
import java.util.List;

public class day03 {
    public static void main(String[] args) throws IOException {
        List<String> lines = Files.readAllLines(Path.of("input.txt"));
        String input = lines.get(0);

        HashSet<Point> houses = new HashSet<>();

        int currentX = 0;
        int currentY = 0;

        houses.add(new Point(currentX, currentY));

        for (char c : input.toCharArray()) {
            switch (c){
                case '^':
                    currentY++;
                    break;
                case 'v':
                    currentY--;
                    break;
                case '>':
                    currentX++;
                    break;
                case '<':
                    currentX--;
                    break;
                default:
                    System.out.println("Not a valid input!");
                    break;
            }
            houses.add(new Point(currentX, currentY));
        }
        System.out.println(houses.size());
    }
}
