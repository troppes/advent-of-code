package day16;

import java.io.File;
import java.io.FileNotFoundException;
import java.util.Scanner;

public class Part1and2 {
    public static void main(String[] args) throws FileNotFoundException {
        File file = new File("src/day16/transmission.txt");
        Scanner reader = new Scanner(file);

        String code = reader.hasNextLine() ? reader.nextLine() : "0";

        MessageDecoder decoder = new MessageDecoder(code);
        System.out.println("Part 1: " + decoder.masterPacket.sumOfVersions());
        System.out.println("Part 2: " + decoder.masterPacket.calculate());
    }
}
