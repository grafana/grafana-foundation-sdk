<?php

namespace Grafana\Foundation\Dashboardv2beta1;

class LibraryPanelKind implements \JsonSerializable
{
    public string $kind;

    public \Grafana\Foundation\Dashboardv2beta1\LibraryPanelKindSpec $spec;

    /**
     * @param \Grafana\Foundation\Dashboardv2beta1\LibraryPanelKindSpec|null $spec
     */
    public function __construct(?\Grafana\Foundation\Dashboardv2beta1\LibraryPanelKindSpec $spec = null)
    {
        $this->kind = "LibraryPanel";
    
        $this->spec = $spec ?: new \Grafana\Foundation\Dashboardv2beta1\LibraryPanelKindSpec();
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
    	/** @var array{id?: float, title?: string, libraryPanel?: mixed} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2beta1\LibraryPanelKindSpec::fromArray($val);
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
