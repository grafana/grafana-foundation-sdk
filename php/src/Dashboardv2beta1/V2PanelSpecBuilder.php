<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\V2PanelSpec>
 */
class V2PanelSpecBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2beta1\V2PanelSpec $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2beta1\V2PanelSpec();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2beta1\V2PanelSpec
     */
    public function build()
    {
        return $this->internal;
    }

    public function id(float $id): static
    {
        $this->internal->id = $id;
    
        return $this;
    }

    public function title(string $title): static
    {
        $this->internal->title = $title;
    
        return $this;
    }

    /**
     * Shown in a info icon tooltip next to panel title
     */
    public function description(string $description): static
    {
        $this->internal->description = $description;
    
        return $this;
    }

    /**
     * Shown in a sub header below the title.
     */
    public function subtitle(string $subtitle): static
    {
        $this->internal->subtitle = $subtitle;
    
        return $this;
    }

    /**
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\DataLink>> $links
     */
    public function links(array $links): static
    {
            $linksResources = [];
            foreach ($links as $r1) {
                    $linksResources[] = $r1->build();
            }
        $this->internal->links = $linksResources;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\V2QueryGroupKind> $data
     */
    public function data(\Grafana\Foundation\Cog\Builder $data): static
    {
        $dataResource = $data->build();
        $this->internal->data = $dataResource;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\VizConfigKind> $vizConfig
     */
    public function vizConfig(\Grafana\Foundation\Cog\Builder $vizConfig): static
    {
        $vizConfigResource = $vizConfig->build();
        $this->internal->vizConfig = $vizConfigResource;
    
        return $this;
    }

    public function transparent(bool $transparent): static
    {
        $this->internal->transparent = $transparent;
    
        return $this;
    }

}
