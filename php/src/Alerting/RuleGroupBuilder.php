<?php

namespace Grafana\Foundation\Alerting;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alerting\RuleGroup>
 */
class RuleGroupBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Alerting\RuleGroup $internal;

    public function __construct(string $title)
    {
    	$this->internal = new \Grafana\Foundation\Alerting\RuleGroup();
    $this->internal->title = $title;
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Alerting\RuleGroup
     */
    public function build()
    {
        return $this->internal;
    }

    public function folderUid(string $folderUid): static
    {
        $this->internal->folderUid = $folderUid;
    
        return $this;
    }

    /**
     * The interval, in seconds, at which all rules in the group are evaluated.
     * If a group contains many rules, the rules are evaluated sequentially.
     */
    public function interval(int $interval): static
    {
        $this->internal->interval = $interval;
    
        return $this;
    }

    /**
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alerting\Rule>> $rules
     */
    public function rules(array $rules): static
    {
            $rulesResources = [];
            foreach ($rules as $r1) {
                    $rulesResources[] = $r1->build();
            }
        $this->internal->rules = $rulesResources;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alerting\Rule> $rule
     */
    public function withRule(\Grafana\Foundation\Cog\Builder $rule): static
    {    
        if ($this->internal->rules === null) {
            $this->internal->rules = [];
        }
        
        $ruleResource = $rule->build();
        $this->internal->rules[] = $ruleResource;
    
        return $this;
    }

    public function title(string $title): static
    {
        $this->internal->title = $title;
    
        return $this;
    }

}
