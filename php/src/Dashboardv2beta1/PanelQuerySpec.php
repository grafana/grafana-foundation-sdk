<?php

namespace Grafana\Foundation\Dashboardv2beta1;

class PanelQuerySpec implements \JsonSerializable
{
    public \Grafana\Foundation\Dashboardv2beta1\DataQueryKind $query;

    public string $refId;

    public bool $hidden;

    /**
     * @param \Grafana\Foundation\Dashboardv2beta1\DataQueryKind|null $query
     * @param string|null $refId
     * @param bool|null $hidden
     */
    public function __construct(?\Grafana\Foundation\Dashboardv2beta1\DataQueryKind $query = null, ?string $refId = null, ?bool $hidden = null)
    {
        $this->query = $query ?: new \Grafana\Foundation\Dashboardv2beta1\DataQueryKind();
        $this->refId = $refId ?: "A";
        $this->hidden = $hidden ?: false;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{query?: mixed, refId?: string, hidden?: bool} $inputData */
        $data = $inputData;
        return new self(
            query: isset($data["query"]) ? (function($input) {
    	/** @var array{kind?: string, group?: string, version?: string, datasource?: mixed, spec?: mixed} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2beta1\DataQueryKind::fromArray($val);
    })($data["query"]) : null,
            refId: $data["refId"] ?? null,
            hidden: $data["hidden"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->query = $this->query;
        $data->refId = $this->refId;
        $data->hidden = $this->hidden;
        return $data;
    }
}
