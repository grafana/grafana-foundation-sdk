<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * Maps special values like Null, NaN (not a number), and boolean values like true and false to a display text and color.
 * See SpecialValueMatch to see the list of special values.
 * For example, you can configure a special value mapping so that null values appear as N/A.
 */
class SpecialValueMap implements \JsonSerializable
{
    public \Grafana\Foundation\Dashboardv2beta1\MappingType $type;

    public \Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1SpecialValueMapOptions $options;

    /**
     * @param \Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1SpecialValueMapOptions|null $options
     */
    public function __construct(?\Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1SpecialValueMapOptions $options = null)
    {
        $this->type = \Grafana\Foundation\Dashboardv2beta1\MappingType::value();
        $this->options = $options ?: new \Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1SpecialValueMapOptions();
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
    	return \Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1SpecialValueMapOptions::fromArray($val);
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
