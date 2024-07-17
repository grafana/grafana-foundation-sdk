<?php

namespace Grafana\Foundation\Preferences;

class NavbarPreference implements \JsonSerializable
{
    /**
     * @var array<string>
     */
    public array $savedItemIds;

    /**
     * @param array<string>|null $savedItemIds
     */
    public function __construct(?array $savedItemIds = null)
    {
        $this->savedItemIds = $savedItemIds ?: [];
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{savedItemIds?: array<string>} $inputData */
        $data = $inputData;
        return new self(
            savedItemIds: $data["savedItemIds"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "savedItemIds" => $this->savedItemIds,
        ];
        return $data;
    }
}
