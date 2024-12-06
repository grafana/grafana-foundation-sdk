<?php

namespace Grafana\Foundation\Nodegraph;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Nodegraph\NodeOptions>
 */
class NodeOptionsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Nodegraph\NodeOptions $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Nodegraph\NodeOptions();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Nodegraph\NodeOptions
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * Unit for the main stat to override what ever is set in the data frame.
     */
    public function mainStatUnit(string $mainStatUnit): static
    {
        $this->internal->mainStatUnit = $mainStatUnit;
    
        return $this;
    }
    /**
     * Unit for the secondary stat to override what ever is set in the data frame.
     */
    public function secondaryStatUnit(string $secondaryStatUnit): static
    {
        $this->internal->secondaryStatUnit = $secondaryStatUnit;
    
        return $this;
    }
    /**
     * Define which fields are shown as part of the node arc (colored circle around the node).
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Nodegraph\ArcOption>> $arcs
     */
    public function arcs(array $arcs): static
    {
            $arcsResources = [];
            foreach ($arcs as $r1) {
                    $arcsResources[] = $r1->build();
            }
        $this->internal->arcs = $arcsResources;
    
        return $this;
    }

}
