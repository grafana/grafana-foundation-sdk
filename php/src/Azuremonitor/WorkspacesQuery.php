<?php

namespace Grafana\Foundation\Azuremonitor;

class WorkspacesQuery implements \JsonSerializable
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
        $this->kind = "WorkspacesQuery";
    
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
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->kind = $this->kind;
        $data->subscription = $this->subscription;
        if (isset($this->rawQuery)) {
            $data->rawQuery = $this->rawQuery;
        }
        return $data;
    }
}
