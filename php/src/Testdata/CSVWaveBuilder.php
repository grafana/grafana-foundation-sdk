<?php

namespace Grafana\Foundation\Testdata;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Testdata\CSVWave>
 */
class CSVWaveBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Testdata\CSVWave $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Testdata\CSVWave();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Testdata\CSVWave
     */
    public function build()
    {
        return $this->internal;
    }

    public function timeStep(int $timeStep): static
    {
        $this->internal->timeStep = $timeStep;
    
        return $this;
    }
    public function name(string $name): static
    {
        $this->internal->name = $name;
    
        return $this;
    }
    public function valuesCSV(string $valuesCSV): static
    {
        $this->internal->valuesCSV = $valuesCSV;
    
        return $this;
    }
    public function labels(string $labels): static
    {
        $this->internal->labels = $labels;
    
        return $this;
    }

}
