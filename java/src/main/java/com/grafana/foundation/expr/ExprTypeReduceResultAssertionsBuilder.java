// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.expr;

import java.util.List;

public class ExprTypeReduceResultAssertionsBuilder implements com.grafana.foundation.cog.Builder<ExprTypeReduceResultAssertions> {
    protected final ExprTypeReduceResultAssertions internal;
    
    public ExprTypeReduceResultAssertionsBuilder() {
        this.internal = new ExprTypeReduceResultAssertions();
    }
    public ExprTypeReduceResultAssertionsBuilder maxFrames(Long maxFrames) {
        this.internal.maxFrames = maxFrames;
        return this;
    }
    
    public ExprTypeReduceResultAssertionsBuilder type(ExprTypeReduceResultAssertionsType type) {
        this.internal.type = type;
        return this;
    }
    
    public ExprTypeReduceResultAssertionsBuilder typeVersion(List<Long> typeVersion) {
        this.internal.typeVersion = typeVersion;
        return this;
    }
    public ExprTypeReduceResultAssertions build() {
        return this.internal;
    }
}
