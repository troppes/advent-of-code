package day04;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.security.MessageDigest;
import java.security.NoSuchAlgorithmException;
import java.util.HexFormat;
import java.util.List;

public class day04 {
    public static void main(String[] args) throws IOException, NoSuchAlgorithmException {
        List<String> lines = Files.readAllLines(Path.of("input.txt"));
        String input = lines.get(0);

        int counter = 0;
        while (true) {

            MessageDigest md = MessageDigest.getInstance("MD5");

            String pw = input + counter;
            md.update(pw.getBytes());
            byte[] bytes = md.digest();

            String generatedPassword = HexFormat.of().formatHex(bytes);

            if (generatedPassword.startsWith("000000")) {
                System.out.println(counter);
                break;
            }

            counter++;
        }

    }

}
