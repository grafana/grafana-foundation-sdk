<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutKind>
 */
class AutoGridLayoutBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutKind $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutKind();
    $this->internal->kind = "AutoGridLayout";
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutKind
     */
    public function build()
    {
        return $this->internal;
    }

    public function maxColumnCount(float $maxColumnCount): static
    {
        $this->internal->spec->maxColumnCount = $maxColumnCount;
    
        return $this;
    }

    public function columnWidthMode(\Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutSpecColumnWidthMode $columnWidthMode): static
    {
        $this->internal->spec->columnWidthMode = $columnWidthMode;
    
        return $this;
    }

    public function columnWidth(float $columnWidth): static
    {
        $this->internal->spec->columnWidth = $columnWidth;
    
        return $this;
    }

    public function rowHeightMode(\Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutSpecRowHeightMode $rowHeightMode): static
    {
        $this->internal->spec->rowHeightMode = $rowHeightMode;
    
        return $this;
    }

    public function rowHeight(float $rowHeight): static
    {
        $this->internal->spec->rowHeight = $rowHeight;
    
        return $this;
    }

    public function fillScreen(bool $fillScreen): static
    {
        $this->internal->spec->fillScreen = $fillScreen;
    
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
        $this->internal->spec->items = $itemsResources;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutItemKind> $item
     */
    public function item(\Grafana\Foundation\Cog\Builder $item): static
    {
        $itemResource = $item->build();
        $this->internal->spec->items[] = $itemResource;
    
        return $this;
    }

}
