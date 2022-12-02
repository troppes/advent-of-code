package day12;

import java.util.*;
import java.util.stream.Collectors;

public class Cave {

    private final ArrayList<Cave> connections = new ArrayList<>();
    private final String name;
    private boolean bigCave = false;
    private boolean part2 = false;

    Cave(String name) {
        this.name = name;
        if (isUpper(name)) bigCave = true;
    }

    public static boolean isUpper(String s) {
        for (char c : s.toCharArray()) {
            if (!Character.isUpperCase(c))
                return false;
        }
        return true;
    }

    public String getName() {
        return name;
    }

    public void setPart2(boolean part2) {
        this.part2 = part2;
    }

    public ArrayList<Cave> getConnections() {
        return connections;
    }

    public void addConnection(Cave b) {
        if (!connections.contains(b)) {
            connections.add(b);
            b.addConnection(this);
        }
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        return this.name.equals(((Cave) o).getName());
    }

    @Override
    public int hashCode() {
        return Objects.hash(name);
    }

    public boolean canVisit(List<Cave> visited) {
        if(name.equals("start")) return false;
        if(part2){
            var smallCavesVisited = visited.stream().filter(cave -> !cave.bigCave).collect(Collectors.toList());
            Set<Cave> distinctSmallCavesVisited = new HashSet<>(smallCavesVisited);
            return bigCave || (smallCavesVisited.size() <= distinctSmallCavesVisited.size()+1);
        }else{
            return bigCave || !visited.contains(this);
        }
    }
}
