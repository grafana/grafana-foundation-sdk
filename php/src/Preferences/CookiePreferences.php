<?php

namespace Grafana\Foundation\Preferences;

class CookiePreferences implements \JsonSerializable
{
    /**
     * @var mixed|null
     */
    public $analytics;

    /**
     * @var mixed|null
     */
    public $performance;

    /**
     * @var mixed|null
     */
    public $functional;

    /**
     * @param mixed|null $analytics
     * @param mixed|null $performance
     * @param mixed|null $functional
     */
    public function __construct( $analytics = null,  $performance = null,  $functional = null)
    {
        $this->analytics = $analytics;
        $this->performance = $performance;
        $this->functional = $functional;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{analytics?: mixed, performance?: mixed, functional?: mixed} $inputData */
        $data = $inputData;
        return new self(
            analytics: $data["analytics"] ?? null,
            performance: $data["performance"] ?? null,
            functional: $data["functional"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        if (isset($this->analytics)) {
            $data->analytics = $this->analytics;
        }
        if (isset($this->performance)) {
            $data->performance = $this->performance;
        }
        if (isset($this->functional)) {
            $data->functional = $this->functional;
        }
        return $data;
    }
}
