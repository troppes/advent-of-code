package day8;

import java.io.File;
import java.io.FileNotFoundException;
import java.util.*;
import java.util.stream.Collectors;

public class Part2 {
    public static void main(String[] args) throws FileNotFoundException {
        File file = new File("src/day8/inputs.txt");
        Scanner reader = new Scanner(file);

        int value = 0;

        while (reader.hasNextLine()) {

            String[] line = reader.nextLine().split(" \\| ");

            List<Code> codes = Arrays.stream(line[0].split(" ")).map(Code::new).sorted().collect(Collectors.toList());
            List<Code> outputs = Arrays.stream(line[1].split(" ")).map(Code::new).collect(Collectors.toList());

            Code[] segments = new Code[10];

            // Since Sorted and every number only exists once, we can use this to determine the number
            segments[1] = codes.get(0); // 1
            segments[4] = codes.get(2); // 4
            segments[7] = codes.get(1); // 7
            segments[8] = codes.get(9); // 8

            ArrayList<Code> fives = new ArrayList<>(), sixes = new ArrayList<>();
            // Find other segments
            for (Code c: codes) {
                if(c.length() == 5) fives.add(c);
                if(c.length() == 6) sixes.add(c);
            }

            // Figure out Numbers in Length Five elements
            for (Code c: fives) {
                if(c.contains(segments[7])){ //3
                    segments[3] = c;
                }else if(c.intersect(segments[4], 3)){ // 5
                    segments[5] = c;
                }else{ // 2
                    segments[2] = c;
                }
            }

            // Figure out Numbers in Length Six elements
            for (Code c: sixes) {
                if(c.contains(segments[4])){ // 9
                    segments[9] = c;
                }else if(c.contains(segments[7])){ // 0
                    segments[0] = c;
                }else{ // 6
                    segments[6] = c;
                }
            }
            StringBuilder newNumber = new StringBuilder();

            for (Code out : outputs) {
                for(int i = 0; i < segments.length; i++){
                    if(out.equals(segments[i])) newNumber.append(i);
                }
            }
            value += Integer.parseInt(String.valueOf(newNumber));
        }
        System.out.println(value);
    }
}
