package day05;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.List;
import java.util.stream.Collectors;

public class day05Part2 {
    public static void main(String[] args) throws IOException {
        List<String> lines = Files.readAllLines(Path.of("input.txt"));
        lines = lines.stream().filter(s -> {
            char[] chars = s.toCharArray();
            for (int j = 1; j < chars.length; j ++) {
                char c1 = chars[j - 1];
                char c2 = chars[j];

                for (int i = j + 2; i < chars.length; i++) {
                    if (chars[i - 1] == c1 && chars[i] == c2) {
                        return true;
                    }
                }
            }
            return false;
        }).filter(s -> {
            char[] chars = s.toCharArray();
            for (int j = 0; j < chars.length - 2; j ++) {
                if (chars[j] == chars[j + 2]) {
                    return true;
                }
            }
            return false;
        }).collect(Collectors.toList());

        System.out.println(lines.size());
    }
}
