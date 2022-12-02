package day13;

public class Paper {

    private boolean[][] grid;

    Paper(int x, int y) {
        grid = new boolean[y][x];
    }

    public void addCoordinate(int x, int y) {
        grid[y][x] = true;
    }

    public void fold(int n, boolean horizontal) {
        if (horizontal) {
            for (int y = n + 1; y < grid.length; y++) {
                for (int x = 0; x < grid[0].length; x++) {
                    if (grid[y][x]) {
                        grid[y][x] = false; // current point false
                        grid[n - (y - n)][x] = true;
                    }
                }
            }

            // Resize
            int newSizeY = grid.length / 2;
            boolean[][] newGrid = new boolean[newSizeY][grid[0].length];
            for (int y = 0; y < newSizeY; y++) {
                System.arraycopy(grid[y], 0, newGrid[y], 0, grid[0].length);
            }
            grid = newGrid;

        } else {
            for (int y = 0; y < grid.length; y++) {
                for (int x = n+1; x < grid[0].length; x++) {
                    if (grid[y][x]) {
                        grid[y][x] = false; // current point false
                        grid[y][n - (x - n)] = true;
                    }
                }
            }

            // Resize
            int newSizeX = grid[0].length / 2;
            boolean[][] newGrid = new boolean[grid.length][newSizeX];
            for (int y = 0; y < grid.length; y++) {
                System.arraycopy(grid[y], 0, newGrid[y], 0, newSizeX);
            }
            grid = newGrid;
        }
    }

    public void print() {
        for (boolean[] row : grid) {
            StringBuilder line = new StringBuilder();
            for (boolean cell : row) {
                if (cell) {
                    line.append("# ");
                } else {
                    line.append("  ");
                }
            }
            System.out.println(line);
        }
    }

    public int countDots() {
        int dots = 0;
        for (boolean[] booleans : grid) {
            for (int x = 0; x < grid[0].length; x++) {
                if (booleans[x]) dots++;
            }
        }
        return dots;
    }
}
