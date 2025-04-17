<?php

namespace Grafana\Foundation\Dashboard;

class DashboardDashboardTemplating implements \JsonSerializable
{
    /**
     * List of configured template variables with their saved values along with some other metadata
     * @var array<\Grafana\Foundation\Dashboard\VariableModel>|null
     */
    public ?array $list;

    /**
     * @param array<\Grafana\Foundation\Dashboard\VariableModel>|null $list
     */
    public function __construct(?array $list = null)
    {
        $this->list = $list;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{list?: array<mixed>} $inputData */
        $data = $inputData;
        return new self(
            list: array_filter(array_map((function($input) {
    	/** @var array{id?: string, type?: string, name?: string, label?: string, hide?: int, skipUrlSync?: bool, description?: string, query?: string|array<string, mixed>, datasource?: mixed, allFormat?: string, current?: mixed, multi?: bool, options?: array<mixed>, refresh?: int, sort?: int, includeAll?: bool, allValue?: string, regex?: string, auto?: bool, auto_min?: string, auto_count?: int} */
    $val = $input;
    	return \Grafana\Foundation\Dashboard\VariableModel::fromArray($val);
    }), $data["list"] ?? [])),
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        if (isset($this->list)) {
            $data->list = $this->list;
        }
        return $data;
    }
}
