// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;


public class NotebookLayoutItemKindBuilder implements com.grafana.foundation.cog.Builder<NotebookLayoutItemKind> {
    protected final NotebookLayoutItemKind internal;
    
    public NotebookLayoutItemKindBuilder() {
        this.internal = new NotebookLayoutItemKind();
        this.internal.kind = "NotebookLayoutItem";
    }
    public NotebookLayoutItemKindBuilder spec(com.grafana.foundation.cog.Builder<NotebookLayoutItemSpec> spec) {
    NotebookLayoutItemSpec specResource = spec.build();
        this.internal.spec = specResource;
        return this;
    }
    public NotebookLayoutItemKind build() {
        return this.internal;
    }
}
