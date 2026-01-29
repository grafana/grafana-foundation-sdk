<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\PanelKind>
 */
class PanelBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2beta1\PanelKind $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2beta1\PanelKind();
    $this->internal->kind = "Panel";
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2beta1\PanelKind
     */
    public function build()
    {
        return $this->internal;
    }

    public function id(float $id): static
    {
        $this->internal->spec->id = $id;
    
        return $this;
    }

    public function title(string $title): static
    {
        $this->internal->spec->title = $title;
    
        return $this;
    }

    public function description(string $description): static
    {
        $this->internal->spec->description = $description;
    
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
        $this->internal->spec->links = $linksResources;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\QueryGroupKind> $data
     */
    public function data(\Grafana\Foundation\Cog\Builder $data): static
    {
        $dataResource = $data->build();
        $this->internal->spec->data = $dataResource;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\VizConfigKind> $vizConfig
     */
    public function visualization(\Grafana\Foundation\Cog\Builder $vizConfig): static
    {
        $vizConfigResource = $vizConfig->build();
        $this->internal->spec->vizConfig = $vizConfigResource;
    
        return $this;
    }

    public function transparent(bool $transparent): static
    {
        $this->internal->spec->transparent = $transparent;
    
        return $this;
    }

}
