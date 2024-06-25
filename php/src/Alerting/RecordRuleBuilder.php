<?php

namespace Grafana\Foundation\Alerting;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alerting\RecordRule>
 */
class RecordRuleBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Alerting\RecordRule $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Alerting\RecordRule();
    }

    /**
     * @return \Grafana\Foundation\Alerting\RecordRule
     */
    public function build()
    {
        return $this->internal;
    }

    public function from(string $from): static
    {
        $this->internal->from = $from;
    
        return $this;
    }
    public function metric(string $metric): static
    {
        $this->internal->metric = $metric;
    
        return $this;
    }

}
