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
     * Builds the object.
     * @return \Grafana\Foundation\Alerting\RecordRule
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * Which expression node should be used as the input for the recorded metric.
     */
    public function from(string $from): static
    {
        $this->internal->from = $from;
    
        return $this;
    }
    /**
     * Name of the recorded metric.
     */
    public function metric(string $metric): static
    {
        $this->internal->metric = $metric;
    
        return $this;
    }
    /**
     * Which data source should be used to write the output of the recording rule, specified by UID.
     */
    public function targetDatasourceUid(string $targetDatasourceUid): static
    {
        $this->internal->targetDatasourceUid = $targetDatasourceUid;
    
        return $this;
    }

}
