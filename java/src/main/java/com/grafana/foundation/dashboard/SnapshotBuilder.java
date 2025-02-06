// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;


public class SnapshotBuilder implements com.grafana.foundation.cog.Builder<Snapshot> {
    protected final Snapshot internal;
    
    public SnapshotBuilder() {
        this.internal = new Snapshot();
    }
    public SnapshotBuilder expires(String expires) {
        this.internal.expires = expires;
        return this;
    }
    
    public SnapshotBuilder external(Boolean external) {
        this.internal.external = external;
        return this;
    }
    
    public SnapshotBuilder externalUrl(String externalUrl) {
        this.internal.externalUrl = externalUrl;
        return this;
    }
    
    public SnapshotBuilder id(Integer id) {
        this.internal.id = id;
        return this;
    }
    
    public SnapshotBuilder key(String key) {
        this.internal.key = key;
        return this;
    }
    
    public SnapshotBuilder name(String name) {
        this.internal.name = name;
        return this;
    }
    
    public SnapshotBuilder orgId(Integer orgId) {
        this.internal.orgId = orgId;
        return this;
    }
    
    public SnapshotBuilder url(String url) {
        this.internal.url = url;
        return this;
    }
    
    public SnapshotBuilder dashboard(com.grafana.foundation.cog.Builder<Dashboard> dashboard) {
        this.internal.dashboard = dashboard.build();
        return this;
    }
    public Snapshot build() {
        return this.internal;
    }
}
