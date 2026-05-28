<?php

namespace Grafana\Foundation\Dashboardv2;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2\TabsLayoutTabKind>
 */
class TabBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2\TabsLayoutTabKind $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2\TabsLayoutTabKind();
    $this->internal->kind = "TabsLayoutTab";
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2\TabsLayoutTabKind
     */
    public function build()
    {
        return $this->internal;
    }

    public function title(string $title): static
    {
        $this->internal->spec->title = $title;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2\GridLayoutKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2\RowsLayoutKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2\AutoGridLayoutKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2\TabsLayoutKind> $layout
     */
    public function layout( $layout): static
    {
        $layoutResource = $layout->build();
        $this->internal->spec->layout = $layoutResource;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2\ConditionalRenderingGroupKind> $conditionalRendering
     */
    public function conditionalRendering(\Grafana\Foundation\Cog\Builder $conditionalRendering): static
    {
        $conditionalRenderingResource = $conditionalRendering->build();
        $this->internal->spec->conditionalRendering = $conditionalRenderingResource;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2\TabRepeatOptions> $repeat
     */
    public function repeat(\Grafana\Foundation\Cog\Builder $repeat): static
    {
        $repeatResource = $repeat->build();
        $this->internal->spec->repeat = $repeatResource;
    
        return $this;
    }

    /**
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2\QueryVariableKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2\TextVariableKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2\ConstantVariableKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2\DatasourceVariableKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2\IntervalVariableKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2\CustomVariableKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2\GroupByVariableKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2\AdhocVariableKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2\SwitchVariableKind>> $variables
     */
    public function variables(array $variables): static
    {
            $variablesResources = [];
            foreach ($variables as $r1) {
                    $variablesResources[] = $r1->build();
            }
        $this->internal->spec->variables = $variablesResources;
    
        return $this;
    }

}
