<?php

namespace Grafana\Foundation\Canvas;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Canvas\CanvasElementOptions>
 */
class CanvasElementOptionsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Canvas\CanvasElementOptions $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Canvas\CanvasElementOptions();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Canvas\CanvasElementOptions
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
    public function type(string $type): static
    {
        $this->internal->type = $type;
    
        return $this;
    }
    /**
     * TODO: figure out how to define this (element config(s))
     * @param mixed $config
     */
    public function config( $config): static
    {
        $this->internal->config = $config;
    
        return $this;
    }
    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Canvas\Constraint> $constraint
     */
    public function constraint(\Grafana\Foundation\Cog\Builder $constraint): static
    {
        $constraintResource = $constraint->build();
        $this->internal->constraint = $constraintResource;
    
        return $this;
    }
    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Canvas\Placement> $placement
     */
    public function placement(\Grafana\Foundation\Cog\Builder $placement): static
    {
        $placementResource = $placement->build();
        $this->internal->placement = $placementResource;
    
        return $this;
    }
    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Canvas\BackgroundConfig> $background
     */
    public function background(\Grafana\Foundation\Cog\Builder $background): static
    {
        $backgroundResource = $background->build();
        $this->internal->background = $backgroundResource;
    
        return $this;
    }
    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Canvas\LineConfig> $border
     */
    public function border(\Grafana\Foundation\Cog\Builder $border): static
    {
        $borderResource = $border->build();
        $this->internal->border = $borderResource;
    
        return $this;
    }
    /**
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Canvas\CanvasConnection>> $connections
     */
    public function connections(array $connections): static
    {
            $connectionsResources = [];
            foreach ($connections as $r1) {
                    $connectionsResources[] = $r1->build();
            }
        $this->internal->connections = $connectionsResources;
    
        return $this;
    }

}
