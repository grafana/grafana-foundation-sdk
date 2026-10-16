// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import java.util.List;

public class CodeCellContentSpecBuilder implements com.grafana.foundation.cog.Builder<CodeCellContentSpec> {
    protected final CodeCellContentSpec internal;
    
    public CodeCellContentSpecBuilder() {
        this.internal = new CodeCellContentSpec();
    }
    public CodeCellContentSpecBuilder language(String language) {
        this.internal.language = language;
        return this;
    }
    
    public CodeCellContentSpecBuilder code(String code) {
        this.internal.code = code;
        return this;
    }
    
    public CodeCellContentSpecBuilder highlight(List<Long> highlight) {
        this.internal.highlight = highlight;
        return this;
    }
    
    public CodeCellContentSpecBuilder annotation(String annotation) {
        this.internal.annotation = annotation;
        return this;
    }
    public CodeCellContentSpec build() {
        return this.internal;
    }
}
