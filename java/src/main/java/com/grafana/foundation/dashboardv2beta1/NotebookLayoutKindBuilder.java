// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;


public class NotebookLayoutKindBuilder implements com.grafana.foundation.cog.Builder<NotebookLayoutKind> {
    protected final NotebookLayoutKind internal;
    
    public NotebookLayoutKindBuilder() {
        this.internal = new NotebookLayoutKind();
        this.internal.kind = "NotebookLayout";
    }
    public NotebookLayoutKindBuilder spec(com.grafana.foundation.cog.Builder<NotebookLayoutSpec> spec) {
    NotebookLayoutSpec specResource = spec.build();
        this.internal.spec = specResource;
        return this;
    }
    public NotebookLayoutKind build() {
        return this.internal;
    }
}
