<?php

namespace Grafana\Foundation\Expr;

class ExprTypeSqlResultAssertions implements \JsonSerializable
{
    /**
     * Maximum frame count
     */
    public ?int $maxFrames;

    /**
     * Type asserts that the frame matches a known type structure.
     * Possible enum values:
     *  - `""` 
     *  - `"timeseries-wide"` 
     *  - `"timeseries-long"` 
     *  - `"timeseries-many"` 
     *  - `"timeseries-multi"` 
     *  - `"directory-listing"` 
     *  - `"table"` 
     *  - `"numeric-wide"` 
     *  - `"numeric-multi"` 
     *  - `"numeric-long"` 
     *  - `"log-lines"` 
     */
    public ?\Grafana\Foundation\Expr\ExprTypeSqlResultAssertionsType $type;

    /**
     * TypeVersion is the version of the Type property. Versions greater than 0.0 correspond to the dataplane
     * contract documentation https://grafana.github.io/dataplane/contract/.
     * @var array<int>
     */
    public array $typeVersion;

    /**
     * @param int|null $maxFrames
     * @param \Grafana\Foundation\Expr\ExprTypeSqlResultAssertionsType|null $type
     * @param array<int>|null $typeVersion
     */
    public function __construct(?int $maxFrames = null, ?\Grafana\Foundation\Expr\ExprTypeSqlResultAssertionsType $type = null, ?array $typeVersion = null)
    {
        $this->maxFrames = $maxFrames;
        $this->type = $type;
        $this->typeVersion = $typeVersion ?: [];
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{maxFrames?: int, type?: string, typeVersion?: array<int>} $inputData */
        $data = $inputData;
        return new self(
            maxFrames: $data["maxFrames"] ?? null,
            type: isset($data["type"]) ? (function($input) { return \Grafana\Foundation\Expr\ExprTypeSqlResultAssertionsType::fromValue($input); })($data["type"]) : null,
            typeVersion: $data["typeVersion"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "typeVersion" => $this->typeVersion,
        ];
        if (isset($this->maxFrames)) {
            $data["maxFrames"] = $this->maxFrames;
        }
        if (isset($this->type)) {
            $data["type"] = $this->type;
        }
        return $data;
    }
}
