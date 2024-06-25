<?php

namespace Grafana\Foundation\Azuremonitor;

class ResourceNamesQuery implements \JsonSerializable
{
    public ?string $rawQuery;

    public string $kind;

    public string $subscription;

    public string $resourceGroup;

    public string $metricNamespace;

    /**
     * @param string|null $rawQuery
     * @param string|null $subscription
     * @param string|null $resourceGroup
     * @param string|null $metricNamespace
     */
    public function __construct(?string $rawQuery = null, ?string $subscription = null, ?string $resourceGroup = null, ?string $metricNamespace = null)
    {
        $this->rawQuery = $rawQuery;
        $this->kind = "ResourceNamesQuery";
    
        $this->subscription = $subscription ?: "";
        $this->resourceGroup = $resourceGroup ?: "";
        $this->metricNamespace = $metricNamespace ?: "";
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{rawQuery?: string, kind?: string, subscription?: string, resourceGroup?: string, metricNamespace?: string} $inputData */
        $data = $inputData;
        return new self(
            rawQuery: $data["rawQuery"] ?? null,
            subscription: $data["subscription"] ?? null,
            resourceGroup: $data["resourceGroup"] ?? null,
            metricNamespace: $data["metricNamespace"] ?? null,
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
            "metricNamespace" => $this->metricNamespace,
        ];
        if (isset($this->rawQuery)) {
            $data["rawQuery"] = $this->rawQuery;
        }
        return $data;
    }
}
