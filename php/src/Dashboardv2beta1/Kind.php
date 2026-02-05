<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * --- Common types ---
 */
class Kind implements \JsonSerializable
{
    public string $kind;

    /**
     * @var mixed
     */
    public $spec;

    /**
     * @var mixed|null
     */
    public $metadata;

    /**
     * @param string|null $kind
     * @param mixed|null $spec
     * @param mixed|null $metadata
     */
    public function __construct(?string $kind = null,  $spec = null,  $metadata = null)
    {
        $this->kind = $kind ?: "";
        $this->spec = $spec ?: null;
        $this->metadata = $metadata;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{kind?: string, spec?: mixed, metadata?: mixed} $inputData */
        $data = $inputData;
        return new self(
            kind: $data["kind"] ?? null,
            spec: $data["spec"] ?? null,
            metadata: $data["metadata"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->kind = $this->kind;
        $data->spec = $this->spec;
        if (isset($this->metadata)) {
            $data->metadata = $this->metadata;
        }
        return $data;
    }
}
