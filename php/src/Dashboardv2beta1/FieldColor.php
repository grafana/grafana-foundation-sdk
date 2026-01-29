<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * Map a field to a color.
 */
class FieldColor implements \JsonSerializable
{
    /**
     * The main color scheme mode.
     */
    public \Grafana\Foundation\Dashboardv2beta1\FieldColorModeId $mode;

    /**
     * The fixed color value for fixed or shades color modes.
     */
    public ?string $fixedColor;

    /**
     * Some visualizations need to know how to assign a series color from by value color schemes.
     */
    public ?\Grafana\Foundation\Dashboardv2beta1\FieldColorSeriesByMode $seriesBy;

    /**
     * @param \Grafana\Foundation\Dashboardv2beta1\FieldColorModeId|null $mode
     * @param string|null $fixedColor
     * @param \Grafana\Foundation\Dashboardv2beta1\FieldColorSeriesByMode|null $seriesBy
     */
    public function __construct(?\Grafana\Foundation\Dashboardv2beta1\FieldColorModeId $mode = null, ?string $fixedColor = null, ?\Grafana\Foundation\Dashboardv2beta1\FieldColorSeriesByMode $seriesBy = null)
    {
        $this->mode = $mode ?: \Grafana\Foundation\Dashboardv2beta1\FieldColorModeId::Thresholds();
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
            mode: isset($data["mode"]) ? (function($input) { return \Grafana\Foundation\Dashboardv2beta1\FieldColorModeId::fromValue($input); })($data["mode"]) : null,
            fixedColor: $data["fixedColor"] ?? null,
            seriesBy: isset($data["seriesBy"]) ? (function($input) { return \Grafana\Foundation\Dashboardv2beta1\FieldColorSeriesByMode::fromValue($input); })($data["seriesBy"]) : null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->mode = $this->mode;
        if (isset($this->fixedColor)) {
            $data->fixedColor = $this->fixedColor;
        }
        if (isset($this->seriesBy)) {
            $data->seriesBy = $this->seriesBy;
        }
        return $data;
    }
}
