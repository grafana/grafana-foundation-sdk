<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * Custom variable value
 */
class CustomVariableValue implements \JsonSerializable
{
    /**
     * The format name or function used in the expression
     */
    public ?string $formatter;

    /**
     * @param string|null $formatter
     */
    public function __construct(?string $formatter = null)
    {
        $this->formatter = $formatter;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{formatter?: string} $inputData */
        $data = $inputData;
        return new self(
            formatter: $data["formatter"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        if (isset($this->formatter)) {
            $data->formatter = $this->formatter;
        }
        return $data;
    }
}
