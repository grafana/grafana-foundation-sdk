<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * Maps regular expressions to replacement text and a color.
 * For example, if a value is www.example.com, you can configure a regex value mapping so that Grafana displays www and truncates the domain.
 */
class RegexMap implements \JsonSerializable
{
    public \Grafana\Foundation\Dashboardv2beta1\MappingType $type;

    /**
     * Regular expression to match against and the result to apply when the value matches the regex
     */
    public \Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1RegexMapOptions $options;

    /**
     * @param \Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1RegexMapOptions|null $options
     */
    public function __construct(?\Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1RegexMapOptions $options = null)
    {
        $this->type = \Grafana\Foundation\Dashboardv2beta1\MappingType::value();
        $this->options = $options ?: new \Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1RegexMapOptions();
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{type?: "regex", options?: mixed} $inputData */
        $data = $inputData;
        return new self(
            options: isset($data["options"]) ? (function($input) {
    	/** @var array{pattern?: string, result?: mixed} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1RegexMapOptions::fromArray($val);
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
