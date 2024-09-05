<?php

namespace Grafana\Foundation\Alerting;

/**
 * A Route is a node that contains definitions of how to handle alerts. This is modified
 * from the upstream alertmanager in that it adds the ObjectMatchers property.
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alerting\NotificationPolicy>
 */
class NotificationPolicyBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Alerting\NotificationPolicy $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Alerting\NotificationPolicy();
    }

    /**
     * @return \Grafana\Foundation\Alerting\NotificationPolicy
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * A Route is a node that contains definitions of how to handle alerts. This is modified
     * from the upstream alertmanager in that it adds the ObjectMatchers property.
     * @param array<string> $activeTimeIntervals
     */
    public function activeTimeIntervals(array $activeTimeIntervals): static
    {
        $this->internal->activeTimeIntervals = $activeTimeIntervals;
    
        return $this;
    }
    /**
     * A Route is a node that contains definitions of how to handle alerts. This is modified
     * from the upstream alertmanager in that it adds the ObjectMatchers property.
     */
    public function continue(bool $continue): static
    {
        $this->internal->continue = $continue;
    
        return $this;
    }
    /**
     * A Route is a node that contains definitions of how to handle alerts. This is modified
     * from the upstream alertmanager in that it adds the ObjectMatchers property.
     * @param array<string> $groupBy
     */
    public function groupBy(array $groupBy): static
    {
        $this->internal->groupBy = $groupBy;
    
        return $this;
    }
    /**
     * A Route is a node that contains definitions of how to handle alerts. This is modified
     * from the upstream alertmanager in that it adds the ObjectMatchers property.
     */
    public function groupInterval(string $groupInterval): static
    {
        $this->internal->groupInterval = $groupInterval;
    
        return $this;
    }
    /**
     * A Route is a node that contains definitions of how to handle alerts. This is modified
     * from the upstream alertmanager in that it adds the ObjectMatchers property.
     */
    public function groupWait(string $groupWait): static
    {
        $this->internal->groupWait = $groupWait;
    
        return $this;
    }
    /**
     * A Route is a node that contains definitions of how to handle alerts. This is modified
     * from the upstream alertmanager in that it adds the ObjectMatchers property.
     * @param array<string, string> $match
     */
    public function match(array $match): static
    {
        $this->internal->match = $match;
    
        return $this;
    }
    /**
     * A Route is a node that contains definitions of how to handle alerts. This is modified
     * from the upstream alertmanager in that it adds the ObjectMatchers property.
     * @param array<string, string> $matchRe
     */
    public function matchRe(array $matchRe): static
    {
        $this->internal->matchRe = $matchRe;
    
        return $this;
    }
    /**
     * A Route is a node that contains definitions of how to handle alerts. This is modified
     * from the upstream alertmanager in that it adds the ObjectMatchers property.
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alerting\Matcher>> $matchers
     */
    public function matchers(array $matchers): static
    {
            $matchersResources = [];
            foreach ($matchers as $r1) {
                    $matchersResources[] = $r1->build();
            }
        $this->internal->matchers = $matchersResources;
    
        return $this;
    }
    /**
     * A Route is a node that contains definitions of how to handle alerts. This is modified
     * from the upstream alertmanager in that it adds the ObjectMatchers property.
     * @param array<string> $muteTimeIntervals
     */
    public function muteTimeIntervals(array $muteTimeIntervals): static
    {
        $this->internal->muteTimeIntervals = $muteTimeIntervals;
    
        return $this;
    }
    /**
     * A Route is a node that contains definitions of how to handle alerts. This is modified
     * from the upstream alertmanager in that it adds the ObjectMatchers property.
     * @param array<array<string>> $objectMatchers
     */
    public function objectMatchers(array $objectMatchers): static
    {
        $this->internal->objectMatchers = $objectMatchers;
    
        return $this;
    }
    /**
     * A Route is a node that contains definitions of how to handle alerts. This is modified
     * from the upstream alertmanager in that it adds the ObjectMatchers property.
     */
    public function provenance(string $provenance): static
    {
        $this->internal->provenance = $provenance;
    
        return $this;
    }
    /**
     * A Route is a node that contains definitions of how to handle alerts. This is modified
     * from the upstream alertmanager in that it adds the ObjectMatchers property.
     */
    public function receiver(string $receiver): static
    {
        $this->internal->receiver = $receiver;
    
        return $this;
    }
    /**
     * A Route is a node that contains definitions of how to handle alerts. This is modified
     * from the upstream alertmanager in that it adds the ObjectMatchers property.
     */
    public function repeatInterval(string $repeatInterval): static
    {
        $this->internal->repeatInterval = $repeatInterval;
    
        return $this;
    }
    /**
     * A Route is a node that contains definitions of how to handle alerts. This is modified
     * from the upstream alertmanager in that it adds the ObjectMatchers property.
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alerting\NotificationPolicy>> $routes
     */
    public function routes(array $routes): static
    {
            $routesResources = [];
            foreach ($routes as $r1) {
                    $routesResources[] = $r1->build();
            }
        $this->internal->routes = $routesResources;
    
        return $this;
    }

}
