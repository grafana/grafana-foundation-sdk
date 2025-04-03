<?php

namespace Grafana\Foundation\Dashboard;

/**
 * Maps special values like Null, NaN (not a number), and boolean values like true and false to a display text and color.
 * See SpecialValueMatch to see the list of special values.
 * For example, you can configure a special value mapping so that null values appear as N/A.
 */
class SpecialValueMap implements \JsonSerializable
{
    public \Grafana\Foundation\Dashboard\MappingType $type;

    public \Grafana\Foundation\Dashboard\DashboardSpecialValueMapOptions $options;

    /**
     * @param \Grafana\Foundation\Dashboard\DashboardSpecialValueMapOptions|null $options
     */
    public function __construct(?\Grafana\Foundation\Dashboard\DashboardSpecialValueMapOptions $options = null)
    {
        $this->type = \Grafana\Foundation\Dashboard\MappingType::valueToText();
        $this->options = $options ?: new \Grafana\Foundation\Dashboard\DashboardSpecialValueMapOptions();
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{type?: "special", options?: mixed} $inputData */
        $data = $inputData;
        return new self(
            options: isset($data["options"]) ? (function($input) {
    	/** @var array{match?: string, result?: mixed} */
    $val = $input;
    	return \Grafana\Foundation\Dashboard\DashboardSpecialValueMapOptions::fromArray($val);
    })($data["options"]) : null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "type" => $this->type,
            "options" => $this->options,
        ];
        return $data;
    }
}
