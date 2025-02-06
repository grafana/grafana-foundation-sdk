// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.expr;

import java.util.List;

public class ExprTypeResampleResultAssertionsBuilder implements com.grafana.foundation.cog.Builder<ExprTypeResampleResultAssertions> {
    protected final ExprTypeResampleResultAssertions internal;
    
    public ExprTypeResampleResultAssertionsBuilder() {
        this.internal = new ExprTypeResampleResultAssertions();
    }
    public ExprTypeResampleResultAssertionsBuilder maxFrames(Long maxFrames) {
        this.internal.maxFrames = maxFrames;
        return this;
    }
    
    public ExprTypeResampleResultAssertionsBuilder type(ExprTypeResampleResultAssertionsType type) {
        this.internal.type = type;
        return this;
    }
    
    public ExprTypeResampleResultAssertionsBuilder typeVersion(List<Long> typeVersion) {
        this.internal.typeVersion = typeVersion;
        return this;
    }
    public ExprTypeResampleResultAssertions build() {
        return this.internal;
    }
}
