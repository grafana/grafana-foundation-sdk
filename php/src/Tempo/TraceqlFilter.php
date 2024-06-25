<?php

namespace Grafana\Foundation\Tempo;

class TraceqlFilter implements \JsonSerializable
{
    /**
     * Uniquely identify the filter, will not be used in the query generation
     */
    public string $id;

    /**
     * The tag for the search filter, for example: .http.status_code, .service.name, status
     */
    public ?string $tag;

    /**
     * The operator that connects the tag to the value, for example: =, >, !=, =~
     */
    public ?string $operator;

    /**
     * The value for the search filter
     * @var string|array<string>|null
     */
    public $value;

    /**
     * The type of the value, used for example to check whether we need to wrap the value in quotes when generating the query
     */
    public ?string $valueType;

    /**
     * The scope of the filter, can either be unscoped/all scopes, resource or span
     */
    public ?\Grafana\Foundation\Tempo\TraceqlSearchScope $scope;

    /**
     * @param string|null $id
     * @param string|null $tag
     * @param string|null $operator
     * @param string|array<string>|null $value
     * @param string|null $valueType
     * @param \Grafana\Foundation\Tempo\TraceqlSearchScope|null $scope
     */
    public function __construct(?string $id = null, ?string $tag = null, ?string $operator = null,  $value = null, ?string $valueType = null, ?\Grafana\Foundation\Tempo\TraceqlSearchScope $scope = null)
    {
        $this->id = $id ?: "";
        $this->tag = $tag;
        $this->operator = $operator;
        $this->value = $value;
        $this->valueType = $valueType;
        $this->scope = $scope;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{id?: string, tag?: string, operator?: string, value?: string|array<string>, valueType?: string, scope?: string} $inputData */
        $data = $inputData;
        return new self(
            id: $data["id"] ?? null,
            tag: $data["tag"] ?? null,
            operator: $data["operator"] ?? null,
            value: isset($data["value"]) ? (function($input) {
        switch (true) {
        case is_string($input):
            return $input;
        default:
            return $input;
    }
    })($data["value"]) : null,
            valueType: $data["valueType"] ?? null,
            scope: isset($data["scope"]) ? (function($input) { return \Grafana\Foundation\Tempo\TraceqlSearchScope::fromValue($input); })($data["scope"]) : null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "id" => $this->id,
        ];
        if (isset($this->tag)) {
            $data["tag"] = $this->tag;
        }
        if (isset($this->operator)) {
            $data["operator"] = $this->operator;
        }
        if (isset($this->value)) {
            $data["value"] = $this->value;
        }
        if (isset($this->valueType)) {
            $data["valueType"] = $this->valueType;
        }
        if (isset($this->scope)) {
            $data["scope"] = $this->scope;
        }
        return $data;
    }
}
