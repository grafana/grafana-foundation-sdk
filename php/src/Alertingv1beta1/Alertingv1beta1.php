<?php

namespace Grafana\Foundation\Alertingv1beta1;

final class Alertingv1beta1
{

    /**
     * Creates a resource manifest from an AlertRule.
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alertingv1beta1\AlertRuleSpec> $folder
     */
    public static function alertRuleManifest(string $name, \Grafana\Foundation\Cog\Builder $folder): \Grafana\Foundation\Resource\ManifestBuilder
    {
    	$builder = new \Grafana\Foundation\Resource\ManifestBuilder();
        $builder->apiVersion("rules.alerting.grafana.app/v1beta1");
        $builder->kind("AlertRule");
        $builder->metadata(\Grafana\Foundation\Resource\Resource::named($name));
        $builder->spec($folder->build());
    	return $builder;
    }

    /**
     * Creates a resource manifest from a RecordingRule.
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alertingv1beta1\RecordingRuleSpec> $folder
     */
    public static function recordingRuleManifest(string $name, \Grafana\Foundation\Cog\Builder $folder): \Grafana\Foundation\Resource\ManifestBuilder
    {
    	$builder = new \Grafana\Foundation\Resource\ManifestBuilder();
        $builder->apiVersion("rules.alerting.grafana.app/v1beta1");
        $builder->kind("RecordingRule");
        $builder->metadata(\Grafana\Foundation\Resource\Resource::named($name));
        $builder->spec($folder->build());
    	return $builder;
    }

    /**
     * Creates a resource manifest from a RuleSequence.
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alertingv1beta1\RuleSequenceSpec> $folder
     */
    public static function ruleSequenceManifest(string $name, \Grafana\Foundation\Cog\Builder $folder): \Grafana\Foundation\Resource\ManifestBuilder
    {
    	$builder = new \Grafana\Foundation\Resource\ManifestBuilder();
        $builder->apiVersion("rules.alerting.grafana.app/v1beta1");
        $builder->kind("RuleSequence");
        $builder->metadata(\Grafana\Foundation\Resource\Resource::named($name));
        $builder->spec($folder->build());
    	return $builder;
    }

}
