package day02;

import com.sun.tools.jconsole.JConsoleContext;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.Arrays;
import java.util.List;
import java.util.OptionalInt;

public class day02 {
    public static void main(String[] args) throws IOException {
        List<String> lines = Files.readAllLines(Path.of("input.txt"));
        int paper = 0;
        int ribbon = 0;
        for (String line : lines ) {
            int[] dimensions = Arrays.stream(line.split("x")).mapToInt(Integer::parseInt).toArray();

            int l = dimensions[0];
            int w = dimensions[1];
            int h = dimensions[2];
            // 2*l*w + 2*w*h + 2*h*l
            int s1 = l*w;
            int s2 = w*h;
            int s3 = h*l;
            int slack = Math.min(s3, Math.min(s1, s2));
            paper += 2 * s1 + 2 * s2 + 2 * s3 + slack;

            Arrays.sort(dimensions);
            ribbon += (dimensions[0] + dimensions[0] + dimensions[1] + dimensions[1]) + (l * w * h);
        }
        System.out.println("Part 1 :" + paper);
        System.out.println("Part 2 :" + ribbon);
    }
}
