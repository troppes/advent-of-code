package day14;

import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class PolymerGenerator {

    private final Map<String, String> elementsTable = new HashMap<>();
    private final Map<String, Long> elements = new HashMap<>();
    private Map<String, Long> polymerMap = new HashMap<>();

    PolymerGenerator(List<String> polymer) {
        elements.put(polymer.get(0), 1L);

        for (int i = 1; i < polymer.size(); i++) {

            Long elemVal = elements.get(polymer.get(i));
            elements.put(polymer.get(i), elemVal == null ? 1 : elemVal + 1);

            String combination = polymer.get(i - 1) + polymer.get(i);
            Long val = polymerMap.get(combination);
            polymerMap.put(combination, val == null ? 1 : val + 1);
        }
    }

    public void addElem(String key, String value) {
        elementsTable.put(key, value);
    }

    public void processPolymer() {
        Map<String, Long> newPolymer = new HashMap<>();

        for (Map.Entry<String, Long> element : polymerMap.entrySet()) {

            String newElement = elementsTable.get(element.getKey());

            // Put Element in Counting Table
            Long val = elements.get(newElement);
            elements.put(newElement, val == null ? element.getValue() : element.getValue() + val);

            String[] elements = element.getKey().split("");
            String newCombination1 = elements[0] + newElement;
            String newCombination2 = newElement + elements[1];

            Long valNewCombination1 = newPolymer.get(newCombination1);
            newPolymer.put(newCombination1, valNewCombination1 == null ? element.getValue() : element.getValue() + valNewCombination1);

            Long valNewCombination2 = newPolymer.get(newCombination2);
            newPolymer.put(newCombination2, valNewCombination2 == null ? element.getValue() : element.getValue() + valNewCombination2);

        }
        polymerMap = newPolymer;
    }

    public long getCalculation() {

        long max = Long.MIN_VALUE;
        long min = Long.MAX_VALUE;

        for (Map.Entry<String, Long> element : elements.entrySet()) {
            if (element.getValue() > max) max = element.getValue();
            if (element.getValue() < min) min = element.getValue();
        }
        return max - min;
    }
}
