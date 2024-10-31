<?php

namespace Grafana\Foundation\Candlestick;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Candlestick\CandlestickFieldMap>
 */
class CandlestickFieldMapBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Candlestick\CandlestickFieldMap $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Candlestick\CandlestickFieldMap();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Candlestick\CandlestickFieldMap
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * Corresponds to the starting value of the given period
     */
    public function open(string $open): static
    {
        $this->internal->open = $open;
    
        return $this;
    }
    /**
     * Corresponds to the highest value of the given period
     */
    public function high(string $high): static
    {
        $this->internal->high = $high;
    
        return $this;
    }
    /**
     * Corresponds to the lowest value of the given period
     */
    public function low(string $low): static
    {
        $this->internal->low = $low;
    
        return $this;
    }
    /**
     * Corresponds to the final (end) value of the given period
     */
    public function close(string $close): static
    {
        $this->internal->close = $close;
    
        return $this;
    }
    /**
     * Corresponds to the sample count in the given period. (e.g. number of trades)
     */
    public function volume(string $volume): static
    {
        $this->internal->volume = $volume;
    
        return $this;
    }

}
