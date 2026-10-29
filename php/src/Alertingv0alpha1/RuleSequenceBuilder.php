<?php

namespace Grafana\Foundation\Alertingv0alpha1;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alertingv0alpha1\RuleSequenceSpec>
 */
class RuleSequenceBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Alertingv0alpha1\RuleSequenceSpec $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Alertingv0alpha1\RuleSequenceSpec();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Alertingv0alpha1\RuleSequenceSpec
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alertingv0alpha1\IntervalTrigger> $trigger
     */
    public function trigger(\Grafana\Foundation\Cog\Builder $trigger): static
    {
        $triggerResource = $trigger->build();
        $this->internal->trigger = $triggerResource;
    
        return $this;
    }

    /**
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alertingv0alpha1\RuleRef>> $recordingRules
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
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alertingv0alpha1\RuleRef>> $alertingRules
     */
    public function alertingRules(array $alertingRules): static
    {
            $alertingRulesResources = [];
            foreach ($alertingRules as $r1) {
                    $alertingRulesResources[] = $r1->build();
            }
        $this->internal->alertingRules = $alertingRulesResources;
    
        return $this;
    }

}
