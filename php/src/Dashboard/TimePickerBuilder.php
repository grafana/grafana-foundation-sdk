<?php

namespace Grafana\Foundation\Dashboard;

/**
 * Time picker configuration
 * It defines the default config for the time picker and the refresh picker for the specific dashboard.
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\TimePickerConfig>
 */
class TimePickerBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboard\TimePickerConfig $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboard\TimePickerConfig();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboard\TimePickerConfig
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * Whether timepicker is visible or not.
     */
    public function hidden(bool $hidden): static
    {
        $this->internal->hidden = $hidden;
    
        return $this;
    }
    /**
     * Interval options available in the refresh picker dropdown.
     * @param array<string> $refreshIntervals
     */
    public function refreshIntervals(array $refreshIntervals): static
    {
        $this->internal->refreshIntervals = $refreshIntervals;
    
        return $this;
    }
    /**
     * Quick ranges for time picker.
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\TimeOption>> $quickRanges
     */
    public function quickRanges(array $quickRanges): static
    {
            $quickRangesResources = [];
            foreach ($quickRanges as $r1) {
                    $quickRangesResources[] = $r1->build();
            }
        $this->internal->quickRanges = $quickRangesResources;
    
        return $this;
    }
    /**
     * Override the now time by entering a time delay. Use this option to accommodate known delays in data aggregation to avoid null values.
     */
    public function nowDelay(string $nowDelay): static
    {
        $this->internal->nowDelay = $nowDelay;
    
        return $this;
    }

}
