<?php

namespace Grafana\Foundation\Azuremonitor;

class BuilderQueryEditorOperator implements \JsonSerializable
{
    public string $name;

    public string $value;

    public ?string $labelValue;

    /**
     * @param string|null $name
     * @param string|null $value
     * @param string|null $labelValue
     */
    public function __construct(?string $name = null, ?string $value = null, ?string $labelValue = null)
    {
        $this->name = $name ?: "";
        $this->value = $value ?: "";
        $this->labelValue = $labelValue;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{name?: string, value?: string, labelValue?: string} $inputData */
        $data = $inputData;
        return new self(
            name: $data["name"] ?? null,
            value: $data["value"] ?? null,
            labelValue: $data["labelValue"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->name = $this->name;
        $data->value = $this->value;
        if (isset($this->labelValue)) {
            $data->labelValue = $this->labelValue;
        }
        return $data;
    }
}
