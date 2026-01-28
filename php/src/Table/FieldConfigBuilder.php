<?php

namespace Grafana\Foundation\Table;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Table\FieldConfig>
 */
class FieldConfigBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Table\FieldConfig $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Table\FieldConfig();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Table\FieldConfig
     */
    public function build()
    {
        return $this->internal;
    }

    public function width(float $width): static
    {
        $this->internal->width = $width;
    
        return $this;
    }

    public function minWidth(float $minWidth): static
    {
        $this->internal->minWidth = $minWidth;
    
        return $this;
    }

    public function align(\Grafana\Foundation\Common\FieldTextAlignment $align): static
    {
        $this->internal->align = $align;
    
        return $this;
    }

    /**
     * This field is deprecated in favor of using cellOptions
     */
    public function displayMode(\Grafana\Foundation\Common\TableCellDisplayMode $displayMode): static
    {
        $this->internal->displayMode = $displayMode;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Common\TableAutoCellOptions|\Grafana\Foundation\Common\TableSparklineCellOptions|\Grafana\Foundation\Common\TableBarGaugeCellOptions|\Grafana\Foundation\Common\TableColoredBackgroundCellOptions|\Grafana\Foundation\Common\TableColorTextCellOptions|\Grafana\Foundation\Common\TableImageCellOptions|\Grafana\Foundation\Common\TableDataLinksCellOptions|\Grafana\Foundation\Common\TableJsonViewCellOptions $cellOptions
     */
    public function cellOptions( $cellOptions): static
    {
        $this->internal->cellOptions = $cellOptions;
    
        return $this;
    }

    /**
     * ?? default is missing or false ??
     */
    public function hidden(bool $hidden): static
    {
        $this->internal->hidden = $hidden;
    
        return $this;
    }

    public function inspect(bool $inspect): static
    {
        $this->internal->inspect = $inspect;
    
        return $this;
    }

    public function filterable(bool $filterable): static
    {
        $this->internal->filterable = $filterable;
    
        return $this;
    }

    /**
     * Hides any header for a column, useful for columns that show some static content or buttons.
     */
    public function hideHeader(bool $hideHeader): static
    {
        $this->internal->hideHeader = $hideHeader;
    
        return $this;
    }

}
