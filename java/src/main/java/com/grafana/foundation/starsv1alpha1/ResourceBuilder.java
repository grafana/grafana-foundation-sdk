// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.starsv1alpha1;

import java.util.List;

public class ResourceBuilder implements com.grafana.foundation.cog.Builder<Resource> {
    protected final Resource internal;
    
    public ResourceBuilder() {
        this.internal = new Resource();
    }
    public ResourceBuilder group(String group) {
        this.internal.group = group;
        return this;
    }
    
    public ResourceBuilder kind(String kind) {
        this.internal.kind = kind;
        return this;
    }
    
    public ResourceBuilder names(List<String> names) {
        this.internal.names = names;
        return this;
    }
    public Resource build() {
        return this.internal;
    }
}
