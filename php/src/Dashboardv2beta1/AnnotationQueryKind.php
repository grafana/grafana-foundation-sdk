<?php

namespace Grafana\Foundation\Dashboardv2beta1;

class AnnotationQueryKind implements \JsonSerializable
{
    public string $kind;

    public \Grafana\Foundation\Dashboardv2beta1\AnnotationQuerySpec $spec;

    /**
     * @param \Grafana\Foundation\Dashboardv2beta1\AnnotationQuerySpec|null $spec
     */
    public function __construct(?\Grafana\Foundation\Dashboardv2beta1\AnnotationQuerySpec $spec = null)
    {
        $this->kind = "AnnotationQuery";
    
        $this->spec = $spec ?: new \Grafana\Foundation\Dashboardv2beta1\AnnotationQuerySpec();
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{kind?: string, spec?: mixed} $inputData */
        $data = $inputData;
        return new self(
            spec: isset($data["spec"]) ? (function($input) {
    	/** @var array{query?: mixed, enable?: bool, hide?: bool, iconColor?: string, name?: string, builtIn?: bool, filter?: mixed, placement?: "inControlsMenu", legacyOptions?: array<string, mixed>} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2beta1\AnnotationQuerySpec::fromArray($val);
    })($data["spec"]) : null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->kind = $this->kind;
        $data->spec = $this->spec;
        return $data;
    }
}
