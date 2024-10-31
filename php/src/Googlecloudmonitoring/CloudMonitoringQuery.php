<?php

namespace Grafana\Foundation\Googlecloudmonitoring;

class CloudMonitoringQuery implements \JsonSerializable, \Grafana\Foundation\Cog\Dataquery
{
    /**
     * A unique identifier for the query within the list of targets.
     * In server side expressions, the refId is used as a variable name to identify results.
     * By default, the UI will assign A->Z; however setting meaningful names may be useful.
     */
    public string $refId;

    /**
     * If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
     */
    public ?bool $hide;

    /**
     * Specify the query flavor
     * TODO make this required and give it a default
     */
    public ?string $queryType;

    /**
     * Aliases can be set to modify the legend labels. e.g. {{metric.label.xxx}}. See docs for more detail.
     */
    public ?string $aliasBy;

    /**
     * GCM query type.
     * queryType: #QueryType
     * Time Series List sub-query properties.
     */
    public ?\Grafana\Foundation\Googlecloudmonitoring\TimeSeriesList $timeSeriesList;

    /**
     * Time Series sub-query properties.
     */
    public ?\Grafana\Foundation\Googlecloudmonitoring\TimeSeriesQuery $timeSeriesQuery;

    /**
     * SLO sub-query properties.
     */
    public ?\Grafana\Foundation\Googlecloudmonitoring\SLOQuery $sloQuery;

    /**
     * PromQL sub-query properties.
     */
    public ?\Grafana\Foundation\Googlecloudmonitoring\PromQLQuery $promQLQuery;

    /**
     * For mixed data sources the selected datasource is on the query level.
     * For non mixed scenarios this is undefined.
     * TODO find a better way to do this ^ that's friendly to schema
     * TODO this shouldn't be unknown but DataSourceRef | null
     */
    public ?\Grafana\Foundation\Dashboard\DataSourceRef $datasource;

    /**
     * Time interval in milliseconds.
     */
    public ?float $intervalMs;

    /**
     * @param string|null $refId
     * @param bool|null $hide
     * @param string|null $queryType
     * @param string|null $aliasBy
     * @param \Grafana\Foundation\Googlecloudmonitoring\TimeSeriesList|null $timeSeriesList
     * @param \Grafana\Foundation\Googlecloudmonitoring\TimeSeriesQuery|null $timeSeriesQuery
     * @param \Grafana\Foundation\Googlecloudmonitoring\SLOQuery|null $sloQuery
     * @param \Grafana\Foundation\Googlecloudmonitoring\PromQLQuery|null $promQLQuery
     * @param \Grafana\Foundation\Dashboard\DataSourceRef|null $datasource
     * @param float|null $intervalMs
     */
    public function __construct(?string $refId = null, ?bool $hide = null, ?string $queryType = null, ?string $aliasBy = null, ?\Grafana\Foundation\Googlecloudmonitoring\TimeSeriesList $timeSeriesList = null, ?\Grafana\Foundation\Googlecloudmonitoring\TimeSeriesQuery $timeSeriesQuery = null, ?\Grafana\Foundation\Googlecloudmonitoring\SLOQuery $sloQuery = null, ?\Grafana\Foundation\Googlecloudmonitoring\PromQLQuery $promQLQuery = null, ?\Grafana\Foundation\Dashboard\DataSourceRef $datasource = null, ?float $intervalMs = null)
    {
        $this->refId = $refId ?: "";
        $this->hide = $hide;
        $this->queryType = $queryType;
        $this->aliasBy = $aliasBy;
        $this->timeSeriesList = $timeSeriesList;
        $this->timeSeriesQuery = $timeSeriesQuery;
        $this->sloQuery = $sloQuery;
        $this->promQLQuery = $promQLQuery;
        $this->datasource = $datasource;
        $this->intervalMs = $intervalMs;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{refId?: string, hide?: bool, queryType?: string, aliasBy?: string, timeSeriesList?: mixed, timeSeriesQuery?: mixed, sloQuery?: mixed, promQLQuery?: mixed, datasource?: mixed, intervalMs?: float} $inputData */
        $data = $inputData;
        return new self(
            refId: $data["refId"] ?? null,
            hide: $data["hide"] ?? null,
            queryType: $data["queryType"] ?? null,
            aliasBy: $data["aliasBy"] ?? null,
            timeSeriesList: isset($data["timeSeriesList"]) ? (function($input) {
    	/** @var array{projectName?: string, crossSeriesReducer?: string, alignmentPeriod?: string, perSeriesAligner?: string, groupBys?: array<string>, filters?: array<string>, view?: string, title?: string, text?: string, secondaryCrossSeriesReducer?: string, secondaryAlignmentPeriod?: string, secondaryPerSeriesAligner?: string, secondaryGroupBys?: array<string>, preprocessor?: string} */
    $val = $input;
    	return \Grafana\Foundation\Googlecloudmonitoring\TimeSeriesList::fromArray($val);
    })($data["timeSeriesList"]) : null,
            timeSeriesQuery: isset($data["timeSeriesQuery"]) ? (function($input) {
    	/** @var array{projectName?: string, query?: string, graphPeriod?: string} */
    $val = $input;
    	return \Grafana\Foundation\Googlecloudmonitoring\TimeSeriesQuery::fromArray($val);
    })($data["timeSeriesQuery"]) : null,
            sloQuery: isset($data["sloQuery"]) ? (function($input) {
    	/** @var array{projectName?: string, perSeriesAligner?: string, alignmentPeriod?: string, selectorName?: string, serviceId?: string, serviceName?: string, sloId?: string, sloName?: string, goal?: float, lookbackPeriod?: string} */
    $val = $input;
    	return \Grafana\Foundation\Googlecloudmonitoring\SLOQuery::fromArray($val);
    })($data["sloQuery"]) : null,
            promQLQuery: isset($data["promQLQuery"]) ? (function($input) {
    	/** @var array{projectName?: string, expr?: string, step?: string} */
    $val = $input;
    	return \Grafana\Foundation\Googlecloudmonitoring\PromQLQuery::fromArray($val);
    })($data["promQLQuery"]) : null,
            datasource: isset($data["datasource"]) ? (function($input) {
    	/** @var array{type?: string, uid?: string} */
    $val = $input;
    	return \Grafana\Foundation\Dashboard\DataSourceRef::fromArray($val);
    })($data["datasource"]) : null,
            intervalMs: $data["intervalMs"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "refId" => $this->refId,
        ];
        if (isset($this->hide)) {
            $data["hide"] = $this->hide;
        }
        if (isset($this->queryType)) {
            $data["queryType"] = $this->queryType;
        }
        if (isset($this->aliasBy)) {
            $data["aliasBy"] = $this->aliasBy;
        }
        if (isset($this->timeSeriesList)) {
            $data["timeSeriesList"] = $this->timeSeriesList;
        }
        if (isset($this->timeSeriesQuery)) {
            $data["timeSeriesQuery"] = $this->timeSeriesQuery;
        }
        if (isset($this->sloQuery)) {
            $data["sloQuery"] = $this->sloQuery;
        }
        if (isset($this->promQLQuery)) {
            $data["promQLQuery"] = $this->promQLQuery;
        }
        if (isset($this->datasource)) {
            $data["datasource"] = $this->datasource;
        }
        if (isset($this->intervalMs)) {
            $data["intervalMs"] = $this->intervalMs;
        }
        return $data;
    }

    public function dataqueryType(): string
    {
        return "cloud-monitoring";
    }
}
