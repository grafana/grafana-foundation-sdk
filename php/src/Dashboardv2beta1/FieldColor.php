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
     * The end color for the gradient color mode (smallest value). Only used when mode is gradient.
     */
    public ?string $gradientColorTo;

    /**
     * Some visualizations need to know how to assign a series color from by value color schemes.
     */
    public ?\Grafana\Foundation\Dashboardv2beta1\FieldColorSeriesByMode $seriesBy;

    /**
     * @param \Grafana\Foundation\Dashboardv2beta1\FieldColorModeId|null $mode
     * @param string|null $fixedColor
     * @param string|null $gradientColorTo
     * @param \Grafana\Foundation\Dashboardv2beta1\FieldColorSeriesByMode|null $seriesBy
     */
    public function __construct(?\Grafana\Foundation\Dashboardv2beta1\FieldColorModeId $mode = null, ?string $fixedColor = null, ?string $gradientColorTo = null, ?\Grafana\Foundation\Dashboardv2beta1\FieldColorSeriesByMode $seriesBy = null)
    {
        $this->mode = $mode ?: \Grafana\Foundation\Dashboardv2beta1\FieldColorModeId::Thresholds();
        $this->fixedColor = $fixedColor;
        $this->gradientColorTo = $gradientColorTo;
        $this->seriesBy = $seriesBy;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{mode?: string, fixedColor?: string, gradientColorTo?: string, seriesBy?: string} $inputData */
        $data = $inputData;
        return new self(
            mode: isset($data["mode"]) ? (function($input) { return \Grafana\Foundation\Dashboardv2beta1\FieldColorModeId::fromValue($input); })($data["mode"]) : null,
            fixedColor: $data["fixedColor"] ?? null,
            gradientColorTo: $data["gradientColorTo"] ?? null,
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
        if (isset($this->gradientColorTo)) {
            $data->gradientColorTo = $this->gradientColorTo;
        }
        if (isset($this->seriesBy)) {
            $data->seriesBy = $this->seriesBy;
        }
        return $data;
    }
}
