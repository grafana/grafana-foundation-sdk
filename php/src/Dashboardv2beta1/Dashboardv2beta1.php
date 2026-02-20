<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class Dashboardv2beta1
{

    /**
     * Creates a resource manifest from a Dashboard.
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\Dashboard> $dashboard
     */
    public static function manifest(string $name, \Grafana\Foundation\Cog\Builder $dashboard): \Grafana\Foundation\Resource\ManifestBuilder
    {
    	$builder = new \Grafana\Foundation\Resource\ManifestBuilder();
        $builder->apiVersion("dashboard.grafana.app/v2beta1");
        $builder->kind("Dashboard");
        $builder->metadata(\Grafana\Foundation\Resource\Resource::named($name));
        $builder->spec($dashboard->build());
    	return $builder;
    }

}
