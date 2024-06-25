<?php

namespace Grafana\Foundation\Xychart;

class XYDimensionConfig implements \JsonSerializable
{
    public int $frame;

    public ?string $x;

    /**
     * @var array<string>|null
     */
    public ?array $exclude;

    /**
     * @param int|null $frame
     * @param string|null $x
     * @param array<string>|null $exclude
     */
    public function __construct(?int $frame = null, ?string $x = null, ?array $exclude = null)
    {
        $this->frame = $frame ?: 0;
        $this->x = $x;
        $this->exclude = $exclude;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{frame?: int, x?: string, exclude?: array<string>} $inputData */
        $data = $inputData;
        return new self(
            frame: $data["frame"] ?? null,
            x: $data["x"] ?? null,
            exclude: $data["exclude"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "frame" => $this->frame,
        ];
        if (isset($this->x)) {
            $data["x"] = $this->x;
        }
        if (isset($this->exclude)) {
            $data["exclude"] = $this->exclude;
        }
        return $data;
    }
}
