<?php

namespace Grafana\Foundation\Azuremonitor;

class AzureMonitorResource implements \JsonSerializable
{
    public ?string $subscription;

    public ?string $resourceGroup;

    public ?string $resourceName;

    public ?string $metricNamespace;

    public ?string $region;

    /**
     * @param string|null $subscription
     * @param string|null $resourceGroup
     * @param string|null $resourceName
     * @param string|null $metricNamespace
     * @param string|null $region
     */
    public function __construct(?string $subscription = null, ?string $resourceGroup = null, ?string $resourceName = null, ?string $metricNamespace = null, ?string $region = null)
    {
        $this->subscription = $subscription;
        $this->resourceGroup = $resourceGroup;
        $this->resourceName = $resourceName;
        $this->metricNamespace = $metricNamespace;
        $this->region = $region;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{subscription?: string, resourceGroup?: string, resourceName?: string, metricNamespace?: string, region?: string} $inputData */
        $data = $inputData;
        return new self(
            subscription: $data["subscription"] ?? null,
            resourceGroup: $data["resourceGroup"] ?? null,
            resourceName: $data["resourceName"] ?? null,
            metricNamespace: $data["metricNamespace"] ?? null,
            region: $data["region"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        if (isset($this->subscription)) {
            $data->subscription = $this->subscription;
        }
        if (isset($this->resourceGroup)) {
            $data->resourceGroup = $this->resourceGroup;
        }
        if (isset($this->resourceName)) {
            $data->resourceName = $this->resourceName;
        }
        if (isset($this->metricNamespace)) {
            $data->metricNamespace = $this->metricNamespace;
        }
        if (isset($this->region)) {
            $data->region = $this->region;
        }
        return $data;
    }
}
