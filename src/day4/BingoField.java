package day4;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;

public class BingoField {
    ArrayList<ArrayList<Integer>> numbers = new ArrayList<>();
    ArrayList<ArrayList<Boolean>> numbersCorrect = new ArrayList<>();

    public void addLine(ArrayList<Integer> line) {
        numbers.add(line);

        ArrayList<Boolean> boolLine = new ArrayList<>(Arrays.asList(new Boolean[line.size()]));
        Collections.fill(boolLine, Boolean.FALSE);
        numbersCorrect.add(boolLine);

    }

    public int checkIfNumberExists(Integer number){
        for (int i = 0; i < numbers.size(); i++) {
            int index = numbers.get(i).indexOf(number);
            if(index != -1){
                numbersCorrect.get(i).set(index, true);
            }
        }
        if(checkWinCondition()){
            int notMarkedSum = 0;
            for (int i = 0; i < numbersCorrect.size(); i++) {
                for(int j = 0; j < numbersCorrect.get(i).size(); j++){
                    if(!numbersCorrect.get(i).get(j)){
                        notMarkedSum += numbers.get(i).get(j);
                    }
                }
            }
            return notMarkedSum*number;
        }
        return -1;
    }

    public boolean checkWinCondition(){
        for (ArrayList<Boolean> booleans : numbersCorrect) {
            int index = booleans.indexOf(Boolean.FALSE);
            if (index == -1) {
                return true;
            }
        }
        boolean foundWin = true;
        for (int i = 0; i < numbersCorrect.get(0).size(); i++) {
            for (ArrayList<Boolean> booleans : numbersCorrect) {
                if (!booleans.get(i)) {
                    foundWin = false;
                    break;
                }
            }
        }
        return foundWin;
    }

}
