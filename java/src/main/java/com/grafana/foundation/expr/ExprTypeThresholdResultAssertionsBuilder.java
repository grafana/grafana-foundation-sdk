// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.expr;

import java.util.List;

public class ExprTypeThresholdResultAssertionsBuilder implements com.grafana.foundation.cog.Builder<ExprTypeThresholdResultAssertions> {
    protected final ExprTypeThresholdResultAssertions internal;
    
    public ExprTypeThresholdResultAssertionsBuilder() {
        this.internal = new ExprTypeThresholdResultAssertions();
    }
    public ExprTypeThresholdResultAssertionsBuilder maxFrames(Long maxFrames) {
        this.internal.maxFrames = maxFrames;
        return this;
    }
    
    public ExprTypeThresholdResultAssertionsBuilder type(ExprTypeThresholdResultAssertionsType type) {
        this.internal.type = type;
        return this;
    }
    
    public ExprTypeThresholdResultAssertionsBuilder typeVersion(List<Long> typeVersion) {
        this.internal.typeVersion = typeVersion;
        return this;
    }
    public ExprTypeThresholdResultAssertions build() {
        return this.internal;
    }
}
