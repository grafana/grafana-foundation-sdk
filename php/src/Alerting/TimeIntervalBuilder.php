<?php

namespace Grafana\Foundation\Alerting;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alerting\TimeInterval>
 */
class TimeIntervalBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Alerting\TimeInterval $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Alerting\TimeInterval();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Alerting\TimeInterval
     */
    public function build()
    {
        return $this->internal;
    }

    public function name(string $name): static
    {
        $this->internal->name = $name;
    
        return $this;
    }
    /**
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alerting\TimeIntervalItem>> $timeIntervals
     */
    public function timeIntervals(array $timeIntervals): static
    {
            $timeIntervalsResources = [];
            foreach ($timeIntervals as $r1) {
                    $timeIntervalsResources[] = $r1->build();
            }
        $this->internal->timeIntervals = $timeIntervalsResources;
    
        return $this;
    }

}
