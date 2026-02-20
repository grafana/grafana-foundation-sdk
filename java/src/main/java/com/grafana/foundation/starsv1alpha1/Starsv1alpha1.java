// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.starsv1alpha1;

import com.grafana.foundation.resource.ManifestBuilder;
import com.grafana.foundation.resource.Manifest;
import com.grafana.foundation.resource.Resource;
public class Starsv1alpha1 {    
    /**
     * Creates a resource manifest from Stars.
     */
    public static com.grafana.foundation.cog.Builder<Manifest> manifest(String name,com.grafana.foundation.cog.Builder<Stars> stars) {
        ManifestBuilder builder = new ManifestBuilder();
        builder.apiVersion("preferences.grafana.app/v1alpha1");
        builder.metadata(Resource.named(name));
        builder.kind("Stars");
        builder.spec(stars.build());
        return builder;
    }
}
