package day9;

import java.awt.*;
import java.util.ArrayList;
import java.util.Collections;
import java.util.Comparator;

public class Field {

    ArrayList<ArrayList<Integer>> field = new ArrayList<>();
    ArrayList<ArrayList<Integer>> basins = new ArrayList<>();
    ArrayList<Point> pointsInBasin = new ArrayList<>();

    void addRow(ArrayList<Integer> row) {
        field.add(row);
    }

    boolean isPositionLowPoint(int x, int y) {
        int currentPoint = field.get(y).get(x);

        if (x != 0 && currentPoint >= field.get(y).get(x - 1)) return false;
        if (x != (field.get(y).size() - 1) && currentPoint >= field.get(y).get(x + 1)) return false;
        if (y != 0 && currentPoint >= field.get(y - 1).get(x)) return false;
        if (y != (field.size() - 1) && currentPoint >= field.get(y + 1).get(x)) return false;

        return true;
    }

    int getRiskLevels() {
        int riskLevel = 0;
        for (int y = 0; y < field.size(); y++) {
            for (int x = 0; x < field.get(y).size(); x++) {
                int point = field.get(y).get(x);
                if (isPositionLowPoint(x, y)) riskLevel += (point + 1);
            }
        }
        return riskLevel;
    }

    int getBasinsLevels() {
        for (int y = 0; y < field.size(); y++) {
            for (int x = 0; x < field.get(y).size(); x++) {
                if (!pointsInBasin.contains(new Point(x, y))) {
                    ArrayList<Integer> arrayList = new ArrayList<>();
                    findBasin(x, y, arrayList);
                    basins.add(arrayList);
                }
            }
        }
        basins.sort(Comparator.comparingInt(ArrayList::size));
        return basins.get(basins.size() - 1).size() * basins.get(basins.size() - 2).size() * basins.get(basins.size() - 3).size();
    }

    void findBasin(int x, int y, ArrayList<Integer> list) {

        int currentPoint = field.get(y).get(x);
        if (currentPoint < 9 && !pointsInBasin.contains(new Point(x, y))) {
            list.add(currentPoint);
            pointsInBasin.add(new Point(x, y));

            if (x != 0) findBasin(x - 1, y, list);
            if (x != (field.get(y).size() - 1)) findBasin(x + 1, y, list);
            if (y != 0) findBasin(x, y - 1, list);
            if (y != (field.size() - 1)) findBasin(x, y + 1, list);
        }
    }
}
