package day5;

import java.util.List;
import java.util.Objects;

public class Grid {

    private final int gridSize;
    private final int[][] grid;

    Grid(int gridSize){
        this.gridSize = gridSize;
        grid = new int[gridSize][gridSize];
    }

    public void addVentLine(List<Integer> start, List<Integer> end, boolean disableDiagonal) {

        int startingPoint, endingPoint;
        boolean horizontal = false;

        if (Objects.equals(start.get(0), end.get(0))) { // vertical
            startingPoint = Math.min(start.get(1), end.get(1));
            endingPoint = Math.max(start.get(1), end.get(1));
        } else if (Objects.equals(start.get(1), end.get(1))) { // horizontal
            horizontal = true;
            startingPoint = Math.min(start.get(0), end.get(0));
            endingPoint = Math.max(start.get(0), end.get(0));
        } else { // diagonal
            if(disableDiagonal) return;

            while(!Objects.equals(start.get(0), end.get(0)) && !Objects.equals(start.get(1), end.get(1))){

                grid[start.get(1)][start.get(0)] += 1;

                if(start.get(0) < end.get(0)) start.set(0, (start.get(0) + 1));
                if(start.get(1) < end.get(1)) start.set(1, (start.get(1) + 1));

                if(start.get(0) > end.get(0)) start.set(0, (start.get(0) - 1));
                if(start.get(1) > end.get(1)) start.set(1, (start.get(1) - 1));

            }
            grid[end.get(1)][end.get(0)] += 1;
            return;
        }

        for (int i = startingPoint; i <= endingPoint; i++) {
            if (horizontal) {
                grid[start.get(1)][i] += 1;
            } else {
                grid[i][start.get(0)] += 1;
            }
        }

    }

    public int getDoubleAndHigherVents() {
        int counter = 0;
        for (int i = 0; i < gridSize; i++) {
            for (int j = 0; j < gridSize; j++) {
                if (grid[i][j] > 1) counter++;
            }
        }
        return counter;
    }

}
