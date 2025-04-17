<?php

namespace Grafana\Foundation\Common;

/**
 * TODO docs
 */
class PointsConfig implements \JsonSerializable
{
    public ?\Grafana\Foundation\Common\VisibilityMode $showPoints;

    public ?float $pointSize;

    public ?string $pointColor;

    public ?string $pointSymbol;

    /**
     * @param \Grafana\Foundation\Common\VisibilityMode|null $showPoints
     * @param float|null $pointSize
     * @param string|null $pointColor
     * @param string|null $pointSymbol
     */
    public function __construct(?\Grafana\Foundation\Common\VisibilityMode $showPoints = null, ?float $pointSize = null, ?string $pointColor = null, ?string $pointSymbol = null)
    {
        $this->showPoints = $showPoints;
        $this->pointSize = $pointSize;
        $this->pointColor = $pointColor;
        $this->pointSymbol = $pointSymbol;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{showPoints?: string, pointSize?: float, pointColor?: string, pointSymbol?: string} $inputData */
        $data = $inputData;
        return new self(
            showPoints: isset($data["showPoints"]) ? (function($input) { return \Grafana\Foundation\Common\VisibilityMode::fromValue($input); })($data["showPoints"]) : null,
            pointSize: $data["pointSize"] ?? null,
            pointColor: $data["pointColor"] ?? null,
            pointSymbol: $data["pointSymbol"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        if (isset($this->showPoints)) {
            $data->showPoints = $this->showPoints;
        }
        if (isset($this->pointSize)) {
            $data->pointSize = $this->pointSize;
        }
        if (isset($this->pointColor)) {
            $data->pointColor = $this->pointColor;
        }
        if (isset($this->pointSymbol)) {
            $data->pointSymbol = $this->pointSymbol;
        }
        return $data;
    }
}
