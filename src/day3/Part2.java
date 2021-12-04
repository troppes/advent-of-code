package day3;

import java.io.File;
import java.io.FileNotFoundException;
import java.util.ArrayList;
import java.util.Scanner;
import java.util.stream.Collectors;

public class Part2 {
    public static void main(String[] args) throws FileNotFoundException {
        File readings = new File("src/day3/diagnosticReport.txt");
        Scanner reader = new Scanner(readings);

        ArrayList<String> list = new ArrayList<>();

        while (reader.hasNextLine()) {
            String line = reader.nextLine();
            list.add(line);
        }

        String oxy = getBits(list, true);
        String co2 = getBits(list, false);
        System.out.println(Integer.parseInt(oxy,2) * Integer.parseInt(co2,2));
    }

    public static String getBits(ArrayList<String> list, boolean most){
        int oneCounter = 0;
        for (int i = 0; i < list.get(0).length(); i++) {
            for (String bin : list) {
                if (bin.charAt(i) == '1') {
                    oneCounter++;
                }
            }
            int finalI = i;
            if(list.size() == 1) return list.get(0);
            if(most){
                if (oneCounter >= list.size() / 2) {
                    list = (ArrayList<String>) list.stream().filter(s -> s.charAt(finalI) == '1').collect(Collectors.toList());
                } else {
                    list = (ArrayList<String>) list.stream().filter(s -> s.charAt(finalI) == '0').collect(Collectors.toList());
                }
            }else{
                if (oneCounter < list.size() / 2) {
                    list = (ArrayList<String>) list.stream().filter(s -> s.charAt(finalI) == '1').collect(Collectors.toList());
                } else {
                    list = (ArrayList<String>) list.stream().filter(s -> s.charAt(finalI) == '0').collect(Collectors.toList());
                }
            }
            oneCounter = 0;
        }
        return list.get(0);
    }
}
