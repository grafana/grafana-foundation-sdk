<?php

namespace Grafana\Foundation\Athena;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Athena\ConnectionArgs>
 */
class ConnectionArgsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Athena\ConnectionArgs $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Athena\ConnectionArgs();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Athena\ConnectionArgs
     */
    public function build()
    {
        return $this->internal;
    }

    public function region(string $region): static
    {
        $this->internal->region = $region;
    
        return $this;
    }
    public function catalog(string $catalog): static
    {
        $this->internal->catalog = $catalog;
    
        return $this;
    }
    public function database(string $database): static
    {
        $this->internal->database = $database;
    
        return $this;
    }
    public function resultReuseEnabled(bool $resultReuseEnabled): static
    {
        $this->internal->resultReuseEnabled = $resultReuseEnabled;
    
        return $this;
    }
    public function resultReuseMaxAgeInMinutes(float $resultReuseMaxAgeInMinutes): static
    {
        $this->internal->resultReuseMaxAgeInMinutes = $resultReuseMaxAgeInMinutes;
    
        return $this;
    }

}
