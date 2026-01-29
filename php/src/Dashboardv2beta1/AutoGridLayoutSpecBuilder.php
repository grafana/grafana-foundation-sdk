<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutSpec>
 */
class AutoGridLayoutSpecBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutSpec $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutSpec();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutSpec
     */
    public function build()
    {
        return $this->internal;
    }

    public function maxColumnCount(float $maxColumnCount): static
    {
        $this->internal->maxColumnCount = $maxColumnCount;
    
        return $this;
    }

    public function columnWidthMode(\Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutSpecColumnWidthMode $columnWidthMode): static
    {
        $this->internal->columnWidthMode = $columnWidthMode;
    
        return $this;
    }

    public function columnWidth(float $columnWidth): static
    {
        $this->internal->columnWidth = $columnWidth;
    
        return $this;
    }

    public function rowHeightMode(\Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutSpecRowHeightMode $rowHeightMode): static
    {
        $this->internal->rowHeightMode = $rowHeightMode;
    
        return $this;
    }

    public function rowHeight(float $rowHeight): static
    {
        $this->internal->rowHeight = $rowHeight;
    
        return $this;
    }

    public function fillScreen(bool $fillScreen): static
    {
        $this->internal->fillScreen = $fillScreen;
    
        return $this;
    }

    /**
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutItemKind>> $items
     */
    public function items(array $items): static
    {
            $itemsResources = [];
            foreach ($items as $r1) {
                    $itemsResources[] = $r1->build();
            }
        $this->internal->items = $itemsResources;
    
        return $this;
    }

}
