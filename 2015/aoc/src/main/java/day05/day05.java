package day05;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.security.NoSuchAlgorithmException;
import java.util.List;
import java.util.stream.Collectors;

public class day05 {
    public static void main(String[] args) throws IOException, NoSuchAlgorithmException {
        List<String> lines = Files.readAllLines(Path.of("input.txt"));
        lines = lines.stream().filter(s -> {
            int v = 0;
            char[] chars = s.toCharArray();
            for (char c : chars) {
                if (c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u') {
                    v++;
                }
            }
            return v >= 3;
        }).filter(s -> {
            char[] chars = s.toCharArray();
            char curr = chars[0];
            for (int i = 1; i < chars.length; i++) {
                if (curr == chars[i]) {
                    return true;
                } else {
                    curr = chars[i];
                }
            }
            return false;
        }).filter(s -> !s.contains("ab") && !s.contains("cd") && !s.contains("pq") && !s.contains("xy")).collect(Collectors.toList());

        System.out.println(lines.size());

    }

}
