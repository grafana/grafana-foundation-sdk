<?php

namespace Grafana\Foundation\Googlecloudmonitoring;

/**
 * @deprecated This type is for migration purposes only. Replaced by TimeSeriesList Metric sub-query properties.
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Googlecloudmonitoring\MetricQuery>
 */
class MetricQueryBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Googlecloudmonitoring\MetricQuery $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Googlecloudmonitoring\MetricQuery();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Googlecloudmonitoring\MetricQuery
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * GCP project to execute the query against.
     */
    public function projectName(string $projectName): static
    {
        $this->internal->projectName = $projectName;
    
        return $this;
    }
    /**
     * Alignment function to be used. Defaults to ALIGN_MEAN.
     */
    public function perSeriesAligner(string $perSeriesAligner): static
    {
        $this->internal->perSeriesAligner = $perSeriesAligner;
    
        return $this;
    }
    /**
     * Alignment period to use when regularizing data. Defaults to cloud-monitoring-auto.
     */
    public function alignmentPeriod(string $alignmentPeriod): static
    {
        $this->internal->alignmentPeriod = $alignmentPeriod;
    
        return $this;
    }
    /**
     * Aliases can be set to modify the legend labels. e.g. {{metric.label.xxx}}. See docs for more detail.
     */
    public function aliasBy(string $aliasBy): static
    {
        $this->internal->aliasBy = $aliasBy;
    
        return $this;
    }
    public function editorMode(string $editorMode): static
    {
        $this->internal->editorMode = $editorMode;
    
        return $this;
    }
    public function metricType(string $metricType): static
    {
        $this->internal->metricType = $metricType;
    
        return $this;
    }
    /**
     * Reducer applied across a set of time-series values. Defaults to REDUCE_NONE.
     */
    public function crossSeriesReducer(string $crossSeriesReducer): static
    {
        $this->internal->crossSeriesReducer = $crossSeriesReducer;
    
        return $this;
    }
    /**
     * Array of labels to group data by.
     * @param array<string> $groupBys
     */
    public function groupBys(array $groupBys): static
    {
        $this->internal->groupBys = $groupBys;
    
        return $this;
    }
    /**
     * Array of filters to query data by. Labels that can be filtered on are defined by the metric.
     * @param array<string> $filters
     */
    public function filters(array $filters): static
    {
        $this->internal->filters = $filters;
    
        return $this;
    }
    public function metricKind(\Grafana\Foundation\Googlecloudmonitoring\MetricKind $metricKind): static
    {
        $this->internal->metricKind = $metricKind;
    
        return $this;
    }
    public function valueType(string $valueType): static
    {
        $this->internal->valueType = $valueType;
    
        return $this;
    }
    public function view(string $view): static
    {
        $this->internal->view = $view;
    
        return $this;
    }
    /**
     * MQL query to be executed.
     */
    public function query(string $query): static
    {
        $this->internal->query = $query;
    
        return $this;
    }
    /**
     * Preprocessor is not part of the API, but is used to store the preprocessor and not affect the UI for the rest of parameters
     */
    public function preprocessor(\Grafana\Foundation\Googlecloudmonitoring\PreprocessorType $preprocessor): static
    {
        $this->internal->preprocessor = $preprocessor;
    
        return $this;
    }
    /**
     * To disable the graphPeriod, it should explictly be set to 'disabled'.
     */
    public function graphPeriod(string $graphPeriod): static
    {
        $this->internal->graphPeriod = $graphPeriod;
    
        return $this;
    }

}
