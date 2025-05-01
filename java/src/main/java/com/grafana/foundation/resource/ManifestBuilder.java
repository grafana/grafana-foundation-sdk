// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.resource;


public class ManifestBuilder implements com.grafana.foundation.cog.Builder<Manifest> {
    protected final Manifest internal;
    
    public ManifestBuilder() {
        this.internal = new Manifest();
    }
    public ManifestBuilder apiVersion(String apiVersion) {
        this.internal.apiVersion = apiVersion;
        return this;
    }
    
    public ManifestBuilder kind(String kind) {
        this.internal.kind = kind;
        return this;
    }
    
    public ManifestBuilder metadata(com.grafana.foundation.cog.Builder<Metadata> metadata) {
    Metadata metadataResource = metadata.build();
        this.internal.metadata = metadataResource;
        return this;
    }
    
    public ManifestBuilder spec(Object spec) {
        this.internal.spec = spec;
        return this;
    }
    public Manifest build() {
        return this.internal;
    }
}
