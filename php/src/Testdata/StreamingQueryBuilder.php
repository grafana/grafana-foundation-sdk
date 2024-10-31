<?php

namespace Grafana\Foundation\Testdata;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Testdata\StreamingQuery>
 */
class StreamingQueryBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Testdata\StreamingQuery $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Testdata\StreamingQuery();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Testdata\StreamingQuery
     */
    public function build()
    {
        return $this->internal;
    }

    public function type(\Grafana\Foundation\Testdata\StreamingQueryType $type): static
    {
        $this->internal->type = $type;
    
        return $this;
    }
    public function speed(int $speed): static
    {
        $this->internal->speed = $speed;
    
        return $this;
    }
    public function spread(int $spread): static
    {
        $this->internal->spread = $spread;
    
        return $this;
    }
    public function noise(int $noise): static
    {
        $this->internal->noise = $noise;
    
        return $this;
    }
    public function bands(int $bands): static
    {
        $this->internal->bands = $bands;
    
        return $this;
    }
    public function url(string $url): static
    {
        $this->internal->url = $url;
    
        return $this;
    }

}
