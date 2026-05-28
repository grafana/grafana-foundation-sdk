<?php

namespace Grafana\Foundation\Dashboardv2;

class TabsLayoutTabKind implements \JsonSerializable
{
    public string $kind;

    public \Grafana\Foundation\Dashboardv2\TabsLayoutTabSpec $spec;

    /**
     * @param \Grafana\Foundation\Dashboardv2\TabsLayoutTabSpec|null $spec
     */
    public function __construct(?\Grafana\Foundation\Dashboardv2\TabsLayoutTabSpec $spec = null)
    {
        $this->kind = "TabsLayoutTab";
    
        $this->spec = $spec ?: new \Grafana\Foundation\Dashboardv2\TabsLayoutTabSpec();
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
    	/** @var array{title?: string, layout?: mixed|mixed|mixed|mixed, conditionalRendering?: mixed, repeat?: mixed, variables?: array<mixed|mixed|mixed|mixed|mixed|mixed|mixed|mixed|mixed>} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2\TabsLayoutTabSpec::fromArray($val);
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
