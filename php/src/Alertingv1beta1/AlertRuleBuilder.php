<?php

namespace Grafana\Foundation\Alertingv1beta1;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alertingv1beta1\AlertRuleSpec>
 */
class AlertRuleBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Alertingv1beta1\AlertRuleSpec $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Alertingv1beta1\AlertRuleSpec();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Alertingv1beta1\AlertRuleSpec
     */
    public function build()
    {
        return $this->internal;
    }

    public function title(string $title): static
    {
        $this->internal->title = $title;
    
        return $this;
    }

    public function paused(bool $paused): static
    {
        $this->internal->paused = $paused;
    
        return $this;
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
     * @param array<string, string> $labels
     */
    public function labels(array $labels): static
    {
        $this->internal->labels = $labels;
    
        return $this;
    }

    /**
     * @param array<string, string> $annotations
     */
    public function annotations(array $annotations): static
    {
        $this->internal->annotations = $annotations;
    
        return $this;
    }

    public function for(string $for): static
    {
        $this->internal->for = $for;
    
        return $this;
    }

    public function keepFiringFor(string $keepFiringFor): static
    {
        $this->internal->keepFiringFor = $keepFiringFor;
    
        return $this;
    }

    public function missingSeriesEvalsToResolve(int $missingSeriesEvalsToResolve): static
    {
        if (!($missingSeriesEvalsToResolve >= 0)) {
            throw new \ValueError('$missingSeriesEvalsToResolve must be >= 0');
        }
        $this->internal->missingSeriesEvalsToResolve = $missingSeriesEvalsToResolve;
    
        return $this;
    }

    public function noDataState(\Grafana\Foundation\Alertingv1beta1\NoDataState $noDataState): static
    {
        $this->internal->noDataState = $noDataState;
    
        return $this;
    }

    public function execErrState(\Grafana\Foundation\Alertingv1beta1\ExecErrState $execErrState): static
    {
        $this->internal->execErrState = $execErrState;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alertingv1beta1\SimplifiedRouting>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alertingv1beta1\NamedRoutingTree> $notificationSettings
     */
    public function notificationSettings( $notificationSettings): static
    {
        $notificationSettingsResource = $notificationSettings->build();
        $this->internal->notificationSettings = $notificationSettingsResource;
    
        return $this;
    }

    /**
     * @param array<string, \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alertingv1beta1\Expression>> $expressions
     */
    public function expressions(array $expressions): static
    {
            $expressionsResources = [];
            foreach ($expressions as $key1 => $val1) {
                    $expressionsResources[$key1] = $val1->build();
            }
        $this->internal->expressions = $expressionsResources;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alertingv1beta1\PanelRef> $panelRef
     */
    public function panelRef(\Grafana\Foundation\Cog\Builder $panelRef): static
    {
        $panelRefResource = $panelRef->build();
        $this->internal->panelRef = $panelRefResource;
    
        return $this;
    }

}
