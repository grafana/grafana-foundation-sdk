<?php

namespace Grafana\Foundation\Resource;

class Manifest implements \JsonSerializable
{
    public string $apiVersion;

    public string $kind;

    public \Grafana\Foundation\Resource\Metadata $metadata;

    /**
     * @var mixed
     */
    public $spec;

    /**
     * @param string|null $apiVersion
     * @param string|null $kind
     * @param \Grafana\Foundation\Resource\Metadata|null $metadata
     * @param mixed|null $spec
     */
    public function __construct(?string $apiVersion = null, ?string $kind = null, ?\Grafana\Foundation\Resource\Metadata $metadata = null,  $spec = null)
    {
        $this->apiVersion = $apiVersion ?: "";
        $this->kind = $kind ?: "";
        $this->metadata = $metadata ?: new \Grafana\Foundation\Resource\Metadata();
        $this->spec = $spec ?: null;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{apiVersion?: string, kind?: string, metadata?: mixed, spec?: mixed} $inputData */
        $data = $inputData;
        return new self(
            apiVersion: $data["apiVersion"] ?? null,
            kind: $data["kind"] ?? null,
            metadata: isset($data["metadata"]) ? (function($input) {
    	/** @var array{name?: string, namespace?: string, labels?: array<string, string>, annotations?: array<string, string>, uid?: string, resourceVersion?: string, generation?: int, creationTimestamp?: string, updateTimestamp?: string, deletionTimestamp?: string} */
    $val = $input;
    	return \Grafana\Foundation\Resource\Metadata::fromArray($val);
    })($data["metadata"]) : null,
            spec: $data["spec"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->apiVersion = $this->apiVersion;
        $data->kind = $this->kind;
        $data->metadata = $this->metadata;
        $data->spec = $this->spec;
        return $data;
    }
}
