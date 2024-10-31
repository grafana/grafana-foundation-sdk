<?php

namespace Grafana\Foundation\Testdata;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Testdata\PulseWaveQuery>
 */
class PulseWaveQueryBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Testdata\PulseWaveQuery $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Testdata\PulseWaveQuery();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Testdata\PulseWaveQuery
     */
    public function build()
    {
        return $this->internal;
    }

    public function offCount(int $offCount): static
    {
        $this->internal->offCount = $offCount;
    
        return $this;
    }
    public function offValue(float $offValue): static
    {
        $this->internal->offValue = $offValue;
    
        return $this;
    }
    public function onCount(int $onCount): static
    {
        $this->internal->onCount = $onCount;
    
        return $this;
    }
    public function onValue(float $onValue): static
    {
        $this->internal->onValue = $onValue;
    
        return $this;
    }
    public function timeStep(int $timeStep): static
    {
        $this->internal->timeStep = $timeStep;
    
        return $this;
    }

}
