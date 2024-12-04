package day04;

import util.Direction;
import util.Grid;
import util.Point;

import java.util.ArrayList;
import java.util.List;

public class XMASFinder {
    public static List<List<Point>> findXMAS(Grid<Character> grid) {
        List<List<Point>> occurrences = new ArrayList<>();

        for (int y = 0; y < grid.getHeight(); y++) {
            for (int x = 0; x < grid.getWidth(); x++) {
                Point start = new Point(x, y);
                if (grid.get(start) == 'X') {
                    occurrences.addAll(searchFromPosition(grid, "XMAS", start));
                }
            }
        }
        return occurrences;
    }

    private static List<List<Point>> searchFromPosition(Grid<Character> grid, String word, Point start) {
        List<List<Point>> positions = new ArrayList<>();

        for (Direction dir : Direction.values()) {
            List<Point> path = new ArrayList<>();
            Point current = start.copy();
            boolean found = true;

            for (int i = 0; i < word.length(); i++) {
                if (!grid.isValid(current) || grid.get(current) != word.charAt(i)) {
                    found = false;
                    break;
                }

                path.add(current);
                current = current.step(dir);
            }

            if (found) {
                positions.add(path);
            }
        }
        return positions;
    }

    public static int findXMasPatterns(Grid<Character> grid) {
        int count = 0;

        for (int y = 1; y < grid.getHeight(); y++) {
            for (int x = 1; x < grid.getWidth(); x++) {
                Point center = new Point(x, y);
                if (grid.isValid(center) && grid.get(center) == 'A') {
                    if (isValidXMasPattern(grid, center)) {
                        count++;
                    }
                }
            }
        }
        return count;
    }

    private static boolean isValidXMasPattern(Grid<Character> grid, Point center) {
        Point[] topPattern = {
                center.step(Direction.UP_LEFT),
                center.step(Direction.UP_RIGHT)
        };

        Point[] bottomPattern = {
                center.step(Direction.DOWN_LEFT),
                center.step(Direction.DOWN_RIGHT)
        };

        for (Point p : topPattern) {
            if (!grid.isValid(p)) return false;
        }
        for (Point p : bottomPattern) {
            if (!grid.isValid(p)) return false;
        }

        char topLeft = grid.get(topPattern[0]);
        char topRight = grid.get(topPattern[1]);
        char bottomLeft = grid.get(bottomPattern[0]);
        char bottomRight = grid.get(bottomPattern[1]);

        return (
                // MAS from top-left to bottom-right AND MAS from top-right to bottom-left
                (isMAS(topLeft, bottomRight) && isMAS(topRight, bottomLeft)) ||
                        // SAM from top-left to bottom-right AND SAM from top-right to bottom-left
                        (isSAM(topLeft, bottomRight) && isSAM(topRight, bottomLeft)) ||
                        // MAS from top-left to bottom-right AND SAM from top-right to bottom-left
                        (isMAS(topLeft,  bottomRight) && isSAM(topRight, bottomLeft)) ||
                        // SAM from top-left to bottom-right AND MAS from top-right to bottom-left
                        (isSAM(topLeft,  bottomRight) && isMAS(topRight, bottomLeft))
        );
    }

    private static boolean isMAS(char M, char S) {
        return M == 'M' && S  == 'S';
    }

    private static boolean isSAM(char S, char M) {
        return S == 'S' && M == 'M';
    }
}
