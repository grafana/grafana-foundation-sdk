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

    public function enableInfiniteScrolling(bool $enableInfiniteScrolling): static
    {
        $this->internal->enableInfiniteScrolling = $enableInfiniteScrolling;
    
        return $this;
    }

    /**
     * TODO: figure out how to define callbacks
     * @param mixed $onClickFilterLabel
     */
    public function onClickFilterLabel( $onClickFilterLabel): static
    {
        $this->internal->onClickFilterLabel = $onClickFilterLabel;
    
        return $this;
    }

    /**
     * @param mixed $onClickFilterOutLabel
     */
    public function onClickFilterOutLabel( $onClickFilterOutLabel): static
    {
        $this->internal->onClickFilterOutLabel = $onClickFilterOutLabel;
    
        return $this;
    }

    /**
     * @param mixed $isFilterLabelActive
     */
    public function isFilterLabelActive( $isFilterLabelActive): static
    {
        $this->internal->isFilterLabelActive = $isFilterLabelActive;
    
        return $this;
    }

    /**
     * @param mixed $onClickFilterString
     */
    public function onClickFilterString( $onClickFilterString): static
    {
        $this->internal->onClickFilterString = $onClickFilterString;
    
        return $this;
    }

    /**
     * @param mixed $onClickFilterOutString
     */
    public function onClickFilterOutString( $onClickFilterOutString): static
    {
        $this->internal->onClickFilterOutString = $onClickFilterOutString;
    
        return $this;
    }

    /**
     * @param mixed $onClickShowField
     */
    public function onClickShowField( $onClickShowField): static
    {
        $this->internal->onClickShowField = $onClickShowField;
    
        return $this;
    }

    /**
     * @param mixed $onClickHideField
     */
    public function onClickHideField( $onClickHideField): static
    {
        $this->internal->onClickHideField = $onClickHideField;
    
        return $this;
    }

    /**
     * @param mixed $logRowMenuIconsBefore
     */
    public function logRowMenuIconsBefore( $logRowMenuIconsBefore): static
    {
        $this->internal->logRowMenuIconsBefore = $logRowMenuIconsBefore;
    
        return $this;
    }

    /**
     * @param mixed $logRowMenuIconsAfter
     */
    public function logRowMenuIconsAfter( $logRowMenuIconsAfter): static
    {
        $this->internal->logRowMenuIconsAfter = $logRowMenuIconsAfter;
    
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

    /**
     * @param array<string> $displayedFields
     */
    public function displayedFields(array $displayedFields): static
    {
        $this->internal->displayedFields = $displayedFields;
    
        return $this;
    }

}
