// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.expr;

import java.util.List;

public class ExprTypeMathResultAssertionsBuilder implements com.grafana.foundation.cog.Builder<ExprTypeMathResultAssertions> {
    protected final ExprTypeMathResultAssertions internal;
    
    public ExprTypeMathResultAssertionsBuilder() {
        this.internal = new ExprTypeMathResultAssertions();
    }
    public ExprTypeMathResultAssertionsBuilder maxFrames(Long maxFrames) {
        this.internal.maxFrames = maxFrames;
        return this;
    }
    
    public ExprTypeMathResultAssertionsBuilder type(ExprTypeMathResultAssertionsType type) {
        this.internal.type = type;
        return this;
    }
    
    public ExprTypeMathResultAssertionsBuilder typeVersion(List<Long> typeVersion) {
        this.internal.typeVersion = typeVersion;
        return this;
    }
    public ExprTypeMathResultAssertions build() {
        return this.internal;
    }
}
