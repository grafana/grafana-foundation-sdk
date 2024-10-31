<?php

namespace Grafana\Foundation\Googlecloudmonitoring;

/**
 * @deprecated Use AnnotationQuery instead. Legacy annotation query properties for migration purposes.
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Googlecloudmonitoring\LegacyCloudMonitoringAnnotationQuery>
 */
class LegacyCloudMonitoringAnnotationQueryBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Googlecloudmonitoring\LegacyCloudMonitoringAnnotationQuery $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Googlecloudmonitoring\LegacyCloudMonitoringAnnotationQuery();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Googlecloudmonitoring\LegacyCloudMonitoringAnnotationQuery
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
    public function metricType(string $metricType): static
    {
        $this->internal->metricType = $metricType;
    
        return $this;
    }
    /**
     * Query refId.
     */
    public function refId(string $refId): static
    {
        $this->internal->refId = $refId;
    
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
