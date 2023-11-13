package day03;

import java.awt.*;
import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.HashSet;
import java.util.List;

public class day03part2 {
    public static void main(String[] args) throws IOException {
        List<String> lines = Files.readAllLines(Path.of("input.txt"));
        String input = lines.get(0);

        HashSet<Point> houses = new HashSet<>();

        int currentX = 0;
        int currentY = 0;

        int currentRoboX = 0;
        int currentRoboY = 0;

        boolean isRobo = false;

        houses.add(new Point(currentX, currentY));

        for (char c : input.toCharArray()) {
            switch (c){
                case '^':
                    if (isRobo){
                        currentY++;
                    } else {
                        currentRoboY++;
                    }
                    break;
                case 'v':
                    if (isRobo){
                        currentY--;
                    } else {
                        currentRoboY--;
                    }
                    break;
                case '>':
                    if (isRobo){
                        currentX++;
                    } else {
                        currentRoboX++;
                    }
                    break;
                case '<':
                    if (isRobo){
                        currentX--;
                    } else {
                        currentRoboX--;
                    }
                    break;
                default:
                    System.out.println("Not a valid input!");
                    break;
            }
            isRobo = !isRobo;
            houses.add(new Point(currentX, currentY));
            houses.add(new Point(currentRoboX, currentRoboY));
        }
        System.out.println(houses.size());
    }
}
