<?php

namespace Grafana\Foundation\Dashboardv2beta1;

class MarkdownCellContentKind implements \JsonSerializable
{
    public string $kind;

    public \Grafana\Foundation\Dashboardv2beta1\MarkdownCellContentSpec $spec;

    /**
     * @param \Grafana\Foundation\Dashboardv2beta1\MarkdownCellContentSpec|null $spec
     */
    public function __construct(?\Grafana\Foundation\Dashboardv2beta1\MarkdownCellContentSpec $spec = null)
    {
        $this->kind = "Markdown";
    
        $this->spec = $spec ?: new \Grafana\Foundation\Dashboardv2beta1\MarkdownCellContentSpec();
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
    	/** @var array{text?: string} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2beta1\MarkdownCellContentSpec::fromArray($val);
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
