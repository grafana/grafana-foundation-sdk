<?php

namespace Grafana\Foundation\Testdata;

/**
 * TODO: Should this live here given it's not used in the dataquery?
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Testdata\Scenario>
 */
class ScenarioBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Testdata\Scenario $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Testdata\Scenario();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Testdata\Scenario
     */
    public function build()
    {
        return $this->internal;
    }

    public function id(string $id): static
    {
        $this->internal->id = $id;
    
        return $this;
    }
    public function name(string $name): static
    {
        $this->internal->name = $name;
    
        return $this;
    }
    public function stringInput(string $stringInput): static
    {
        $this->internal->stringInput = $stringInput;
    
        return $this;
    }
    public function description(string $description): static
    {
        $this->internal->description = $description;
    
        return $this;
    }
    public function hideAliasField(bool $hideAliasField): static
    {
        $this->internal->hideAliasField = $hideAliasField;
    
        return $this;
    }

}
