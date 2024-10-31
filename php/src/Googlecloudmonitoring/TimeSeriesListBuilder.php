<?php

namespace Grafana\Foundation\Googlecloudmonitoring;

/**
 * Time Series List sub-query properties.
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Googlecloudmonitoring\TimeSeriesList>
 */
class TimeSeriesListBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Googlecloudmonitoring\TimeSeriesList $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Googlecloudmonitoring\TimeSeriesList();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Googlecloudmonitoring\TimeSeriesList
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
     * Reducer applied across a set of time-series values. Defaults to REDUCE_NONE.
     */
    public function crossSeriesReducer(string $crossSeriesReducer): static
    {
        $this->internal->crossSeriesReducer = $crossSeriesReducer;
    
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
     * Alignment function to be used. Defaults to ALIGN_MEAN.
     */
    public function perSeriesAligner(string $perSeriesAligner): static
    {
        $this->internal->perSeriesAligner = $perSeriesAligner;
    
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
    /**
     * Data view, defaults to FULL.
     */
    public function view(string $view): static
    {
        $this->internal->view = $view;
    
        return $this;
    }
    /**
     * Only present if a preprocessor is selected. Reducer applied across a set of time-series values. Defaults to REDUCE_NONE.
     */
    public function secondaryCrossSeriesReducer(string $secondaryCrossSeriesReducer): static
    {
        $this->internal->secondaryCrossSeriesReducer = $secondaryCrossSeriesReducer;
    
        return $this;
    }
    /**
     * Only present if a preprocessor is selected. Alignment period to use when regularizing data. Defaults to cloud-monitoring-auto.
     */
    public function secondaryAlignmentPeriod(string $secondaryAlignmentPeriod): static
    {
        $this->internal->secondaryAlignmentPeriod = $secondaryAlignmentPeriod;
    
        return $this;
    }
    /**
     * Only present if a preprocessor is selected. Alignment function to be used. Defaults to ALIGN_MEAN.
     */
    public function secondaryPerSeriesAligner(string $secondaryPerSeriesAligner): static
    {
        $this->internal->secondaryPerSeriesAligner = $secondaryPerSeriesAligner;
    
        return $this;
    }
    /**
     * Only present if a preprocessor is selected. Array of labels to group data by.
     * @param array<string> $secondaryGroupBys
     */
    public function secondaryGroupBys(array $secondaryGroupBys): static
    {
        $this->internal->secondaryGroupBys = $secondaryGroupBys;
    
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
     * Annotation title.
     */
    public function title(string $title): static
    {
        $this->internal->title = $title;
    
        return $this;
    }
    /**
     * Annotation text.
     */
    public function text(string $text): static
    {
        $this->internal->text = $text;
    
        return $this;
    }

}
