package util;

import java.io.BufferedReader;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;
import java.util.function.Function;

public class Grid<T> {
    private final int xSize;
    private final int ySize;
    private final T[][] grid;

    @SuppressWarnings("unchecked")
    private Grid(List<T[]> gridData) {
        this.ySize = gridData.size();
        this.xSize = gridData.getFirst().length;
        this.grid = (T[][]) new Object[ySize][xSize];

        for (int y = 0; y < ySize; y++) {
            grid[y] = gridData.get(y);
        }
    }

    public T get(Point p) {
        if (!isValid(p)) {
            throw new IllegalArgumentException("Position (" + p.x() + "," + p.y() + ") is out of bounds");
        }
        return grid[p.y()][p.x()];
    }

    public void set(Point p, T value) {
        if (!isValid(p)) {
            throw new IllegalArgumentException("Position (" + p.x() + "," + p.y() + ") is out of bounds");
        }
        grid[p.y()][p.x()] = value;
    }

    public boolean isValid(Point point) {
        return point.x() >= 0 && point.x() < xSize && point.y() >= 0 && point.y() < ySize;
    }

    public int getWidth() {
        return xSize;
    }

    public int getHeight() {
        return ySize;
    }

    public static <T> Grid<T> create(BufferedReader reader, Function<String, T[]> parser) throws IOException {
        List<T[]> gridData = new ArrayList<>();
        String line;

        while ((line = reader.readLine()) != null) {
            gridData.add(parser.apply(line));
        }

        return new Grid<>(gridData);
    }

    @Override
    public String toString() {
        StringBuilder sb = new StringBuilder();
        for (int y = 0; y < ySize; y++) {
            for (int x = 0; x < xSize; x++) {
                sb.append(grid[y][x] == null ? "." : grid[y][x].toString()).append(" ");
            }
            sb.append("\n");
        }
        return sb.toString();
    }
}
