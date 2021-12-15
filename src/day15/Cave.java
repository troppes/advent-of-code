package day15;

import java.util.*;

public class Cave {

    private final RiskLevel startingPoint, endPoint;
    private final Set<RiskLevel> settledRiskLevels = new HashSet<>();
    private final Set<RiskLevel> unsettledRiskLevels = new HashSet<>();


    Cave(List<List<Integer>> riskLevels) {
        RiskLevel[][] riskLevelNodes = new RiskLevel[riskLevels.size()][riskLevels.get(0).size()];

        // Create all Nodes
        for (int y = 0; y < riskLevels.size(); y++) {
            for (int x = 0; x < riskLevels.get(0).size(); x++) {
                riskLevelNodes[y][x] = new RiskLevel(x, y, riskLevels.get(y).get(x));
            }
        }

        // Fill all connections
        for (int y = 0; y < riskLevels.size(); y++) {
            for (int x = 0; x < riskLevels.get(0).size(); x++) {
                RiskLevel current = riskLevelNodes[y][x];

                // Horizontal
                if (x != 0) current.addDestination(riskLevelNodes[y][x - 1], riskLevelNodes[y][x - 1].getValue());
                if (x != (riskLevelNodes[y].length - 1))
                    current.addDestination(riskLevelNodes[y][x + 1], riskLevelNodes[y][x + 1].getValue());
                // Vertical
                if (y != 0) current.addDestination(riskLevelNodes[y - 1][x], riskLevelNodes[y - 1][x].getValue());
                if (y != (riskLevelNodes.length - 1))
                    current.addDestination(riskLevelNodes[y + 1][x], riskLevelNodes[y + 1][x].getValue());
            }
        }

        riskLevelNodes[0][0].setRisk(0);
        startingPoint = riskLevelNodes[0][0];
        endPoint = riskLevelNodes[riskLevels.size() - 1][riskLevels.get(0).size() - 1];
    }

    public int getTotalRisk() {
        return endPoint.getRisk();
    }

    public void calcShortestPath() {

        unsettledRiskLevels.clear();
        settledRiskLevels.clear();
        unsettledRiskLevels.add(startingPoint);

        while (unsettledRiskLevels.size() != 0) {
            RiskLevel currentRiskLevel = getLowestRiskNode();
            unsettledRiskLevels.remove(currentRiskLevel);

            for (Map.Entry<RiskLevel, Integer> neighbours : currentRiskLevel.getNeighbours().entrySet()) {
                RiskLevel neighbour = neighbours.getKey();
                Integer edgeWeight = neighbours.getValue();

                if (!settledRiskLevels.contains(neighbour)) {
                    calMinRisk(neighbour, edgeWeight, currentRiskLevel);
                    unsettledRiskLevels.add(neighbour);
                }
            }

            settledRiskLevels.add(currentRiskLevel);
        }
    }

    private RiskLevel getLowestRiskNode() {
        RiskLevel lowestRiskNode = null;
        int lowestRisk = Integer.MAX_VALUE;

        for (RiskLevel c : unsettledRiskLevels) {
            int nodeRisk = c.getRisk();
            if (nodeRisk < lowestRisk) {
                lowestRisk = nodeRisk;
                lowestRiskNode = c;
            }
        }
        return lowestRiskNode;
    }

    private static void calMinRisk(RiskLevel neighbourRiskLevel, Integer edgeWeight, RiskLevel currentRiskLevel) {
        Integer currentRisk = currentRiskLevel.getRisk();

        if (currentRisk + edgeWeight < neighbourRiskLevel.getRisk()) {
            neighbourRiskLevel.setRisk(currentRisk + edgeWeight);
            LinkedList<RiskLevel> lowestRisk = new LinkedList<>(currentRiskLevel.getLowestRisk());
            lowestRisk.add(currentRiskLevel);
            neighbourRiskLevel.setLowestRisk(lowestRisk);
        }
    }

}
