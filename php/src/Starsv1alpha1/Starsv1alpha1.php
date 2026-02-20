<?php

namespace Grafana\Foundation\Starsv1alpha1;

final class Starsv1alpha1
{

    /**
     * Creates a resource manifest from Stars.
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Starsv1alpha1\Stars> $stars
     */
    public static function manifest(string $name, \Grafana\Foundation\Cog\Builder $stars): \Grafana\Foundation\Resource\ManifestBuilder
    {
    	$builder = new \Grafana\Foundation\Resource\ManifestBuilder();
        $builder->apiVersion("preferences.grafana.app/v1alpha1");
        $builder->metadata(\Grafana\Foundation\Resource\Resource::named($name));
        $builder->kind("Stars");
        $builder->spec($stars->build());
    	return $builder;
    }

}
