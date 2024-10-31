<?php

namespace Grafana\Foundation\Dashboard;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\TimePicker>
 */
class TimePickerBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboard\TimePicker $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboard\TimePicker();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboard\TimePicker
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
     * Whether timepicker is collapsed or not. Has no effect on provisioned dashboard.
     */
    public function collapse(bool $collapse): static
    {
        $this->internal->collapse = $collapse;
    
        return $this;
    }
    /**
     * Whether timepicker is enabled or not. Has no effect on provisioned dashboard.
     */
    public function enable(bool $enable): static
    {
        $this->internal->enable = $enable;
    
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
