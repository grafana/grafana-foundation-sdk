<?php

namespace Grafana\Foundation\Alerting;

/**
 * EmbeddedContactPoint is the contact point type that is used
 * by grafanas embedded alertmanager implementation.
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alerting\ContactPoint>
 */
class ContactPointBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Alerting\ContactPoint $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Alerting\ContactPoint();
    }

    /**
     * @return \Grafana\Foundation\Alerting\ContactPoint
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * EmbeddedContactPoint is the contact point type that is used
     * by grafanas embedded alertmanager implementation.
     */
    public function disableResolveMessage(bool $disableResolveMessage): static
    {
        $this->internal->disableResolveMessage = $disableResolveMessage;
    
        return $this;
    }
    /**
     * EmbeddedContactPoint is the contact point type that is used
     * by grafanas embedded alertmanager implementation.
     */
    public function name(string $name): static
    {
        $this->internal->name = $name;
    
        return $this;
    }
    /**
     * EmbeddedContactPoint is the contact point type that is used
     * by grafanas embedded alertmanager implementation.
     */
    public function provenance(string $provenance): static
    {
        $this->internal->provenance = $provenance;
    
        return $this;
    }
    /**
     * EmbeddedContactPoint is the contact point type that is used
     * by grafanas embedded alertmanager implementation.
     * @param mixed $settings
     */
    public function settings( $settings): static
    {
        $this->internal->settings = $settings;
    
        return $this;
    }
    /**
     * EmbeddedContactPoint is the contact point type that is used
     * by grafanas embedded alertmanager implementation.
     */
    public function type(\Grafana\Foundation\Alerting\ContactPointType $type): static
    {
        $this->internal->type = $type;
    
        return $this;
    }
    /**
     * EmbeddedContactPoint is the contact point type that is used
     * by grafanas embedded alertmanager implementation.
     */
    public function uid(string $uid): static
    {
        $this->internal->uid = $uid;
    
        return $this;
    }

}
