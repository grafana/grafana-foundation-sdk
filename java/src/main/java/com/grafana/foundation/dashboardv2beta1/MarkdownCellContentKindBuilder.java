// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;


public class MarkdownCellContentKindBuilder implements com.grafana.foundation.cog.Builder<MarkdownCellContentKind> {
    protected final MarkdownCellContentKind internal;
    
    public MarkdownCellContentKindBuilder() {
        this.internal = new MarkdownCellContentKind();
        this.internal.kind = "Markdown";
    }
    public MarkdownCellContentKindBuilder spec(com.grafana.foundation.cog.Builder<MarkdownCellContentSpec> spec) {
    MarkdownCellContentSpec specResource = spec.build();
        this.internal.spec = specResource;
        return this;
    }
    public MarkdownCellContentKind build() {
        return this.internal;
    }
}
