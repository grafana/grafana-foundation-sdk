<?php

namespace Grafana\Foundation\Common;

/**
 * Footer options
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\TableFooterOptions>
 */
class TableFooterOptionsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Common\TableFooterOptions $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Common\TableFooterOptions();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Common\TableFooterOptions
     */
    public function build()
    {
        return $this->internal;
    }

    public function show(bool $show): static
    {
        $this->internal->show = $show;
    
        return $this;
    }
    /**
     * actually 1 value
     * @param array<string> $reducer
     */
    public function reducer(array $reducer): static
    {
        $this->internal->reducer = $reducer;
    
        return $this;
    }
    /**
     * @param array<string> $fields
     */
    public function fields(array $fields): static
    {
        $this->internal->fields = $fields;
    
        return $this;
    }
    public function enablePagination(bool $enablePagination): static
    {
        $this->internal->enablePagination = $enablePagination;
    
        return $this;
    }
    public function countRows(bool $countRows): static
    {
        $this->internal->countRows = $countRows;
    
        return $this;
    }

}
