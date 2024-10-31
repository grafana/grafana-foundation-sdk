<?php

namespace Grafana\Foundation\Dashboard;

/**
 * Links with references to other dashboards or external resources
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\DashboardLink>
 */
class DashboardLinkBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboard\DashboardLink $internal;

    public function __construct(string $title)
    {
    	$this->internal = new \Grafana\Foundation\Dashboard\DashboardLink();
    $this->internal->title = $title;
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboard\DashboardLink
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * Title to display with the link
     */
    public function title(string $title): static
    {
        $this->internal->title = $title;
    
        return $this;
    }
    /**
     * Link type. Accepted values are dashboards (to refer to another dashboard) and link (to refer to an external resource)
     */
    public function type(\Grafana\Foundation\Dashboard\DashboardLinkType $type): static
    {
        $this->internal->type = $type;
    
        return $this;
    }
    /**
     * Icon name to be displayed with the link
     */
    public function icon(string $icon): static
    {
        $this->internal->icon = $icon;
    
        return $this;
    }
    /**
     * Tooltip to display when the user hovers their mouse over it
     */
    public function tooltip(string $tooltip): static
    {
        $this->internal->tooltip = $tooltip;
    
        return $this;
    }
    /**
     * Link URL. Only required/valid if the type is link
     */
    public function url(string $url): static
    {
        $this->internal->url = $url;
    
        return $this;
    }
    /**
     * List of tags to limit the linked dashboards. If empty, all dashboards will be displayed. Only valid if the type is dashboards
     * @param array<string> $tags
     */
    public function tags(array $tags): static
    {
        $this->internal->tags = $tags;
    
        return $this;
    }
    /**
     * If true, all dashboards links will be displayed in a dropdown. If false, all dashboards links will be displayed side by side. Only valid if the type is dashboards
     */
    public function asDropdown(bool $asDropdown): static
    {
        $this->internal->asDropdown = $asDropdown;
    
        return $this;
    }
    /**
     * If true, the link will be opened in a new tab
     */
    public function targetBlank(bool $targetBlank): static
    {
        $this->internal->targetBlank = $targetBlank;
    
        return $this;
    }
    /**
     * If true, includes current template variables values in the link as query params
     */
    public function includeVars(bool $includeVars): static
    {
        $this->internal->includeVars = $includeVars;
    
        return $this;
    }
    /**
     * If true, includes current time range in the link as query params
     */
    public function keepTime(bool $keepTime): static
    {
        $this->internal->keepTime = $keepTime;
    
        return $this;
    }

}
