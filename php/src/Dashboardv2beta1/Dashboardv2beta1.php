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

    public static function rows(): \Grafana\Foundation\Dashboardv2beta1\RowsBuilder
    {
    	$builder = new \Grafana\Foundation\Dashboardv2beta1\RowsBuilder();
    	return $builder;
    }

    public static function row(string $title): \Grafana\Foundation\Dashboardv2beta1\RowBuilder
    {
    	$builder = new \Grafana\Foundation\Dashboardv2beta1\RowBuilder();
        $builder->title($title);
    	return $builder;
    }

    public static function autoGrid(): \Grafana\Foundation\Dashboardv2beta1\AutoGridBuilder
    {
    	$builder = new \Grafana\Foundation\Dashboardv2beta1\AutoGridBuilder();
    	return $builder;
    }

    public static function autoGridItem(string $name): \Grafana\Foundation\Dashboardv2beta1\AutoGridItemBuilder
    {
    	$builder = new \Grafana\Foundation\Dashboardv2beta1\AutoGridItemBuilder();
        $builder->name($name);
    	return $builder;
    }

    public static function tabs(): \Grafana\Foundation\Dashboardv2beta1\TabsBuilder
    {
    	$builder = new \Grafana\Foundation\Dashboardv2beta1\TabsBuilder();
    	return $builder;
    }

    public static function tab(string $title): \Grafana\Foundation\Dashboardv2beta1\TabBuilder
    {
    	$builder = new \Grafana\Foundation\Dashboardv2beta1\TabBuilder();
        $builder->title($title);
    	return $builder;
    }

    public static function grid(): \Grafana\Foundation\Dashboardv2beta1\GridBuilder
    {
    	$builder = new \Grafana\Foundation\Dashboardv2beta1\GridBuilder();
    	return $builder;
    }

    public static function gridItem(string $name): \Grafana\Foundation\Dashboardv2beta1\GridItemBuilder
    {
    	$builder = new \Grafana\Foundation\Dashboardv2beta1\GridItemBuilder();
        $builder->name($name);
    	return $builder;
    }

}
