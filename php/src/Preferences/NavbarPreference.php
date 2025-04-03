<?php

namespace Grafana\Foundation\Preferences;

class NavbarPreference implements \JsonSerializable
{
    /**
     * @var array<string>
     */
    public array $bookmarkUrls;

    /**
     * @param array<string>|null $bookmarkUrls
     */
    public function __construct(?array $bookmarkUrls = null)
    {
        $this->bookmarkUrls = $bookmarkUrls ?: [];
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{bookmarkUrls?: array<string>} $inputData */
        $data = $inputData;
        return new self(
            bookmarkUrls: $data["bookmarkUrls"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "bookmarkUrls" => $this->bookmarkUrls,
        ];
        return $data;
    }
}
