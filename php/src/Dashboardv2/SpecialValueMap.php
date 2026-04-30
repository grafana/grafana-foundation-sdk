<?php

namespace Grafana\Foundation\Dashboardv2;

/**
 * Maps special values like Null, NaN (not a number), and boolean values like true and false to a display text and color.
 * See SpecialValueMatch to see the list of special values.
 * For example, you can configure a special value mapping so that null values appear as N/A.
 */
class SpecialValueMap implements \JsonSerializable
{
    public \Grafana\Foundation\Dashboardv2\MappingType $type;

    public \Grafana\Foundation\Dashboardv2\Dashboardv2SpecialValueMapOptions $options;

    /**
     * @param \Grafana\Foundation\Dashboardv2\Dashboardv2SpecialValueMapOptions|null $options
     */
    public function __construct(?\Grafana\Foundation\Dashboardv2\Dashboardv2SpecialValueMapOptions $options = null)
    {
        $this->type = \Grafana\Foundation\Dashboardv2\MappingType::value();
        $this->options = $options ?: new \Grafana\Foundation\Dashboardv2\Dashboardv2SpecialValueMapOptions();
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
    	return \Grafana\Foundation\Dashboardv2\Dashboardv2SpecialValueMapOptions::fromArray($val);
    })($data["options"]) : null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->type = $this->type;
        $data->options = $this->options;
        return $data;
    }
}
