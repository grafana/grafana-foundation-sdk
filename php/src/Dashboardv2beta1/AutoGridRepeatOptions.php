<?php

namespace Grafana\Foundation\Dashboardv2beta1;

class AutoGridRepeatOptions implements \JsonSerializable
{
    public \Grafana\Foundation\Dashboardv2beta1\RepeatMode $mode;

    public string $value;

    /**
     * @param string|null $value
     */
    public function __construct(?string $value = null)
    {
        $this->mode = \Grafana\Foundation\Dashboardv2beta1\RepeatMode::variable();
        $this->value = $value ?: "";
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{mode?: "variable", value?: string} $inputData */
        $data = $inputData;
        return new self(
            value: $data["value"] ?? null,
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
        return $data;
    }
}
