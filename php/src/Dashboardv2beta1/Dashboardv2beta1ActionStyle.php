<?php

namespace Grafana\Foundation\Dashboardv2beta1;

class Dashboardv2beta1ActionStyle implements \JsonSerializable
{
    public ?string $backgroundColor;

    /**
     * @param string|null $backgroundColor
     */
    public function __construct(?string $backgroundColor = null)
    {
        $this->backgroundColor = $backgroundColor;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{backgroundColor?: string} $inputData */
        $data = $inputData;
        return new self(
            backgroundColor: $data["backgroundColor"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        if (isset($this->backgroundColor)) {
            $data->backgroundColor = $this->backgroundColor;
        }
        return $data;
    }
}
