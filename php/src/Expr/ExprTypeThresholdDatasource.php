<?php

namespace Grafana\Foundation\Expr;

class ExprTypeThresholdDatasource implements \JsonSerializable
{
    /**
     * The apiserver version
     */
    public ?string $apiVersion;

    /**
     * The datasource plugin type
     */
    public string $type;

    /**
     * Datasource UID (NOTE: name in k8s)
     */
    public ?string $uid;

    /**
     * @param string|null $apiVersion
     * @param string|null $uid
     */
    public function __construct(?string $apiVersion = null, ?string $uid = null)
    {
        $this->apiVersion = $apiVersion;
        $this->type = "__expr__";
    
        $this->uid = $uid;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{apiVersion?: string, type?: string, uid?: string} $inputData */
        $data = $inputData;
        return new self(
            apiVersion: $data["apiVersion"] ?? null,
            uid: $data["uid"] ?? null,
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
        if (isset($this->apiVersion)) {
            $data["apiVersion"] = $this->apiVersion;
        }
        if (isset($this->uid)) {
            $data["uid"] = $this->uid;
        }
        return $data;
    }
}
