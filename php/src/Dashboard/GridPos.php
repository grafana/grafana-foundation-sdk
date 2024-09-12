<?php

namespace Grafana\Foundation\Dashboard;

/**
 * Position and dimensions of a panel in the grid
 */
class GridPos implements \JsonSerializable
{
    /**
     * Panel height. The height is the number of rows from the top edge of the panel.
     */
    public int $h;

    /**
     * Panel width. The width is the number of columns from the left edge of the panel.
     */
    public int $w;

    /**
     * Panel x. The x coordinate is the number of columns from the left edge of the grid
     */
    public int $x;

    /**
     * Panel y. The y coordinate is the number of rows from the top edge of the grid
     */
    public int $y;

    /**
     * Whether the panel is fixed within the grid. If true, the panel will not be affected by other panels' interactions
     */
    public ?bool $static;

    /**
     * @param int|null $h
     * @param int|null $w
     * @param int|null $x
     * @param int|null $y
     * @param bool|null $static
     */
    public function __construct(?int $h = null, ?int $w = null, ?int $x = null, ?int $y = null, ?bool $static = null)
    {
        $this->h = $h ?: 9;
        $this->w = $w ?: 12;
        $this->x = $x ?: 0;
        $this->y = $y ?: 0;
        $this->static = $static;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{h?: int, w?: int, x?: int, y?: int, static?: bool} $inputData */
        $data = $inputData;
        return new self(
            h: $data["h"] ?? null,
            w: $data["w"] ?? null,
            x: $data["x"] ?? null,
            y: $data["y"] ?? null,
            static: $data["static"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "h" => $this->h,
            "w" => $this->w,
            "x" => $this->x,
            "y" => $this->y,
        ];
        if (isset($this->static)) {
            $data["static"] = $this->static;
        }
        return $data;
    }
}
