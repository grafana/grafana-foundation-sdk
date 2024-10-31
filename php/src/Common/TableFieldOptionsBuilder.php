<?php

namespace Grafana\Foundation\Common;

/**
 * Field options for each field within a table (e.g 10, "The String", 64.20, etc.)
 * Generally defines alignment, filtering capabilties, display options, etc.
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\TableFieldOptions>
 */
class TableFieldOptionsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Common\TableFieldOptions $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Common\TableFieldOptions();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Common\TableFieldOptions
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
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\TableAutoCellOptions>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\TableSparklineCellOptions>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\TableBarGaugeCellOptions>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\TableColoredBackgroundCellOptions>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\TableColorTextCellOptions>|\Grafana\Foundation\Common\TableImageCellOptions|\Grafana\Foundation\Common\TableDataLinksCellOptions|\Grafana\Foundation\Common\TableJsonViewCellOptions $cellOptions
     */
    public function cellOptions( $cellOptions): static
    {
        /** @var \Grafana\Foundation\Common\TableAutoCellOptions|\Grafana\Foundation\Common\TableSparklineCellOptions|\Grafana\Foundation\Common\TableBarGaugeCellOptions|\Grafana\Foundation\Common\TableColoredBackgroundCellOptions|\Grafana\Foundation\Common\TableColorTextCellOptions|\Grafana\Foundation\Common\TableImageCellOptions|\Grafana\Foundation\Common\TableDataLinksCellOptions|\Grafana\Foundation\Common\TableJsonViewCellOptions $cellOptionsResource */
        $cellOptionsResource = $cellOptions instanceof \Grafana\Foundation\Cog\Builder ? $cellOptions->build() : $cellOptions;
        $this->internal->cellOptions = $cellOptionsResource;
    
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
