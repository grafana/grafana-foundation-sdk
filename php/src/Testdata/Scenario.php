<?php

namespace Grafana\Foundation\Testdata;

/**
 * TODO: Should this live here given it's not used in the dataquery?
 */
class Scenario implements \JsonSerializable
{
    public string $id;

    public string $name;

    public string $stringInput;

    public ?string $description;

    public ?bool $hideAliasField;

    /**
     * @param string|null $id
     * @param string|null $name
     * @param string|null $stringInput
     * @param string|null $description
     * @param bool|null $hideAliasField
     */
    public function __construct(?string $id = null, ?string $name = null, ?string $stringInput = null, ?string $description = null, ?bool $hideAliasField = null)
    {
        $this->id = $id ?: "";
        $this->name = $name ?: "";
        $this->stringInput = $stringInput ?: "";
        $this->description = $description;
        $this->hideAliasField = $hideAliasField;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{id?: string, name?: string, stringInput?: string, description?: string, hideAliasField?: bool} $inputData */
        $data = $inputData;
        return new self(
            id: $data["id"] ?? null,
            name: $data["name"] ?? null,
            stringInput: $data["stringInput"] ?? null,
            description: $data["description"] ?? null,
            hideAliasField: $data["hideAliasField"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->id = $this->id;
        $data->name = $this->name;
        $data->stringInput = $this->stringInput;
        if (isset($this->description)) {
            $data->description = $this->description;
        }
        if (isset($this->hideAliasField)) {
            $data->hideAliasField = $this->hideAliasField;
        }
        return $data;
    }
}
