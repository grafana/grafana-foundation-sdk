<?php

namespace Grafana\Foundation\Dashboardv2;

final class Dashboardv2
{

    /**
     * Creates a resource manifest from a Dashboard.
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2\Dashboard> $spec
     */
    public static function manifest(string $name, \Grafana\Foundation\Cog\Builder $spec): \Grafana\Foundation\Resource\ManifestBuilder
    {
    	$builder = new \Grafana\Foundation\Resource\ManifestBuilder();
        $builder->apiVersion("dashboard.grafana.app/v2");
        $builder->kind("Dashboard");
        $builder->metadata(\Grafana\Foundation\Resource\Resource::named($name));
        $builder->spec($spec->build());
    	return $builder;
    }

    public static function rows(): \Grafana\Foundation\Dashboardv2\RowsBuilder
    {
    	$builder = new \Grafana\Foundation\Dashboardv2\RowsBuilder();
    	return $builder;
    }

    public static function row(string $title): \Grafana\Foundation\Dashboardv2\RowBuilder
    {
    	$builder = new \Grafana\Foundation\Dashboardv2\RowBuilder();
        $builder->title($title);
    	return $builder;
    }

    public static function autoGrid(): \Grafana\Foundation\Dashboardv2\AutoGridBuilder
    {
    	$builder = new \Grafana\Foundation\Dashboardv2\AutoGridBuilder();
    	return $builder;
    }

    public static function autoGridItem(string $name): \Grafana\Foundation\Dashboardv2\AutoGridItemBuilder
    {
    	$builder = new \Grafana\Foundation\Dashboardv2\AutoGridItemBuilder();
        $builder->name($name);
    	return $builder;
    }

    public static function tabs(): \Grafana\Foundation\Dashboardv2\TabsBuilder
    {
    	$builder = new \Grafana\Foundation\Dashboardv2\TabsBuilder();
    	return $builder;
    }

    public static function tab(string $title): \Grafana\Foundation\Dashboardv2\TabBuilder
    {
    	$builder = new \Grafana\Foundation\Dashboardv2\TabBuilder();
        $builder->title($title);
    	return $builder;
    }

    public static function grid(): \Grafana\Foundation\Dashboardv2\GridBuilder
    {
    	$builder = new \Grafana\Foundation\Dashboardv2\GridBuilder();
    	return $builder;
    }

    public static function gridItem(string $name): \Grafana\Foundation\Dashboardv2\GridItemBuilder
    {
    	$builder = new \Grafana\Foundation\Dashboardv2\GridItemBuilder();
        $builder->name($name);
    	return $builder;
    }

}
