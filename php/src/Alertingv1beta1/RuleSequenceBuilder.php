<?php

namespace Grafana\Foundation\Alertingv1beta1;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alertingv1beta1\RuleSequenceSpec>
 */
class RuleSequenceBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Alertingv1beta1\RuleSequenceSpec $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Alertingv1beta1\RuleSequenceSpec();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Alertingv1beta1\RuleSequenceSpec
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alertingv1beta1\IntervalTrigger> $trigger
     */
    public function trigger(\Grafana\Foundation\Cog\Builder $trigger): static
    {
        $triggerResource = $trigger->build();
        $this->internal->trigger = $triggerResource;
    
        return $this;
    }

    /**
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alertingv1beta1\RuleRef>> $recordingRules
     */
    public function recordingRules(array $recordingRules): static
    {
            $recordingRulesResources = [];
            foreach ($recordingRules as $r1) {
                    $recordingRulesResources[] = $r1->build();
            }
        $this->internal->recordingRules = $recordingRulesResources;
    
        return $this;
    }

    /**
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alertingv1beta1\RuleRef>> $alertRules
     */
    public function alertRules(array $alertRules): static
    {
            $alertRulesResources = [];
            foreach ($alertRules as $r1) {
                    $alertRulesResources[] = $r1->build();
            }
        $this->internal->alertRules = $alertRulesResources;
    
        return $this;
    }

}
