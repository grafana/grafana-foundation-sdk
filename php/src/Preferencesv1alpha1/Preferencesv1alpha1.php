<?php

namespace Grafana\Foundation\Preferencesv1alpha1;

final class Preferencesv1alpha1
{

    /**
     * Creates a resource manifest from Preferences.
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Preferencesv1alpha1\Preferences> $preferences
     */
    public static function manifest(string $name, \Grafana\Foundation\Cog\Builder $preferences): \Grafana\Foundation\Resource\ManifestBuilder
    {
    	$builder = new \Grafana\Foundation\Resource\ManifestBuilder();
        $builder->apiVersion("preferences.grafana.app/v1alpha1");
        $builder->kind("Preferences");
        $builder->metadata(\Grafana\Foundation\Resource\Resource::named($name));
        $builder->spec($preferences->build());
    	return $builder;
    }

}
