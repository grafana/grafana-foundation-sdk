<?php

namespace Grafana\Foundation\Folderv1beta1;

final class Folderv1beta1
{

    /**
     * Creates a resource manifest from a Folder.
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Folderv1beta1\Folder> $folder
     */
    public static function manifest(string $name, \Grafana\Foundation\Cog\Builder $folder): \Grafana\Foundation\Resource\ManifestBuilder
    {
    	$builder = new \Grafana\Foundation\Resource\ManifestBuilder();
        $builder->apiVersion("folder.grafana.app/v1beta1");
        $builder->kind("Folder");
        $builder->metadata(\Grafana\Foundation\Resource\Resource::named($name));
        $builder->spec($folder->build());
    	return $builder;
    }

}
