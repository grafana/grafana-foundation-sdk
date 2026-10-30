<?php

namespace Grafana\Foundation\Alertingv1beta1;

class RuleRef implements \JsonSerializable
{
    /**
     * name is the metadata.name of an AlertRule or RecordingRule resource.
     */
    public string $name;

    /**
     * @param string|null $name
     */
    public function __construct(?string $name = null)
    {
        $this->name = $name ?: "";
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{name?: string} $inputData */
        $data = $inputData;
        return new self(
            name: $data["name"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->name = $this->name;
        return $data;
    }
}
