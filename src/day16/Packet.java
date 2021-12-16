package day16;

import java.util.List;
import java.util.OptionalLong;

public class Packet {
    private final int version;
    private final int typeId;
    private final Long value; // -1 if empty
    private final List<Packet> subPackets; // null if empty

    public Packet(int version, int typeId, Long value, List<Packet> subPackets) {
        this.version = version;
        this.typeId = typeId;
        this.value = value;
        this.subPackets = subPackets;
    }

    // Heavily inspired from: https://github.com/topaz
    public long calculate() {
        switch (typeId) {
            case 0:
                return subPackets.stream().mapToLong(Packet::calculate).sum();
            case 1:
                return subPackets.stream().mapToLong(Packet::calculate).reduce(1L, (a, b) -> a * b);
            case 2: {
                OptionalLong result = subPackets.stream().mapToLong(Packet::calculate).min();
                if (result.isPresent()) return result.getAsLong();
            }
            case 3: {
                OptionalLong result = subPackets.stream().mapToLong(Packet::calculate).max();
                if (result.isPresent()) return result.getAsLong();
            }
            case 4:
                return value;
            case 5:
                return (subPackets.get(0).calculate() > subPackets.get(1).calculate()) ? 1 : 0;
            case 6:
                return (subPackets.get(0).calculate() < subPackets.get(1).calculate()) ? 1 : 0;
            case 7:
                return (subPackets.get(0).calculate() == subPackets.get(1).calculate()) ? 1 : 0;
            default:
                throw new IllegalStateException("TypeId not found: " + typeId);
        }
    }

    public int sumOfVersions() {
        if (value != -1) {
            return version;
        } else {
            return version + subPackets.stream().mapToInt(Packet::sumOfVersions).sum();
        }
    }

}
