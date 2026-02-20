// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.resource;

public class Resource {    
    /**
     * Creates metadata for a named resource.
     */
    public static com.grafana.foundation.cog.Builder<Metadata> named(String name) {
        MetadataBuilder builder = new MetadataBuilder();
        builder.name(name);
        return builder;
    }
}
