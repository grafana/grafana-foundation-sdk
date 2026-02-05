<?php

namespace Grafana\Foundation\Resource;

class Metadata implements \JsonSerializable
{
    public string $name;

    public ?string $namespace;

    /**
     * @var array<string, string>|null
     */
    public ?array $labels;

    /**
     * @var array<string, string>|null
     */
    public ?array $annotations;

    public ?string $uid;

    public ?string $resourceVersion;

    public ?int $generation;

    public ?string $creationTimestamp;

    public ?string $updateTimestamp;

    public ?string $deletionTimestamp;

    /**
     * @param string|null $name
     * @param string|null $namespace
     * @param array<string, string>|null $labels
     * @param array<string, string>|null $annotations
     * @param string|null $uid
     * @param string|null $resourceVersion
     * @param int|null $generation
     * @param string|null $creationTimestamp
     * @param string|null $updateTimestamp
     * @param string|null $deletionTimestamp
     */
    public function __construct(?string $name = null, ?string $namespace = null, ?array $labels = null, ?array $annotations = null, ?string $uid = null, ?string $resourceVersion = null, ?int $generation = null, ?string $creationTimestamp = null, ?string $updateTimestamp = null, ?string $deletionTimestamp = null)
    {
        $this->name = $name ?: "";
        $this->namespace = $namespace;
        $this->labels = $labels;
        $this->annotations = $annotations;
        $this->uid = $uid;
        $this->resourceVersion = $resourceVersion;
        $this->generation = $generation;
        $this->creationTimestamp = $creationTimestamp;
        $this->updateTimestamp = $updateTimestamp;
        $this->deletionTimestamp = $deletionTimestamp;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{name?: string, namespace?: string, labels?: array<string, string>, annotations?: array<string, string>, uid?: string, resourceVersion?: string, generation?: int, creationTimestamp?: string, updateTimestamp?: string, deletionTimestamp?: string} $inputData */
        $data = $inputData;
        return new self(
            name: $data["name"] ?? null,
            namespace: $data["namespace"] ?? null,
            labels: $data["labels"] ?? null,
            annotations: $data["annotations"] ?? null,
            uid: $data["uid"] ?? null,
            resourceVersion: $data["resourceVersion"] ?? null,
            generation: $data["generation"] ?? null,
            creationTimestamp: $data["creationTimestamp"] ?? null,
            updateTimestamp: $data["updateTimestamp"] ?? null,
            deletionTimestamp: $data["deletionTimestamp"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->name = $this->name;
        if (isset($this->namespace)) {
            $data->namespace = $this->namespace;
        }
        if (isset($this->labels)) {
            $data->labels = $this->labels;
        }
        if (isset($this->annotations)) {
            $data->annotations = $this->annotations;
        }
        if (isset($this->uid)) {
            $data->uid = $this->uid;
        }
        if (isset($this->resourceVersion)) {
            $data->resourceVersion = $this->resourceVersion;
        }
        if (isset($this->generation)) {
            $data->generation = $this->generation;
        }
        if (isset($this->creationTimestamp)) {
            $data->creationTimestamp = $this->creationTimestamp;
        }
        if (isset($this->updateTimestamp)) {
            $data->updateTimestamp = $this->updateTimestamp;
        }
        if (isset($this->deletionTimestamp)) {
            $data->deletionTimestamp = $this->deletionTimestamp;
        }
        return $data;
    }
}
