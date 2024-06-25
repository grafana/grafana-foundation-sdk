<?php

namespace Grafana\Foundation\Common;

/**
 * TODO docs
 */
class ReduceDataOptions implements \JsonSerializable
{
    /**
     * If true show each row value
     */
    public ?bool $values;

    /**
     * if showing all values limit
     */
    public ?float $limit;

    /**
     * When !values, pick one value for the whole field
     * @var array<string>
     */
    public array $calcs;

    /**
     * Which fields to show.  By default this is only numeric fields
     */
    public ?string $fields;

    /**
     * @param bool|null $values
     * @param float|null $limit
     * @param array<string>|null $calcs
     * @param string|null $fields
     */
    public function __construct(?bool $values = null, ?float $limit = null, ?array $calcs = null, ?string $fields = null)
    {
        $this->values = $values;
        $this->limit = $limit;
        $this->calcs = $calcs ?: [];
        $this->fields = $fields;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{values?: bool, limit?: float, calcs?: array<string>, fields?: string} $inputData */
        $data = $inputData;
        return new self(
            values: $data["values"] ?? null,
            limit: $data["limit"] ?? null,
            calcs: $data["calcs"] ?? null,
            fields: $data["fields"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "calcs" => $this->calcs,
        ];
        if (isset($this->values)) {
            $data["values"] = $this->values;
        }
        if (isset($this->limit)) {
            $data["limit"] = $this->limit;
        }
        if (isset($this->fields)) {
            $data["fields"] = $this->fields;
        }
        return $data;
    }
}
