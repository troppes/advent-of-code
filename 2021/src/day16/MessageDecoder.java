package day16;

import java.util.ArrayDeque;
import java.util.ArrayList;
import java.util.Queue;
import java.util.stream.Collectors;

public class MessageDecoder {

    Packet masterPacket;

    MessageDecoder(String hexString) {

        hexString = hexString.replaceAll("0", "0000");
        hexString = hexString.replaceAll("1", "0001");
        hexString = hexString.replaceAll("2", "0010");
        hexString = hexString.replaceAll("3", "0011");
        hexString = hexString.replaceAll("4", "0100");
        hexString = hexString.replaceAll("5", "0101");
        hexString = hexString.replaceAll("6", "0110");
        hexString = hexString.replaceAll("7", "0111");
        hexString = hexString.replaceAll("8", "1000");
        hexString = hexString.replaceAll("9", "1001");
        hexString = hexString.replaceAll("A", "1010");
        hexString = hexString.replaceAll("B", "1011");
        hexString = hexString.replaceAll("C", "1100");
        hexString = hexString.replaceAll("D", "1101");
        hexString = hexString.replaceAll("E", "1110");
        hexString = hexString.replaceAll("F", "1111");

        masterPacket = decode(hexString.chars().mapToObj(c -> (char) c).collect(Collectors.toCollection(ArrayDeque::new)));
    }

    // Heavily inspired from: https://github.com/topaz
    private Packet decode(Queue<Character> queue) {

        int version = bitsToInt(queue, 3);
        int typeId = bitsToInt(queue, 3);

        if (typeId == 4) {
            return new Packet(version, typeId, getLiteralValue(queue), null);
        } else {
            char length = queue.remove();
            ArrayList<Packet> packets = new ArrayList<>();
            if (length == '0') {
                int payloadLength = bitsToInt(queue, 15);
                int bitsToAdd = queue.size() - payloadLength;
                while (queue.size() != bitsToAdd) { // >= are 0
                    packets.add(decode(queue));
                }
            } else if (length == '1') {
                int packetCount = bitsToInt(queue, 11);
                for (int i = 0; i < packetCount; i++) {
                    packets.add(decode(queue));
                }
            }

            return new Packet(version, typeId, -1L, packets);
        }
    }

    private static int bitsToInt(Queue<Character> q, int numberOfBits) {
        StringBuilder s = new StringBuilder();
        for (int i = 0; i < numberOfBits; i++) {
            s.append(q.remove());
        }
        return Integer.parseInt(s.toString(), 2);
    }

    private static long getLiteralValue(Queue<Character> q) {

        StringBuilder s = new StringBuilder();
        while (!String.valueOf(q.remove()).equals("0")) {
            for (int j = 0; j < 4; j++) {
                s.append(q.remove());
            }
        }
        // add the last 4
        for (int j = 0; j < 4; j++) {
            s.append(q.remove());
        }

        return Long.parseLong(s.toString(), 2);
    }
}
