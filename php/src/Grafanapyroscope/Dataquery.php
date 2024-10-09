<?php

namespace Grafana\Foundation\Grafanapyroscope;

class Dataquery implements \JsonSerializable, \Grafana\Foundation\Cog\Dataquery
{
    /**
     * Specifies the query label selectors.
     */
    public string $labelSelector;

    /**
     * Specifies the query span selectors.
     * @var array<string>|null
     */
    public ?array $spanSelector;

    /**
     * Specifies the type of profile to query.
     */
    public string $profileTypeId;

    /**
     * Allows to group the results.
     * @var array<string>
     */
    public array $groupBy;

    /**
     * Sets the maximum number of nodes in the flamegraph.
     */
    public ?int $maxNodes;

    /**
     * A unique identifier for the query within the list of targets.
     * In server side expressions, the refId is used as a variable name to identify results.
     * By default, the UI will assign A->Z; however setting meaningful names may be useful.
     */
    public string $refId;

    /**
     * true if query is disabled (ie should not be returned to the dashboard)
     * Note this does not always imply that the query should not be executed since
     * the results from a hidden query may be used as the input to other queries (SSE etc)
     */
    public ?bool $hide;

    /**
     * Specify the query flavor
     * TODO make this required and give it a default
     */
    public ?string $queryType;

    /**
     * For mixed data sources the selected datasource is on the query level.
     * For non mixed scenarios this is undefined.
     * TODO find a better way to do this ^ that's friendly to schema
     * TODO this shouldn't be unknown but DataSourceRef | null
     */
    public ?\Grafana\Foundation\Dashboard\DataSourceRef $datasource;

    /**
     * @param string|null $labelSelector
     * @param array<string>|null $spanSelector
     * @param string|null $profileTypeId
     * @param array<string>|null $groupBy
     * @param int|null $maxNodes
     * @param string|null $refId
     * @param bool|null $hide
     * @param string|null $queryType
     * @param \Grafana\Foundation\Dashboard\DataSourceRef|null $datasource
     */
    public function __construct(?string $labelSelector = null, ?array $spanSelector = null, ?string $profileTypeId = null, ?array $groupBy = null, ?int $maxNodes = null, ?string $refId = null, ?bool $hide = null, ?string $queryType = null, ?\Grafana\Foundation\Dashboard\DataSourceRef $datasource = null)
    {
        $this->labelSelector = $labelSelector ?: "{}";
        $this->spanSelector = $spanSelector;
        $this->profileTypeId = $profileTypeId ?: "";
        $this->groupBy = $groupBy ?: [];
        $this->maxNodes = $maxNodes;
        $this->refId = $refId ?: "";
        $this->hide = $hide;
        $this->queryType = $queryType;
        $this->datasource = $datasource;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{labelSelector?: string, spanSelector?: array<string>, profileTypeId?: string, groupBy?: array<string>, maxNodes?: int, refId?: string, hide?: bool, queryType?: string, datasource?: mixed} $inputData */
        $data = $inputData;
        return new self(
            labelSelector: $data["labelSelector"] ?? null,
            spanSelector: $data["spanSelector"] ?? null,
            profileTypeId: $data["profileTypeId"] ?? null,
            groupBy: $data["groupBy"] ?? null,
            maxNodes: $data["maxNodes"] ?? null,
            refId: $data["refId"] ?? null,
            hide: $data["hide"] ?? null,
            queryType: $data["queryType"] ?? null,
            datasource: isset($data["datasource"]) ? (function($input) {
    	/** @var array{type?: string, uid?: string} */
    $val = $input;
    	return \Grafana\Foundation\Dashboard\DataSourceRef::fromArray($val);
    })($data["datasource"]) : null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "labelSelector" => $this->labelSelector,
            "profileTypeId" => $this->profileTypeId,
            "groupBy" => $this->groupBy,
            "refId" => $this->refId,
        ];
        if (isset($this->spanSelector)) {
            $data["spanSelector"] = $this->spanSelector;
        }
        if (isset($this->maxNodes)) {
            $data["maxNodes"] = $this->maxNodes;
        }
        if (isset($this->hide)) {
            $data["hide"] = $this->hide;
        }
        if (isset($this->queryType)) {
            $data["queryType"] = $this->queryType;
        }
        if (isset($this->datasource)) {
            $data["datasource"] = $this->datasource;
        }
        return $data;
    }

    public function dataqueryType(): string
    {
        return "grafanapyroscope";
    }
}
