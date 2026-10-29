<?php

namespace Grafana\Foundation\Alertingv1beta1;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alertingv1beta1\PanelRef>
 */
class PanelRefBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Alertingv1beta1\PanelRef $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Alertingv1beta1\PanelRef();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Alertingv1beta1\PanelRef
     */
    public function build()
    {
        return $this->internal;
    }

    public function dashboardUID(string $dashboardUID): static
    {
        if (!(strlen($dashboardUID) >= 1)) {
            throw new \ValueError('strlen($dashboardUID) must be >= 1');
        }
        if (!preg_match("/^[a-zA-Z0-9_-]+$/", $dashboardUID)) {
            throw new \ValueError('$dashboardUID must match regex ^[a-zA-Z0-9_-]+$');
        }
        $this->internal->dashboardUID = $dashboardUID;
    
        return $this;
    }

    public function panelID(int $panelID): static
    {
        if (!($panelID > 0)) {
            throw new \ValueError('$panelID must be > 0');
        }
        $this->internal->panelID = $panelID;
    
        return $this;
    }

}
