<?php

namespace Grafana\Foundation\Dashboard;

/**
 * Annotation event field mapping. Defines how to map a data frame field to an annotation event field.
 */
class AnnotationEventFieldMapping implements \JsonSerializable
{
    /**
     * Source type for the field value.
     */
    public ?\Grafana\Foundation\Dashboard\AnnotationEventFieldSource $source;

    /**
     * Constant value to use when source is "text".
     */
    public ?string $value;

    /**
     * Regular expression to apply to the field value.
     */
    public ?string $regex;

    /**
     * @param \Grafana\Foundation\Dashboard\AnnotationEventFieldSource|null $source
     * @param string|null $value
     * @param string|null $regex
     */
    public function __construct(?\Grafana\Foundation\Dashboard\AnnotationEventFieldSource $source = null, ?string $value = null, ?string $regex = null)
    {
        $this->source = $source;
        $this->value = $value;
        $this->regex = $regex;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{source?: string, value?: string, regex?: string} $inputData */
        $data = $inputData;
        return new self(
            source: isset($data["source"]) ? (function($input) { return \Grafana\Foundation\Dashboard\AnnotationEventFieldSource::fromValue($input); })($data["source"]) : null,
            value: $data["value"] ?? null,
            regex: $data["regex"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        if (isset($this->source)) {
            $data->source = $this->source;
        }
        if (isset($this->value)) {
            $data->value = $this->value;
        }
        if (isset($this->regex)) {
            $data->regex = $this->regex;
        }
        return $data;
    }
}
