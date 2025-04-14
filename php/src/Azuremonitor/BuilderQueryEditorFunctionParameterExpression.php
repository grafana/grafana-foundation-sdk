<?php

namespace Grafana\Foundation\Azuremonitor;

class BuilderQueryEditorFunctionParameterExpression implements \JsonSerializable
{
    public string $value;

    public \Grafana\Foundation\Azuremonitor\BuilderQueryEditorPropertyType $fieldType;

    public \Grafana\Foundation\Azuremonitor\BuilderQueryEditorExpressionType $type;

    /**
     * @param string|null $value
     * @param \Grafana\Foundation\Azuremonitor\BuilderQueryEditorPropertyType|null $fieldType
     * @param \Grafana\Foundation\Azuremonitor\BuilderQueryEditorExpressionType|null $type
     */
    public function __construct(?string $value = null, ?\Grafana\Foundation\Azuremonitor\BuilderQueryEditorPropertyType $fieldType = null, ?\Grafana\Foundation\Azuremonitor\BuilderQueryEditorExpressionType $type = null)
    {
        $this->value = $value ?: "";
        $this->fieldType = $fieldType ?: \Grafana\Foundation\Azuremonitor\BuilderQueryEditorPropertyType::Number();
        $this->type = $type ?: \Grafana\Foundation\Azuremonitor\BuilderQueryEditorExpressionType::Property();
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{value?: string, fieldType?: string, type?: string} $inputData */
        $data = $inputData;
        return new self(
            value: $data["value"] ?? null,
            fieldType: isset($data["fieldType"]) ? (function($input) { return \Grafana\Foundation\Azuremonitor\BuilderQueryEditorPropertyType::fromValue($input); })($data["fieldType"]) : null,
            type: isset($data["type"]) ? (function($input) { return \Grafana\Foundation\Azuremonitor\BuilderQueryEditorExpressionType::fromValue($input); })($data["type"]) : null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "value" => $this->value,
            "fieldType" => $this->fieldType,
            "type" => $this->type,
        ];
        return $data;
    }
}
