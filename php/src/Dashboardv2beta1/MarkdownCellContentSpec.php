<?php

namespace Grafana\Foundation\Dashboardv2beta1;

class MarkdownCellContentSpec implements \JsonSerializable
{
    public string $text;

    /**
     * @param string|null $text
     */
    public function __construct(?string $text = null)
    {
        $this->text = $text ?: "";
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{text?: string} $inputData */
        $data = $inputData;
        return new self(
            text: $data["text"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->text = $this->text;
        return $data;
    }
}
