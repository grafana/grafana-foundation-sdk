<?php

namespace Grafana\Foundation\Logsnew;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Logsnew\Options>
 */
class OptionsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Logsnew\Options $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Logsnew\Options();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Logsnew\Options
     */
    public function build()
    {
        return $this->internal;
    }

    public function showTime(bool $showTime): static
    {
        $this->internal->showTime = $showTime;
    
        return $this;
    }

    public function wrapLogMessage(bool $wrapLogMessage): static
    {
        $this->internal->wrapLogMessage = $wrapLogMessage;
    
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

    public function enableInfiniteScrolling(bool $enableInfiniteScrolling): static
    {
        $this->internal->enableInfiniteScrolling = $enableInfiniteScrolling;
    
        return $this;
    }

    /**
     * @param mixed $onNewLogsReceived
     */
    public function onNewLogsReceived( $onNewLogsReceived): static
    {
        $this->internal->onNewLogsReceived = $onNewLogsReceived;
    
        return $this;
    }

}
