// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.preferencesv1alpha1;

import com.grafana.foundation.resource.ManifestBuilder;
import com.grafana.foundation.resource.Manifest;
import com.grafana.foundation.resource.Resource;
public class Preferencesv1alpha1 {    
    /**
     * Creates a resource manifest from Preferences.
     */
    public static com.grafana.foundation.cog.Builder<Manifest> manifest(String name,com.grafana.foundation.cog.Builder<Preferences> preferences) {
        ManifestBuilder builder = new ManifestBuilder();
        builder.apiVersion("preferences.grafana.app/v1alpha1");
        builder.kind("Preferences");
        builder.metadata(Resource.named(name));
        builder.spec(preferences.build());
        return builder;
    }
}
