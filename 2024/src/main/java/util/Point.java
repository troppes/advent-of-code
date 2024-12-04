package util;

public record Point(int x, int y) {

    public Point step(Direction direction) {
        return new Point(
                x + direction.getDx(),
                y + direction.getDy()
        );
    }

    public Point copy() {
        return new Point(x, y);
    }
}

