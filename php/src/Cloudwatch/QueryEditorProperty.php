<?php

namespace Grafana\Foundation\Cloudwatch;

class QueryEditorProperty implements \JsonSerializable
{
    public \Grafana\Foundation\Cloudwatch\QueryEditorPropertyType $type;

    public ?string $name;

    /**
     * @param string|null $name
     */
    public function __construct(?string $name = null)
    {
        $this->type = \Grafana\Foundation\Cloudwatch\QueryEditorPropertyType::string();
        $this->name = $name;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{type?: "string", name?: string} $inputData */
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
        $data->type = $this->type;
        if (isset($this->name)) {
            $data->name = $this->name;
        }
        return $data;
    }
}
