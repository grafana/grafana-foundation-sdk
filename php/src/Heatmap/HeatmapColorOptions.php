<?php

namespace Grafana\Foundation\Heatmap;

/**
 * Controls various color options
 */
class HeatmapColorOptions implements \JsonSerializable
{
    /**
     * Sets the color mode
     */
    public ?\Grafana\Foundation\Heatmap\HeatmapColorMode $mode;

    /**
     * Controls the color scheme used
     */
    public string $scheme;

    /**
     * Controls the color fill when in opacity mode
     */
    public string $fill;

    /**
     * Controls the color scale
     */
    public ?\Grafana\Foundation\Heatmap\HeatmapColorScale $scale;

    /**
     * Controls the exponent when scale is set to exponential
     */
    public float $exponent;

    /**
     * Controls the number of color steps
     */
    public int $steps;

    /**
     * Reverses the color scheme
     */
    public bool $reverse;

    /**
     * Sets the minimum value for the color scale
     */
    public ?float $min;

    /**
     * Sets the maximum value for the color scale
     */
    public ?float $max;

    /**
     * @param \Grafana\Foundation\Heatmap\HeatmapColorMode|null $mode
     * @param string|null $scheme
     * @param string|null $fill
     * @param \Grafana\Foundation\Heatmap\HeatmapColorScale|null $scale
     * @param float|null $exponent
     * @param int|null $steps
     * @param bool|null $reverse
     * @param float|null $min
     * @param float|null $max
     */
    public function __construct(?\Grafana\Foundation\Heatmap\HeatmapColorMode $mode = null, ?string $scheme = null, ?string $fill = null, ?\Grafana\Foundation\Heatmap\HeatmapColorScale $scale = null, ?float $exponent = null, ?int $steps = null, ?bool $reverse = null, ?float $min = null, ?float $max = null)
    {
        $this->mode = $mode;
        $this->scheme = $scheme ?: "";
        $this->fill = $fill ?: "";
        $this->scale = $scale;
        $this->exponent = $exponent ?: 0;
        $this->steps = $steps ?: 0;
        $this->reverse = $reverse ?: false;
        $this->min = $min;
        $this->max = $max;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{mode?: string, scheme?: string, fill?: string, scale?: string, exponent?: float, steps?: int, reverse?: bool, min?: float, max?: float} $inputData */
        $data = $inputData;
        return new self(
            mode: isset($data["mode"]) ? (function($input) { return \Grafana\Foundation\Heatmap\HeatmapColorMode::fromValue($input); })($data["mode"]) : null,
            scheme: $data["scheme"] ?? null,
            fill: $data["fill"] ?? null,
            scale: isset($data["scale"]) ? (function($input) { return \Grafana\Foundation\Heatmap\HeatmapColorScale::fromValue($input); })($data["scale"]) : null,
            exponent: $data["exponent"] ?? null,
            steps: $data["steps"] ?? null,
            reverse: $data["reverse"] ?? null,
            min: $data["min"] ?? null,
            max: $data["max"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "scheme" => $this->scheme,
            "fill" => $this->fill,
            "exponent" => $this->exponent,
            "steps" => $this->steps,
            "reverse" => $this->reverse,
        ];
        if (isset($this->mode)) {
            $data["mode"] = $this->mode;
        }
        if (isset($this->scale)) {
            $data["scale"] = $this->scale;
        }
        if (isset($this->min)) {
            $data["min"] = $this->min;
        }
        if (isset($this->max)) {
            $data["max"] = $this->max;
        }
        return $data;
    }
}
