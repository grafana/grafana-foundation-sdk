<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\QueryOptionsSpec>
 */
class QueryOptionsSpecBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2beta1\QueryOptionsSpec $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2beta1\QueryOptionsSpec();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2beta1\QueryOptionsSpec
     */
    public function build()
    {
        return $this->internal;
    }

    public function timeFrom(string $timeFrom): static
    {
        $this->internal->timeFrom = $timeFrom;
    
        return $this;
    }

    public function maxDataPoints(int $maxDataPoints): static
    {
        $this->internal->maxDataPoints = $maxDataPoints;
    
        return $this;
    }

    public function timeShift(string $timeShift): static
    {
        $this->internal->timeShift = $timeShift;
    
        return $this;
    }

    public function queryCachingTTL(int $queryCachingTTL): static
    {
        $this->internal->queryCachingTTL = $queryCachingTTL;
    
        return $this;
    }

    public function interval(string $interval): static
    {
        $this->internal->interval = $interval;
    
        return $this;
    }

    public function cacheTimeout(string $cacheTimeout): static
    {
        $this->internal->cacheTimeout = $cacheTimeout;
    
        return $this;
    }

    public function hideTimeOverride(bool $hideTimeOverride): static
    {
        $this->internal->hideTimeOverride = $hideTimeOverride;
    
        return $this;
    }

    public function timeCompare(string $timeCompare): static
    {
        $this->internal->timeCompare = $timeCompare;
    
        return $this;
    }

}
