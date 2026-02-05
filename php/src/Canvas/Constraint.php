<?php

namespace Grafana\Foundation\Canvas;

class Constraint implements \JsonSerializable
{
    public ?\Grafana\Foundation\Canvas\HorizontalConstraint $horizontal;

    public ?\Grafana\Foundation\Canvas\VerticalConstraint $vertical;

    /**
     * @param \Grafana\Foundation\Canvas\HorizontalConstraint|null $horizontal
     * @param \Grafana\Foundation\Canvas\VerticalConstraint|null $vertical
     */
    public function __construct(?\Grafana\Foundation\Canvas\HorizontalConstraint $horizontal = null, ?\Grafana\Foundation\Canvas\VerticalConstraint $vertical = null)
    {
        $this->horizontal = $horizontal;
        $this->vertical = $vertical;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{horizontal?: string, vertical?: string} $inputData */
        $data = $inputData;
        return new self(
            horizontal: isset($data["horizontal"]) ? (function($input) { return \Grafana\Foundation\Canvas\HorizontalConstraint::fromValue($input); })($data["horizontal"]) : null,
            vertical: isset($data["vertical"]) ? (function($input) { return \Grafana\Foundation\Canvas\VerticalConstraint::fromValue($input); })($data["vertical"]) : null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        if (isset($this->horizontal)) {
            $data->horizontal = $this->horizontal;
        }
        if (isset($this->vertical)) {
            $data->vertical = $this->vertical;
        }
        return $data;
    }
}
