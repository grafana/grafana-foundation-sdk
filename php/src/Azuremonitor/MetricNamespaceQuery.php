<?php

namespace Grafana\Foundation\Azuremonitor;

class MetricNamespaceQuery implements \JsonSerializable
{
    public ?string $rawQuery;

    public string $kind;

    public string $subscription;

    public string $resourceGroup;

    public ?string $metricNamespace;

    public ?string $resourceName;

    /**
     * @param string|null $rawQuery
     * @param string|null $subscription
     * @param string|null $resourceGroup
     * @param string|null $metricNamespace
     * @param string|null $resourceName
     */
    public function __construct(?string $rawQuery = null, ?string $subscription = null, ?string $resourceGroup = null, ?string $metricNamespace = null, ?string $resourceName = null)
    {
        $this->rawQuery = $rawQuery;
        $this->kind = "MetricNamespaceQuery";
    
        $this->subscription = $subscription ?: "";
        $this->resourceGroup = $resourceGroup ?: "";
        $this->metricNamespace = $metricNamespace;
        $this->resourceName = $resourceName;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{rawQuery?: string, kind?: string, subscription?: string, resourceGroup?: string, metricNamespace?: string, resourceName?: string} $inputData */
        $data = $inputData;
        return new self(
            rawQuery: $data["rawQuery"] ?? null,
            subscription: $data["subscription"] ?? null,
            resourceGroup: $data["resourceGroup"] ?? null,
            metricNamespace: $data["metricNamespace"] ?? null,
            resourceName: $data["resourceName"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "kind" => $this->kind,
            "subscription" => $this->subscription,
            "resourceGroup" => $this->resourceGroup,
        ];
        if (isset($this->rawQuery)) {
            $data["rawQuery"] = $this->rawQuery;
        }
        if (isset($this->metricNamespace)) {
            $data["metricNamespace"] = $this->metricNamespace;
        }
        if (isset($this->resourceName)) {
            $data["resourceName"] = $this->resourceName;
        }
        return $data;
    }
}
