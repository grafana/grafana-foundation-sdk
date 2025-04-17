// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.resource;

import java.util.Map;

public class MetadataBuilder implements com.grafana.foundation.cog.Builder<Metadata> {
    protected final Metadata internal;
    
    public MetadataBuilder() {
        this.internal = new Metadata();
    }
    public MetadataBuilder name(String name) {
        this.internal.name = name;
        return this;
    }
    
    public MetadataBuilder namespace(String namespace) {
        this.internal.namespace = namespace;
        return this;
    }
    
    public MetadataBuilder labels(Map<String, String> labels) {
        this.internal.labels = labels;
        return this;
    }
    
    public MetadataBuilder annotations(Map<String, String> annotations) {
        this.internal.annotations = annotations;
        return this;
    }
    
    public MetadataBuilder uid(String uid) {
        this.internal.uid = uid;
        return this;
    }
    
    public MetadataBuilder resourceVersion(String resourceVersion) {
        this.internal.resourceVersion = resourceVersion;
        return this;
    }
    
    public MetadataBuilder generation(Long generation) {
        this.internal.generation = generation;
        return this;
    }
    
    public MetadataBuilder creationTimestamp(String creationTimestamp) {
        this.internal.creationTimestamp = creationTimestamp;
        return this;
    }
    
    public MetadataBuilder updateTimestamp(String updateTimestamp) {
        this.internal.updateTimestamp = updateTimestamp;
        return this;
    }
    
    public MetadataBuilder deletionTimestamp(String deletionTimestamp) {
        this.internal.deletionTimestamp = deletionTimestamp;
        return this;
    }
    public Metadata build() {
        return this.internal;
    }
}
