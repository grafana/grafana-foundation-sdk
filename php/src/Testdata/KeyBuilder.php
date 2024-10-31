<?php

namespace Grafana\Foundation\Testdata;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Testdata\Key>
 */
class KeyBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Testdata\Key $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Testdata\Key();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Testdata\Key
     */
    public function build()
    {
        return $this->internal;
    }

    public function tick(float $tick): static
    {
        $this->internal->tick = $tick;
    
        return $this;
    }
    public function type(string $type): static
    {
        $this->internal->type = $type;
    
        return $this;
    }
    public function uid(string $uid): static
    {
        $this->internal->uid = $uid;
    
        return $this;
    }

}
