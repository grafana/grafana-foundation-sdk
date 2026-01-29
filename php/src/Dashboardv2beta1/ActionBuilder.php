<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\Action>
 */
class ActionBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2beta1\Action $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2beta1\Action();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2beta1\Action
     */
    public function build()
    {
        return $this->internal;
    }

    public function type(\Grafana\Foundation\Dashboardv2beta1\ActionType $type): static
    {
        $this->internal->type = $type;
    
        return $this;
    }

    public function title(string $title): static
    {
        $this->internal->title = $title;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\FetchOptions> $fetch
     */
    public function fetch(\Grafana\Foundation\Cog\Builder $fetch): static
    {
        $fetchResource = $fetch->build();
        $this->internal->fetch = $fetchResource;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\InfinityOptions> $infinity
     */
    public function infinity(\Grafana\Foundation\Cog\Builder $infinity): static
    {
        $infinityResource = $infinity->build();
        $this->internal->infinity = $infinityResource;
    
        return $this;
    }

    public function confirmation(string $confirmation): static
    {
        $this->internal->confirmation = $confirmation;
    
        return $this;
    }

    public function oneClick(bool $oneClick): static
    {
        $this->internal->oneClick = $oneClick;
    
        return $this;
    }

    /**
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\ActionVariable>> $variables
     */
    public function variables(array $variables): static
    {
            $variablesResources = [];
            foreach ($variables as $r1) {
                    $variablesResources[] = $r1->build();
            }
        $this->internal->variables = $variablesResources;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1ActionStyle> $style
     */
    public function style(\Grafana\Foundation\Cog\Builder $style): static
    {
        $styleResource = $style->build();
        $this->internal->style = $styleResource;
    
        return $this;
    }

}
