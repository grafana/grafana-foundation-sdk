<?php

namespace Grafana\Foundation\Common;

/**
 * TODO docs
 */
class OptionsWithTimezones implements \JsonSerializable
{
    /**
     * @var array<string>|null
     */
    public ?array $timezone;

    /**
     * @param array<string>|null $timezone
     */
    public function __construct(?array $timezone = null)
    {
        $this->timezone = $timezone;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{timezone?: array<string>} $inputData */
        $data = $inputData;
        return new self(
            timezone: $data["timezone"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        if (isset($this->timezone)) {
            $data->timezone = $this->timezone;
        }
        return $data;
    }
}
