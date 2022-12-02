package day4;

import java.io.File;
import java.io.FileNotFoundException;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.Scanner;
import java.util.stream.Collectors;

public class Part2 {

    public static void main(String[] args) throws FileNotFoundException {
        File bingo = new File("src/day4/bingo.txt");
        Scanner reader = new Scanner(bingo);

        ArrayList<BingoField> fields = new ArrayList<>();
        List<Integer> numbersDrawn = Arrays.stream(reader.nextLine().split(",")).map(Integer::parseInt).collect(Collectors.toList());

        // Eliminate first empty line
        reader.nextLine();

        BingoField field = new BingoField();
        while (reader.hasNextLine()) {
            String line = reader.nextLine();
            if (line.equals("")) {
                fields.add(field);
                field = new BingoField();
            } else {
                field.addLine(new ArrayList<>(Arrays.stream(line.trim().split("\\s+")).map(Integer::parseInt).collect(Collectors.toList())));
            }
        }

        for (Integer number : numbersDrawn) {
            fields.removeIf(e -> {
                System.out.println(e.checkIfNumberExists(number));
                return e.checkIfNumberExists(number) != -1;
            } );
        }

    }
}
