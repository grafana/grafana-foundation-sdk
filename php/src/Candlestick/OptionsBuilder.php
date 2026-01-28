<?php

namespace Grafana\Foundation\Candlestick;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Candlestick\Options>
 */
class OptionsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Candlestick\Options $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Candlestick\Options();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Candlestick\Options
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * Sets which dimensions are used for the visualization
     */
    public function mode(\Grafana\Foundation\Candlestick\VizDisplayMode $mode): static
    {
        $this->internal->mode = $mode;
    
        return $this;
    }

    /**
     * Sets the style of the candlesticks
     */
    public function candleStyle(\Grafana\Foundation\Candlestick\CandleStyle $candleStyle): static
    {
        $this->internal->candleStyle = $candleStyle;
    
        return $this;
    }

    /**
     * Sets the color strategy for the candlesticks
     */
    public function colorStrategy(\Grafana\Foundation\Candlestick\ColorStrategy $colorStrategy): static
    {
        $this->internal->colorStrategy = $colorStrategy;
    
        return $this;
    }

    /**
     * Map fields to appropriate dimension
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Candlestick\CandlestickFieldMap> $fields
     */
    public function fields(\Grafana\Foundation\Cog\Builder $fields): static
    {
        $fieldsResource = $fields->build();
        $this->internal->fields = $fieldsResource;
    
        return $this;
    }

    /**
     * Set which colors are used when the price movement is up or down
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Candlestick\CandlestickColors> $colors
     */
    public function colors(\Grafana\Foundation\Cog\Builder $colors): static
    {
        $colorsResource = $colors->build();
        $this->internal->colors = $colorsResource;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\VizLegendOptions> $legend
     */
    public function legend(\Grafana\Foundation\Cog\Builder $legend): static
    {
        $legendResource = $legend->build();
        $this->internal->legend = $legendResource;
    
        return $this;
    }

    /**
     * When enabled, all fields will be sent to the graph
     */
    public function includeAllFields(bool $includeAllFields): static
    {
        $this->internal->includeAllFields = $includeAllFields;
    
        return $this;
    }

}
