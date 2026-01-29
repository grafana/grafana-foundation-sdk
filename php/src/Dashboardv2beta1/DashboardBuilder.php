<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\Dashboard>
 */
class DashboardBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2beta1\Dashboard $internal;

    public function __construct(string $title)
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2beta1\Dashboard();
    $this->internal->title = $title;
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2beta1\Dashboard
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\AnnotationQueryKind>> $annotations
     */
    public function annotations(array $annotations): static
    {
            $annotationsResources = [];
            foreach ($annotations as $r1) {
                    $annotationsResources[] = $r1->build();
            }
        $this->internal->annotations = $annotationsResources;
    
        return $this;
    }

    /**
     * Configuration of dashboard cursor sync behavior.
     * "Off" for no shared crosshair or tooltip (default).
     * "Crosshair" for shared crosshair.
     * "Tooltip" for shared crosshair AND shared tooltip.
     */
    public function cursorSync(\Grafana\Foundation\Dashboardv2beta1\DashboardCursorSync $cursorSync): static
    {
        $this->internal->cursorSync = $cursorSync;
    
        return $this;
    }

    /**
     * Description of dashboard.
     */
    public function description(string $description): static
    {
        $this->internal->description = $description;
    
        return $this;
    }

    /**
     * Whether a dashboard is editable or not.
     */
    public function editable(bool $editable): static
    {
        $this->internal->editable = $editable;
    
        return $this;
    }

    /**
     * @param array<string, \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\PanelKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\LibraryPanelKind>> $elements
     */
    public function elements(array $elements): static
    {
            $elementsResources = [];
            foreach ($elements as $key1 => $val1) {
                    $elementsResources[$key1] = $val1->build();
            }
        $this->internal->elements = $elementsResources;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\PanelKind> $panelKind
     */
    public function panel(string $key, \Grafana\Foundation\Cog\Builder $panelKind): static
    {
        $panelKindResource = $panelKind->build();
        $this->internal->elements[$key] = $panelKindResource;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\LibraryPanelKind> $libraryPanelKind
     */
    public function libraryPanel(string $key, \Grafana\Foundation\Cog\Builder $libraryPanelKind): static
    {
        $libraryPanelKindResource = $libraryPanelKind->build();
        $this->internal->elements[$key] = $libraryPanelKindResource;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\GridLayoutKind> $gridLayoutKind
     */
    public function gridLayout(\Grafana\Foundation\Cog\Builder $gridLayoutKind): static
    {
        $gridLayoutKindResource = $gridLayoutKind->build();
        $this->internal->layout = $gridLayoutKindResource;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\RowsLayoutKind> $rowsLayoutKind
     */
    public function rowsLayout(\Grafana\Foundation\Cog\Builder $rowsLayoutKind): static
    {
        $rowsLayoutKindResource = $rowsLayoutKind->build();
        $this->internal->layout = $rowsLayoutKindResource;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutKind> $autoGridLayoutKind
     */
    public function autoGridLayout(\Grafana\Foundation\Cog\Builder $autoGridLayoutKind): static
    {
        $autoGridLayoutKindResource = $autoGridLayoutKind->build();
        $this->internal->layout = $autoGridLayoutKindResource;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\TabsLayoutKind> $tabsLayoutKind
     */
    public function tabsLayout(\Grafana\Foundation\Cog\Builder $tabsLayoutKind): static
    {
        $tabsLayoutKindResource = $tabsLayoutKind->build();
        $this->internal->layout = $tabsLayoutKindResource;
    
        return $this;
    }

    /**
     * Links with references to other dashboards or external websites.
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\DashboardLink>> $links
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
     * When set to true, the dashboard will redraw panels at an interval matching the pixel width.
     * This will keep data "moving left" regardless of the query refresh rate. This setting helps
     * avoid dashboards presenting stale live data.
     */
    public function liveNow(bool $liveNow): static
    {
        $this->internal->liveNow = $liveNow;
    
        return $this;
    }

    /**
     * When set to true, the dashboard will load all panels in the dashboard when it's loaded.
     */
    public function preload(bool $preload): static
    {
        $this->internal->preload = $preload;
    
        return $this;
    }

    /**
     * Plugins only. The version of the dashboard installed together with the plugin.
     * This is used to determine if the dashboard should be updated when the plugin is updated.
     */
    public function revision(int $revision): static
    {
        $this->internal->revision = $revision;
    
        return $this;
    }

    /**
     * Tags associated with dashboard.
     * @param array<string> $tags
     */
    public function tags(array $tags): static
    {
        $this->internal->tags = $tags;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\TimeSettingsSpec> $timeSettings
     */
    public function timeSettings(\Grafana\Foundation\Cog\Builder $timeSettings): static
    {
        $timeSettingsResource = $timeSettings->build();
        $this->internal->timeSettings = $timeSettingsResource;
    
        return $this;
    }

    /**
     * Title of dashboard.
     */
    public function title(string $title): static
    {
        $this->internal->title = $title;
    
        return $this;
    }

    /**
     * Configured template variables.
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\QueryVariableKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\TextVariableKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\ConstantVariableKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\DatasourceVariableKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\IntervalVariableKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\CustomVariableKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\GroupByVariableKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\AdhocVariableKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\SwitchVariableKind>> $variables
     */
    public function variables(array $variables): static
    {
            $variablesResources = [];
            foreach ($variables as $r1) {
                    $variablesResources[] = $r1->build();
            }
        $this->internal->variables = $variablesResources;
    
        return $this;
    }

    /**
     * Configured template variables.
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\QueryVariableKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\TextVariableKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\ConstantVariableKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\DatasourceVariableKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\IntervalVariableKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\CustomVariableKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\GroupByVariableKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\AdhocVariableKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\SwitchVariableKind> $variable
     */
    public function variable( $variable): static
    {
        $variableResource = $variable->build();
        $this->internal->variables[] = $variableResource;
    
        return $this;
    }

}
