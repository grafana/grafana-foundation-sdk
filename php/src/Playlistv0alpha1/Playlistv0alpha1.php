<?php

namespace Grafana\Foundation\Playlistv0alpha1;

final class Playlistv0alpha1
{

    /**
     * Creates a resource manifest from a Playlist.
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Playlistv0alpha1\Playlist> $playlist
     */
    public static function manifest(string $name, \Grafana\Foundation\Cog\Builder $playlist): \Grafana\Foundation\Resource\ManifestBuilder
    {
    	$builder = new \Grafana\Foundation\Resource\ManifestBuilder();
        $builder->apiVersion("playlist.grafana.app/playlistv0alpha1");
        $builder->kind("Playlist");
        $builder->metadata(\Grafana\Foundation\Resource\Resource::named($name));
        $builder->spec($playlist->build());
    	return $builder;
    }

}
