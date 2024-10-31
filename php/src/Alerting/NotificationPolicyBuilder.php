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

    public function continue(bool $continue): static
    {
        $this->internal->continue = $continue;
    
        return $this;
    }
    /**
     * @param array<string> $groupBy
     */
    public function groupBy(array $groupBy): static
    {
        $this->internal->groupBy = $groupBy;
    
        return $this;
    }
    public function groupInterval(string $groupInterval): static
    {
        $this->internal->groupInterval = $groupInterval;
    
        return $this;
    }
    public function groupWait(string $groupWait): static
    {
        $this->internal->groupWait = $groupWait;
    
        return $this;
    }
    /**
     * Deprecated. Remove before v1.0 release.
     * @param array<string, string> $match
     */
    public function match(array $match): static
    {
        $this->internal->match = $match;
    
        return $this;
    }
    /**
     * @param array<string, mixed> $matchRe
     */
    public function matchRe(array $matchRe): static
    {
        $this->internal->matchRe = $matchRe;
    
        return $this;
    }
    /**
     * Matchers is a slice of Matchers that is sortable, implements Stringer, and
     * provides a Matches method to match a LabelSet against all Matchers in the
     * slice. Note that some users of Matchers might require it to be sorted.
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
     * @param array<string> $muteTimeIntervals
     */
    public function muteTimeIntervals(array $muteTimeIntervals): static
    {
        $this->internal->muteTimeIntervals = $muteTimeIntervals;
    
        return $this;
    }
    /**
     * Matchers is a slice of Matchers that is sortable, implements Stringer, and
     * provides a Matches method to match a LabelSet against all Matchers in the
     * slice. Note that some users of Matchers might require it to be sorted.
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alerting\Matcher>> $objectMatchers
     */
    public function objectMatchers(array $objectMatchers): static
    {
            $objectMatchersResources = [];
            foreach ($objectMatchers as $r1) {
                    $objectMatchersResources[] = $r1->build();
            }
        $this->internal->objectMatchers = $objectMatchersResources;
    
        return $this;
    }
    public function provenance(string $provenance): static
    {
        $this->internal->provenance = $provenance;
    
        return $this;
    }
    public function receiver(string $receiver): static
    {
        $this->internal->receiver = $receiver;
    
        return $this;
    }
    public function repeatInterval(string $repeatInterval): static
    {
        $this->internal->repeatInterval = $repeatInterval;
    
        return $this;
    }
    /**
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
