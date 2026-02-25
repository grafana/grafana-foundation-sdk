// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.playlistv0alpha1;

import com.grafana.foundation.resource.ManifestBuilder;
import com.grafana.foundation.resource.Resource;
public class Playlistv0alpha1 {    
    /**
     * Creates a resource manifest from a Playlist.
     */
    public static ManifestBuilder manifest(String name,com.grafana.foundation.cog.Builder<Playlist> playlist) {
        ManifestBuilder builder = new ManifestBuilder();
        builder.apiVersion("playlist.grafana.app/playlistv0alpha1");
        builder.kind("Playlist");
        builder.metadata(Resource.named(name));
        builder.spec(playlist.build());
        return builder;
    }
}
