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
     * Selectable options available in the time picker dropdown. Has no effect on provisioned dashboard.
     * @param array<string> $timeOptions
     */
    public function timeOptions(array $timeOptions): static
    {
        $this->internal->timeOptions = $timeOptions;
    
        return $this;
    }

}
