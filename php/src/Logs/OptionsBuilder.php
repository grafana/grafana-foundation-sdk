<?php

namespace Grafana\Foundation\Logs;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Logs\Options>
 */
class OptionsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Logs\Options $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Logs\Options();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Logs\Options
     */
    public function build()
    {
        return $this->internal;
    }

    public function showLabels(bool $showLabels): static
    {
        $this->internal->showLabels = $showLabels;
    
        return $this;
    }

    public function showCommonLabels(bool $showCommonLabels): static
    {
        $this->internal->showCommonLabels = $showCommonLabels;
    
        return $this;
    }

    public function showTime(bool $showTime): static
    {
        $this->internal->showTime = $showTime;
    
        return $this;
    }

    public function showLogContextToggle(bool $showLogContextToggle): static
    {
        $this->internal->showLogContextToggle = $showLogContextToggle;
    
        return $this;
    }

    public function wrapLogMessage(bool $wrapLogMessage): static
    {
        $this->internal->wrapLogMessage = $wrapLogMessage;
    
        return $this;
    }

    public function prettifyLogMessage(bool $prettifyLogMessage): static
    {
        $this->internal->prettifyLogMessage = $prettifyLogMessage;
    
        return $this;
    }

    public function enableLogDetails(bool $enableLogDetails): static
    {
        $this->internal->enableLogDetails = $enableLogDetails;
    
        return $this;
    }

    public function sortOrder(\Grafana\Foundation\Common\LogsSortOrder $sortOrder): static
    {
        $this->internal->sortOrder = $sortOrder;
    
        return $this;
    }

    public function dedupStrategy(\Grafana\Foundation\Common\LogsDedupStrategy $dedupStrategy): static
    {
        $this->internal->dedupStrategy = $dedupStrategy;
    
        return $this;
    }

}
