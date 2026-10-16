<?php

namespace Grafana\Foundation\Dashboardv2;

class Threshold implements \JsonSerializable
{
    /**
     * Value null means -Infinity
     */
    public ?float $value;

    /**
     * Optional dashboard-variable expression (e.g. `$myVar`) resolved at render time; `value` is the numeric fallback when the expression cannot be resolved to a single finite number.
     */
    public ?string $valueExpr;

    public string $color;

    /**
     * @param float|null $value
     * @param string|null $valueExpr
     * @param string|null $color
     */
    public function __construct(?float $value = null, ?string $valueExpr = null, ?string $color = null)
    {
        $this->value = $value;
        $this->valueExpr = $valueExpr;
        $this->color = $color ?: "";
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{value?: float, valueExpr?: string, color?: string} $inputData */
        $data = $inputData;
        return new self(
            value: $data["value"] ?? null,
            valueExpr: $data["valueExpr"] ?? null,
            color: $data["color"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->color = $this->color;
        if (isset($this->value)) {
            $data->value = $this->value;
        }
        if (isset($this->valueExpr)) {
            $data->valueExpr = $this->valueExpr;
        }
        return $data;
    }
}
