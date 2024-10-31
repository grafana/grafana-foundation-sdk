<?php

namespace Grafana\Foundation\Alerting;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alerting\NotificationTemplate>
 */
class NotificationTemplateBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Alerting\NotificationTemplate $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Alerting\NotificationTemplate();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Alerting\NotificationTemplate
     */
    public function build()
    {
        return $this->internal;
    }

    public function name(string $name): static
    {
        $this->internal->name = $name;
    
        return $this;
    }
    public function provenance(string $provenance): static
    {
        $this->internal->provenance = $provenance;
    
        return $this;
    }
    public function template(string $template): static
    {
        $this->internal->template = $template;
    
        return $this;
    }

}
