<?php

namespace Grafana\Foundation\Resource;

final class Resource
{

    /**
     * Creates metadata for a named resource.
     */
    public static function named(string $name): \Grafana\Foundation\Resource\MetadataBuilder
    {
    	$builder = new \Grafana\Foundation\Resource\MetadataBuilder();
        $builder->name($name);
    	return $builder;
    }

}
