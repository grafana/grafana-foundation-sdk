// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;


public class MarkdownCellContentSpecBuilder implements com.grafana.foundation.cog.Builder<MarkdownCellContentSpec> {
    protected final MarkdownCellContentSpec internal;
    
    public MarkdownCellContentSpecBuilder() {
        this.internal = new MarkdownCellContentSpec();
    }
    public MarkdownCellContentSpecBuilder text(String text) {
        this.internal.text = text;
        return this;
    }
    public MarkdownCellContentSpec build() {
        return this.internal;
    }
}
