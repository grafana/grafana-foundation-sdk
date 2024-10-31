<?php

namespace Grafana\Foundation\Googlecloudmonitoring;

/**
 * SLO sub-query properties.
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Googlecloudmonitoring\SLOQuery>
 */
class SLOQueryBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Googlecloudmonitoring\SLOQuery $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Googlecloudmonitoring\SLOQuery();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Googlecloudmonitoring\SLOQuery
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
     * SLO selector.
     */
    public function selectorName(string $selectorName): static
    {
        $this->internal->selectorName = $selectorName;
    
        return $this;
    }
    /**
     * ID for the service the SLO is in.
     */
    public function serviceId(string $serviceId): static
    {
        $this->internal->serviceId = $serviceId;
    
        return $this;
    }
    /**
     * Name for the service the SLO is in.
     */
    public function serviceName(string $serviceName): static
    {
        $this->internal->serviceName = $serviceName;
    
        return $this;
    }
    /**
     * ID for the SLO.
     */
    public function sloId(string $sloId): static
    {
        $this->internal->sloId = $sloId;
    
        return $this;
    }
    /**
     * Name of the SLO.
     */
    public function sloName(string $sloName): static
    {
        $this->internal->sloName = $sloName;
    
        return $this;
    }
    /**
     * SLO goal value.
     */
    public function goal(float $goal): static
    {
        $this->internal->goal = $goal;
    
        return $this;
    }
    /**
     * Specific lookback period for the SLO.
     */
    public function lookbackPeriod(string $lookbackPeriod): static
    {
        $this->internal->lookbackPeriod = $lookbackPeriod;
    
        return $this;
    }

}
