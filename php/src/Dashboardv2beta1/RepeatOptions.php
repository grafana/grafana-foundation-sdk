<?php

namespace Grafana\Foundation\Dashboardv2beta1;

class RepeatOptions implements \JsonSerializable
{
    public \Grafana\Foundation\Dashboardv2beta1\RepeatMode $mode;

    public string $value;

    public ?\Grafana\Foundation\Dashboardv2beta1\RepeatOptionsDirection $direction;

    public ?int $maxPerRow;

    /**
     * @param string|null $value
     * @param \Grafana\Foundation\Dashboardv2beta1\RepeatOptionsDirection|null $direction
     * @param int|null $maxPerRow
     */
    public function __construct(?string $value = null, ?\Grafana\Foundation\Dashboardv2beta1\RepeatOptionsDirection $direction = null, ?int $maxPerRow = null)
    {
        $this->mode = \Grafana\Foundation\Dashboardv2beta1\RepeatMode::variable();
        $this->value = $value ?: "";
        $this->direction = $direction;
        $this->maxPerRow = $maxPerRow;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{mode?: "variable", value?: string, direction?: string, maxPerRow?: int} $inputData */
        $data = $inputData;
        return new self(
            value: $data["value"] ?? null,
            direction: isset($data["direction"]) ? (function($input) { return \Grafana\Foundation\Dashboardv2beta1\RepeatOptionsDirection::fromValue($input); })($data["direction"]) : null,
            maxPerRow: $data["maxPerRow"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->mode = $this->mode;
        $data->value = $this->value;
        if (isset($this->direction)) {
            $data->direction = $this->direction;
        }
        if (isset($this->maxPerRow)) {
            $data->maxPerRow = $this->maxPerRow;
        }
        return $data;
    }
}
