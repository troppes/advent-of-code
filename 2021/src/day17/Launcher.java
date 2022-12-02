package day17;

import java.util.*;

public class Launcher {

    Set<Integer> xRange, yRange;

    int maxY = -1;
    int successfulThrows = -1;

    public Launcher(List<Integer> xRange, List<Integer> yRange) {

        this.xRange = new HashSet<>();
        this.yRange = new HashSet<>();

        for(int i = xRange.get(0); i <= xRange.get(1); i++){
            this.xRange.add(i);
        }

        for(int i = yRange.get(0); i <= yRange.get(1); i++){
            this.yRange.add(i);
        }

        calculate();
    }

    public int getMaxY() {
        return maxY;
    }

    public int getSuccessfulThrows() {
        return successfulThrows;
    }

    public void calculate(){

        int maxY = -1;
        int count = 0;

        for (int x = 1; x < 1000; x++) {
            for (int y = Collections.min(yRange); y < 1000; y++) {
                int result = launch(x,y);
                if(result != -1){
                    count++;
                }
                if(result > maxY){
                    maxY = result;
                }
            }
        }

        this.successfulThrows = count;
        this.maxY = maxY;
    }

    // Return maxY and -1 if the target is Missed
    public int launch(int veloX, int veloY){

        int x = 0;
        int y = 0;

        int maxY = 0;

        while (x < Collections.max(xRange) && y > Collections.min(yRange)){
            x += veloX;
            y += veloY;

            if(y > maxY) maxY = y;
            if(xRange.contains(x) && yRange.contains(y)) return maxY;

            if (veloX > 0) veloX--;
            veloY--;
        }
        return -1;
    }
}
