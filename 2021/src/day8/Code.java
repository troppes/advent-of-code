package day8;

import java.util.List;
import java.util.Objects;
import java.util.stream.Collectors;

public class Code implements Comparable<Code> {
    private final List<Character> chars;
    private final String code;

    Code(String code){
        this.code = code;
        chars = code.chars().mapToObj((i) -> (char) i).collect(Collectors.toList());
    }

    public List<Character> getChars() {
        return chars;
    }

    public String getCode() {
        return code;
    }

    boolean contains(Code c){
        return chars.containsAll(c.getChars());
    }

    int length(){
        return code.length();
    }

    @Override
    public int compareTo(Code o) {
        if(o == null) throw new NullPointerException();
        return Integer.compare(code.length(), o.getCode().length());
    }

    public boolean intersect(Code segment, int limit) {
        int counter = 0;
        for (Character c1: chars) {
            for(Character c2: segment.getChars()){
                if(c1 == c2) counter++;
            }
        }
        return counter == limit;
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        return contains((Code) o) && length() == ((Code) o).length();
    }

    @Override
    public int hashCode() {
        return Objects.hash(code);
    }
}
