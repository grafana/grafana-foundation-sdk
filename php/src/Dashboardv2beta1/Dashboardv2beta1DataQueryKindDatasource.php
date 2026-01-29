<?php

namespace Grafana\Foundation\Dashboardv2beta1;

class Dashboardv2beta1DataQueryKindDatasource implements \JsonSerializable
{
    public ?string $name;

    /**
     * @param string|null $name
     */
    public function __construct(?string $name = null)
    {
        $this->name = $name;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{name?: string} $inputData */
        $data = $inputData;
        return new self(
            name: $data["name"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        if (isset($this->name)) {
            $data->name = $this->name;
        }
        return $data;
    }
}
