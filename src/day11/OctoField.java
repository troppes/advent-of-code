package day11;

import java.awt.*;
import java.util.ArrayList;

public class OctoField {
    private final ArrayList<ArrayList<Integer>> octopi = new ArrayList<>();
    private final ArrayList<Point> flashed = new ArrayList<>();
    private int flashCounter = 0;

    public void addRow(ArrayList<Integer> line) {
        octopi.add(line);
    }

    public int getFlashes() {
        return flashCounter;
    }

    public void raiseLevels() {
        for (ArrayList<Integer> list : octopi) {
            for (int i = 0; i < list.size(); i++) {
                list.set(i, (list.get(i) + 1));
            }
        }
    }

    public boolean checkFlashes() {
        flashed.clear();
        for (int y = 0; y < octopi.size(); y++) {
            for (int x = 0; x < octopi.get(y).size(); x++) {
                if (octopi.get(y).get(x) > 9 && !flashed.contains(new Point(x, y))) flash(x, y);
            }
        }
        for (Point p : flashed) {
            octopi.get(p.y).set(p.x, 0);
        }
        return flashed.size() == octopi.size() * octopi.get(0).size();
    }

    private void flash(int x, int y) {
        octopi.get(y).set(x, (octopi.get(y).get(x) + 1));

        if (octopi.get(y).get(x) > 9 && !flashed.contains(new Point(x, y))) {
            flashed.add(new Point(x, y));
            flashCounter++;

            // Horizontal
            if (x != 0) flash(x - 1, y);
            if (x != (octopi.get(y).size() - 1)) flash(x + 1, y);
            // Vertical
            if (y != 0) flash(x, y - 1);
            if (y != (octopi.size() - 1)) flash(x, y + 1);
            // Diagonal
            if (x != 0 && y != 0) flash(x - 1, y - 1);
            if (x != (octopi.get(y).size() - 1) && y != 0) flash(x + 1, y - 1);

            if (y != (octopi.size() - 1) && x != (octopi.get(y).size() - 1)) flash(x + 1, y + 1);
            if (y != (octopi.size() - 1) && x != 0) flash(x - 1, y + 1);
        }
    }

}
