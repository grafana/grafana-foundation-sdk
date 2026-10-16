<?php

namespace Grafana\Foundation\Dashboardv2beta1;

class NotebookLayoutKind implements \JsonSerializable
{
    public string $kind;

    public \Grafana\Foundation\Dashboardv2beta1\NotebookLayoutSpec $spec;

    /**
     * @param \Grafana\Foundation\Dashboardv2beta1\NotebookLayoutSpec|null $spec
     */
    public function __construct(?\Grafana\Foundation\Dashboardv2beta1\NotebookLayoutSpec $spec = null)
    {
        $this->kind = "NotebookLayout";
    
        $this->spec = $spec ?: new \Grafana\Foundation\Dashboardv2beta1\NotebookLayoutSpec();
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
    	/** @var array{cells?: array<mixed>} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2beta1\NotebookLayoutSpec::fromArray($val);
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
