<?php

namespace Grafana\Foundation\Preferences;

class QueryHistoryPreference implements \JsonSerializable
{
    /**
     * one of: '' | 'query' | 'starred';
     */
    public ?string $homeTab;

    /**
     * @param string|null $homeTab
     */
    public function __construct(?string $homeTab = null)
    {
        $this->homeTab = $homeTab;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{homeTab?: string} $inputData */
        $data = $inputData;
        return new self(
            homeTab: $data["homeTab"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
        ];
        if (isset($this->homeTab)) {
            $data["homeTab"] = $this->homeTab;
        }
        return $data;
    }
}
