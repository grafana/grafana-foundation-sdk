// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;


public class CodeCellContentKindBuilder implements com.grafana.foundation.cog.Builder<CodeCellContentKind> {
    protected final CodeCellContentKind internal;
    
    public CodeCellContentKindBuilder() {
        this.internal = new CodeCellContentKind();
        this.internal.kind = "Code";
    }
    public CodeCellContentKindBuilder spec(com.grafana.foundation.cog.Builder<CodeCellContentSpec> spec) {
    CodeCellContentSpec specResource = spec.build();
        this.internal.spec = specResource;
        return this;
    }
    public CodeCellContentKind build() {
        return this.internal;
    }
}
