<?php

namespace Grafana\Foundation\Common;

/**
 * TODO docs
 */
class LineConfig implements \JsonSerializable
{
    public ?string $lineColor;

    public ?float $lineWidth;

    public ?\Grafana\Foundation\Common\LineInterpolation $lineInterpolation;

    public ?\Grafana\Foundation\Common\LineStyle $lineStyle;

    /**
     * Indicate if null values should be treated as gaps or connected.
     * When the value is a number, it represents the maximum delta in the
     * X axis that should be considered connected.  For timeseries, this is milliseconds
     * @var bool|float|null
     */
    public $spanNulls;

    /**
     * @param string|null $lineColor
     * @param float|null $lineWidth
     * @param \Grafana\Foundation\Common\LineInterpolation|null $lineInterpolation
     * @param \Grafana\Foundation\Common\LineStyle|null $lineStyle
     * @param bool|float|null $spanNulls
     */
    public function __construct(?string $lineColor = null, ?float $lineWidth = null, ?\Grafana\Foundation\Common\LineInterpolation $lineInterpolation = null, ?\Grafana\Foundation\Common\LineStyle $lineStyle = null,  $spanNulls = null)
    {
        $this->lineColor = $lineColor;
        $this->lineWidth = $lineWidth;
        $this->lineInterpolation = $lineInterpolation;
        $this->lineStyle = $lineStyle;
        $this->spanNulls = $spanNulls;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{lineColor?: string, lineWidth?: float, lineInterpolation?: string, lineStyle?: mixed, spanNulls?: bool|float} $inputData */
        $data = $inputData;
        return new self(
            lineColor: $data["lineColor"] ?? null,
            lineWidth: $data["lineWidth"] ?? null,
            lineInterpolation: isset($data["lineInterpolation"]) ? (function($input) { return \Grafana\Foundation\Common\LineInterpolation::fromValue($input); })($data["lineInterpolation"]) : null,
            lineStyle: isset($data["lineStyle"]) ? (function($input) {
    	/** @var array{fill?: string, dash?: array<float>} */
    $val = $input;
    	return \Grafana\Foundation\Common\LineStyle::fromArray($val);
    })($data["lineStyle"]) : null,
            spanNulls: isset($data["spanNulls"]) ? (function($input) {
        switch (true) {
        case is_bool($input):
            return $input;
        case is_float($input):
            return $input;
        default:
            throw new \ValueError('incorrect value for disjunction');
    }
    })($data["spanNulls"]) : null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
        ];
        if (isset($this->lineColor)) {
            $data["lineColor"] = $this->lineColor;
        }
        if (isset($this->lineWidth)) {
            $data["lineWidth"] = $this->lineWidth;
        }
        if (isset($this->lineInterpolation)) {
            $data["lineInterpolation"] = $this->lineInterpolation;
        }
        if (isset($this->lineStyle)) {
            $data["lineStyle"] = $this->lineStyle;
        }
        if (isset($this->spanNulls)) {
            $data["spanNulls"] = $this->spanNulls;
        }
        return $data;
    }
}
