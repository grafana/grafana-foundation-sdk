<?php

namespace Grafana\Foundation\Azuremonitor;

class ResourceGroupsQuery implements \JsonSerializable
{
    public ?string $rawQuery;

    public string $kind;

    public string $subscription;

    /**
     * @param string|null $rawQuery
     * @param string|null $subscription
     */
    public function __construct(?string $rawQuery = null, ?string $subscription = null)
    {
        $this->rawQuery = $rawQuery;
        $this->kind = "ResourceGroupsQuery";
    
        $this->subscription = $subscription ?: "";
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{rawQuery?: string, kind?: string, subscription?: string} $inputData */
        $data = $inputData;
        return new self(
            rawQuery: $data["rawQuery"] ?? null,
            subscription: $data["subscription"] ?? null,
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
        ];
        if (isset($this->rawQuery)) {
            $data["rawQuery"] = $this->rawQuery;
        }
        return $data;
    }
}
