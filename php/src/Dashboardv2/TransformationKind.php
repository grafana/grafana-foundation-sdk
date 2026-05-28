<?php

namespace Grafana\Foundation\Dashboardv2;

class TransformationKind implements \JsonSerializable
{
    public string $kind;

    /**
     * The group is the transformation ID
     */
    public string $group;

    public \Grafana\Foundation\Dashboardv2\TransformationSpec $spec;

    /**
     * @param string|null $group
     * @param \Grafana\Foundation\Dashboardv2\TransformationSpec|null $spec
     */
    public function __construct(?string $group = null, ?\Grafana\Foundation\Dashboardv2\TransformationSpec $spec = null)
    {
        $this->kind = "Transformation";
    
        $this->group = $group ?: "";
        $this->spec = $spec ?: new \Grafana\Foundation\Dashboardv2\TransformationSpec();
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{kind?: string, group?: string, spec?: mixed} $inputData */
        $data = $inputData;
        return new self(
            group: $data["group"] ?? null,
            spec: isset($data["spec"]) ? (function($input) {
    	/** @var array{disabled?: bool, filter?: mixed, topic?: string, options?: mixed} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2\TransformationSpec::fromArray($val);
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
        $data->group = $this->group;
        $data->spec = $this->spec;
        return $data;
    }
}
