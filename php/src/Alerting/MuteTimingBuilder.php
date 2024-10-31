<?php

namespace Grafana\Foundation\Alerting;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alerting\MuteTiming>
 */
class MuteTimingBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Alerting\MuteTiming $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Alerting\MuteTiming();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Alerting\MuteTiming
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
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alerting\TimeInterval>> $timeIntervals
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
