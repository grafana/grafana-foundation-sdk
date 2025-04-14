<?php

namespace Grafana\Foundation\Azuremonitor;

class BuilderQueryEditorColumnsExpression implements \JsonSerializable
{
    /**
     * @var array<string>|null
     */
    public ?array $columns;

    public \Grafana\Foundation\Azuremonitor\BuilderQueryEditorExpressionType $type;

    /**
     * @param array<string>|null $columns
     * @param \Grafana\Foundation\Azuremonitor\BuilderQueryEditorExpressionType|null $type
     */
    public function __construct(?array $columns = null, ?\Grafana\Foundation\Azuremonitor\BuilderQueryEditorExpressionType $type = null)
    {
        $this->columns = $columns;
        $this->type = $type ?: \Grafana\Foundation\Azuremonitor\BuilderQueryEditorExpressionType::Property();
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{columns?: array<string>, type?: string} $inputData */
        $data = $inputData;
        return new self(
            columns: $data["columns"] ?? null,
            type: isset($data["type"]) ? (function($input) { return \Grafana\Foundation\Azuremonitor\BuilderQueryEditorExpressionType::fromValue($input); })($data["type"]) : null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "type" => $this->type,
        ];
        if (isset($this->columns)) {
            $data["columns"] = $this->columns;
        }
        return $data;
    }
}
