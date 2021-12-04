package day4;

import java.io.File;
import java.io.FileNotFoundException;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.Scanner;
import java.util.stream.Collectors;

public class Part1 {

    public static void main(String[] args) throws FileNotFoundException {
        File bingo = new File("src/day4/bingo.txt");
        Scanner reader = new Scanner(bingo);

        ArrayList<BingoField> fields = new ArrayList<>();
        List<Integer> numbersDrawn = Arrays.stream(reader.nextLine().split(",")).map(Integer::parseInt).collect(Collectors.toList());

        // Eliminate first empty line
        reader.nextLine();

        BingoField field = new BingoField();
        while(reader.hasNextLine()){
            String line = reader.nextLine();
            if(line.equals("")) {
                fields.add(field);
                field = new BingoField();
            }else{
                field.addLine(new ArrayList<>(Arrays.stream(line.trim().split("\\s+")).map(Integer::parseInt).collect(Collectors.toList())));
            }
        }

        boolean foundWinner = false;
        for (Integer number: numbersDrawn) {
            for (BingoField bf: fields) {
                int result = bf.checkIfNumberExists(number);
                if(result != -1){
                    System.out.println(result);
                    foundWinner = true;
                    break;
                }
            }
            if(foundWinner) break;
        }

    }
}
