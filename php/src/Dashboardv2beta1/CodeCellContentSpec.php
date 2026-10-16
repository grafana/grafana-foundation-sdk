<?php

namespace Grafana\Foundation\Dashboardv2beta1;

class CodeCellContentSpec implements \JsonSerializable
{
    public string $language;

    public string $code;

    /**
     * @var array<int>|null
     */
    public ?array $highlight;

    public ?string $annotation;

    /**
     * @param string|null $language
     * @param string|null $code
     * @param array<int>|null $highlight
     * @param string|null $annotation
     */
    public function __construct(?string $language = null, ?string $code = null, ?array $highlight = null, ?string $annotation = null)
    {
        $this->language = $language ?: "";
        $this->code = $code ?: "";
        $this->highlight = $highlight;
        $this->annotation = $annotation;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{language?: string, code?: string, highlight?: array<int>, annotation?: string} $inputData */
        $data = $inputData;
        return new self(
            language: $data["language"] ?? null,
            code: $data["code"] ?? null,
            highlight: $data["highlight"] ?? null,
            annotation: $data["annotation"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->language = $this->language;
        $data->code = $this->code;
        if (isset($this->highlight)) {
            $data->highlight = $this->highlight;
        }
        if (isset($this->annotation)) {
            $data->annotation = $this->annotation;
        }
        return $data;
    }
}
