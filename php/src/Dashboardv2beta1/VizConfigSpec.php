<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * --- Kinds ---
 */
class VizConfigSpec implements \JsonSerializable
{
    /**
     * @var mixed|null
     */
    public $options;

    public \Grafana\Foundation\Dashboardv2beta1\FieldConfigSource $fieldConfig;

    /**
     * @param mixed|null $options
     * @param \Grafana\Foundation\Dashboardv2beta1\FieldConfigSource|null $fieldConfig
     */
    public function __construct( $options = null, ?\Grafana\Foundation\Dashboardv2beta1\FieldConfigSource $fieldConfig = null)
    {
        $this->options = $options;
        $this->fieldConfig = $fieldConfig ?: new \Grafana\Foundation\Dashboardv2beta1\FieldConfigSource();
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{options?: mixed, fieldConfig?: mixed} $inputData */
        $data = $inputData;
        return new self(
            options: $data["options"] ?? null,
            fieldConfig: isset($data["fieldConfig"]) ? (function($input) {
    	/** @var array{defaults?: mixed, overrides?: array<mixed>} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2beta1\FieldConfigSource::fromArray($val);
    })($data["fieldConfig"]) : null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->fieldConfig = $this->fieldConfig;
        if (isset($this->options)) {
            $data->options = $this->options;
        }
        return $data;
    }
}
