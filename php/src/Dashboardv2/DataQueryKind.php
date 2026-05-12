<?php

namespace Grafana\Foundation\Dashboardv2;

class DataQueryKind implements \JsonSerializable
{
    public string $kind;

    public string $group;

    public string $version;

    /**
     * @var array<string, string>|null
     */
    public ?array $labels;

    /**
     * New type for datasource reference
     * Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.
     */
    public ?\Grafana\Foundation\Dashboardv2\Dashboardv2DataQueryKindDatasource $datasource;

    /**
     * @var mixed|null
     */
    public $spec;

    /**
     * @param string|null $group
     * @param string|null $version
     * @param array<string, string>|null $labels
     * @param \Grafana\Foundation\Dashboardv2\Dashboardv2DataQueryKindDatasource|null $datasource
     * @param mixed|null $spec
     */
    public function __construct(?string $group = null, ?string $version = null, ?array $labels = null, ?\Grafana\Foundation\Dashboardv2\Dashboardv2DataQueryKindDatasource $datasource = null,  $spec = null)
    {
        $this->kind = "DataQuery";
    
        $this->group = $group ?: "";
        $this->version = $version ?: "v0";
        $this->labels = $labels;
        $this->datasource = $datasource;
        $this->spec = $spec;
    }

    
    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{kind?: string, group?: string, version?: string, labels?: array<string, string>, datasource?: mixed, spec?: mixed} $inputData */
        $data = $inputData;
        return new self( 
            
             
            group: $data["group"] ?? null, 
            version: $data["version"] ?? null, 
            labels: $data["labels"] ?? null, 
            datasource: isset($data["datasource"]) ? (function($input) {
    	/** @var array{name?: string} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2\Dashboardv2DataQueryKindDatasource::fromArray($val);
    })($data["datasource"]) : null, 
            spec: isset($data['spec']) ? (function($input) {
        /** @var array<string, mixed> $spec */
        $spec = $input['spec'];
        return \Grafana\Foundation\Cog\Runtime::get()->dataqueryFromArray($spec, $input['group'] ?? '');
    })($data) : null,
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
        $data->version = $this->version;
        if (isset($this->labels)) {
            $data->labels = $this->labels;
        }
        if (isset($this->datasource)) {
            $data->datasource = $this->datasource;
        }
        if (isset($this->spec)) {
            $data->spec = $this->spec;
        }
        return $data;
    }
}
