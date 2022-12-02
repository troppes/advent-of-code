package day15;

import java.util.HashMap;
import java.util.LinkedList;
import java.util.List;
import java.util.Map;

public class RiskLevel {


    private List<RiskLevel> lowestRisk = new LinkedList<>();

    private Integer risk = Integer.MAX_VALUE;

    Map<RiskLevel, Integer> neighbours = new HashMap<>();

    int value, x, y;

    public void addDestination(RiskLevel destination, int distance) {
        neighbours.put(destination, distance);
    }

    public RiskLevel(int x, int y, int value) {
        this.x = x;
        this.y = y;
        this.value = value;
    }

    public void setRisk(Integer risk) {
        this.risk = risk;
    }

    public Integer getRisk() {
        return risk;
    }

    public List<RiskLevel> getLowestRisk() {
        return lowestRisk;
    }

    public void setLowestRisk(List<RiskLevel> lowestRisk) {
        this.lowestRisk = lowestRisk;
    }

    public Map<RiskLevel, Integer> getNeighbours() {
        return neighbours;
    }

    public int getValue() {
        return value;
    }

}