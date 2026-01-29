<?php

namespace Grafana\Foundation\Dashboardv2beta1;

class DynamicConfigValue implements \JsonSerializable
{
    public string $id;

    /**
     * @var mixed|null
     */
    public $value;

    /**
     * @param string|null $id
     * @param mixed|null $value
     */
    public function __construct(?string $id = null,  $value = null)
    {
        $this->id = $id ?: "";
        $this->value = $value;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{id?: string, value?: mixed} $inputData */
        $data = $inputData;
        return new self(
            id: $data["id"] ?? null,
            value: $data["value"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->id = $this->id;
        if (isset($this->value)) {
            $data->value = $this->value;
        }
        return $data;
    }
}
