<?php

namespace Grafana\Foundation\Dashboardv2beta1;

class TabsLayoutSpec implements \JsonSerializable
{
    /**
     * @var array<\Grafana\Foundation\Dashboardv2beta1\TabsLayoutTabKind>
     */
    public array $tabs;

    /**
     * @param array<\Grafana\Foundation\Dashboardv2beta1\TabsLayoutTabKind>|null $tabs
     */
    public function __construct(?array $tabs = null)
    {
        $this->tabs = $tabs ?: [];
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{tabs?: array<mixed>} $inputData */
        $data = $inputData;
        return new self(
            tabs: array_filter(array_map((function($input) {
    	/** @var array{kind?: string, spec?: mixed} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2beta1\TabsLayoutTabKind::fromArray($val);
    }), $data["tabs"] ?? [])),
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->tabs = $this->tabs;
        return $data;
    }
}
