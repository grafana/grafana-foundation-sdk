<?php

namespace Grafana\Foundation\Alertingv0alpha1;

/**
 * #RuleSearchHitFields is the per-kind field payload returned on each hit.
 * It carries the union of alert- and recording-rule search fields; only the
 * fields relevant to a hit's kind are populated. This maps to the kind's
 * declared searchFields.
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alertingv0alpha1\RuleSearchHitFields>
 */
class RuleSearchHitFieldsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Alertingv0alpha1\RuleSearchHitFields $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Alertingv0alpha1\RuleSearchHitFields();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Alertingv0alpha1\RuleSearchHitFields
     */
    public function build()
    {
        return $this->internal;
    }

    public function title(string $title): static
    {
        $this->internal->title = $title;
    
        return $this;
    }

    public function folder(string $folder): static
    {
        $this->internal->folder = $folder;
    
        return $this;
    }

    public function type(string $type): static
    {
        $this->internal->type = $type;
    
        return $this;
    }

    public function interval(string $interval): static
    {
        $this->internal->interval = $interval;
    
        return $this;
    }

    public function paused(bool $paused): static
    {
        $this->internal->paused = $paused;
    
        return $this;
    }

    /**
     * @param array<string, string> $labels
     */
    public function labels(array $labels): static
    {
        $this->internal->labels = $labels;
    
        return $this;
    }

    /**
     * @param array<string> $datasourceUIDs
     */
    public function datasourceUIDs(array $datasourceUIDs): static
    {
        $this->internal->datasourceUIDs = $datasourceUIDs;
    
        return $this;
    }

    /**
     * Alert-rule fields.
     * @param array<string, string> $annotations
     */
    public function annotations(array $annotations): static
    {
        $this->internal->annotations = $annotations;
    
        return $this;
    }

    public function for(string $for): static
    {
        $this->internal->for = $for;
    
        return $this;
    }

    public function keepFiringFor(string $keepFiringFor): static
    {
        $this->internal->keepFiringFor = $keepFiringFor;
    
        return $this;
    }

    public function dashboardUID(string $dashboardUID): static
    {
        $this->internal->dashboardUID = $dashboardUID;
    
        return $this;
    }

    public function panelID(int $panelID): static
    {
        $this->internal->panelID = $panelID;
    
        return $this;
    }

    public function receiver(string $receiver): static
    {
        $this->internal->receiver = $receiver;
    
        return $this;
    }

    public function notificationType(string $notificationType): static
    {
        $this->internal->notificationType = $notificationType;
    
        return $this;
    }

    public function routingTree(string $routingTree): static
    {
        $this->internal->routingTree = $routingTree;
    
        return $this;
    }

    /**
     * Recording-rule fields.
     */
    public function metric(string $metric): static
    {
        $this->internal->metric = $metric;
    
        return $this;
    }

    public function targetDatasourceUID(string $targetDatasourceUID): static
    {
        $this->internal->targetDatasourceUID = $targetDatasourceUID;
    
        return $this;
    }

}
