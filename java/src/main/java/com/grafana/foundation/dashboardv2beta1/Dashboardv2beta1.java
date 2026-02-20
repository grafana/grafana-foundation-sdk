// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.grafana.foundation.resource.ManifestBuilder;
import com.grafana.foundation.resource.Manifest;
import com.grafana.foundation.resource.Resource;
public class Dashboardv2beta1 {    
    /**
     * Creates a resource manifest from a Dashboard.
     */
    public static com.grafana.foundation.cog.Builder<Manifest> manifest(String name,com.grafana.foundation.cog.Builder<Dashboard> dashboard) {
        ManifestBuilder builder = new ManifestBuilder();
        builder.apiVersion("dashboard.grafana.app/v2beta1");
        builder.kind("Dashboard");
        builder.metadata(Resource.named(name));
        builder.spec(dashboard.build());
        return builder;
    }
}
