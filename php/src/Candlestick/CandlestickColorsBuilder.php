<?php

namespace Grafana\Foundation\Candlestick;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Candlestick\CandlestickColors>
 */
class CandlestickColorsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Candlestick\CandlestickColors $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Candlestick\CandlestickColors();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Candlestick\CandlestickColors
     */
    public function build()
    {
        return $this->internal;
    }

    public function up(string $up): static
    {
        $this->internal->up = $up;
    
        return $this;
    }
    public function down(string $down): static
    {
        $this->internal->down = $down;
    
        return $this;
    }
    public function flat(string $flat): static
    {
        $this->internal->flat = $flat;
    
        return $this;
    }

}
