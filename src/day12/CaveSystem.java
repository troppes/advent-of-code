package day12;

import java.util.ArrayList;
import java.util.LinkedList;
import java.util.List;
import java.util.Optional;

public class CaveSystem {

    private final ArrayList<Cave> parts = new ArrayList<>();
    private int pathCount = 0;

    public void addCavesAndConnection(String c1, String c2) {

        Optional<Cave> optionalCave1 = parts.stream().filter(e -> e.getName().equals(c1)).findFirst();
        Optional<Cave> optionalCave2 = parts.stream().filter(e -> e.getName().equals(c2)).findFirst();

        Cave cave1 = optionalCave1.orElseGet(() -> new Cave(c1));
        Cave cave2 = optionalCave2.orElseGet(() -> new Cave(c2));

        if (optionalCave1.isEmpty()) parts.add(cave1);
        if (optionalCave2.isEmpty()) parts.add(cave2);

        cave1.addConnection(cave2);

    }

    public void activatePart2() {
        for (Cave c : parts) {
            c.setPart2(true);
        }
    }

    public int findPaths() {
        pathCount = 0;

        Optional<Cave> startCaveOptional = parts.stream().filter(e -> e.getName().equals("start")).findFirst();
        Optional<Cave> endCaveOptional = parts.stream().filter(e -> e.getName().equals("end")).findFirst();
        Cave start = startCaveOptional.orElse(null);
        Cave end = endCaveOptional.orElse(null);

        if (start == null || end == null) return -1;
        findPath(start, end, new LinkedList<>());
        return pathCount;
    }

    private void findPath(Cave from, Cave to, List<Cave> visited) {
        visited.add(from);
        if (from.equals(to)) {
            pathCount++;
        } else {
            for (Cave c : from.getConnections()) {
                if (c.canVisit(visited)) {
                    findPath(c, to, new LinkedList<>(visited));
                }
            }
        }
    }

}
