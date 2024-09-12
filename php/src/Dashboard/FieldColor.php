<?php

namespace Grafana\Foundation\Dashboard;

/**
 * Map a field to a color.
 */
class FieldColor implements \JsonSerializable
{
    /**
     * The main color scheme mode.
     */
    public \Grafana\Foundation\Dashboard\FieldColorModeId $mode;

    /**
     * The fixed color value for fixed or shades color modes.
     */
    public ?string $fixedColor;

    /**
     * Some visualizations need to know how to assign a series color from by value color schemes.
     */
    public ?\Grafana\Foundation\Dashboard\FieldColorSeriesByMode $seriesBy;

    /**
     * @param \Grafana\Foundation\Dashboard\FieldColorModeId|null $mode
     * @param string|null $fixedColor
     * @param \Grafana\Foundation\Dashboard\FieldColorSeriesByMode|null $seriesBy
     */
    public function __construct(?\Grafana\Foundation\Dashboard\FieldColorModeId $mode = null, ?string $fixedColor = null, ?\Grafana\Foundation\Dashboard\FieldColorSeriesByMode $seriesBy = null)
    {
        $this->mode = $mode ?: \Grafana\Foundation\Dashboard\FieldColorModeId::Thresholds();
        $this->fixedColor = $fixedColor;
        $this->seriesBy = $seriesBy;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{mode?: string, fixedColor?: string, seriesBy?: string} $inputData */
        $data = $inputData;
        return new self(
            mode: isset($data["mode"]) ? (function($input) { return \Grafana\Foundation\Dashboard\FieldColorModeId::fromValue($input); })($data["mode"]) : null,
            fixedColor: $data["fixedColor"] ?? null,
            seriesBy: isset($data["seriesBy"]) ? (function($input) { return \Grafana\Foundation\Dashboard\FieldColorSeriesByMode::fromValue($input); })($data["seriesBy"]) : null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "mode" => $this->mode,
        ];
        if (isset($this->fixedColor)) {
            $data["fixedColor"] = $this->fixedColor;
        }
        if (isset($this->seriesBy)) {
            $data["seriesBy"] = $this->seriesBy;
        }
        return $data;
    }
}
