// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.expr;

import java.util.List;

public class ExprTypeSqlResultAssertionsBuilder implements com.grafana.foundation.cog.Builder<ExprTypeSqlResultAssertions> {
    protected final ExprTypeSqlResultAssertions internal;
    
    public ExprTypeSqlResultAssertionsBuilder() {
        this.internal = new ExprTypeSqlResultAssertions();
    }
    public ExprTypeSqlResultAssertionsBuilder maxFrames(Long maxFrames) {
        this.internal.maxFrames = maxFrames;
        return this;
    }
    
    public ExprTypeSqlResultAssertionsBuilder type(ExprTypeSqlResultAssertionsType type) {
        this.internal.type = type;
        return this;
    }
    
    public ExprTypeSqlResultAssertionsBuilder typeVersion(List<Long> typeVersion) {
        this.internal.typeVersion = typeVersion;
        return this;
    }
    public ExprTypeSqlResultAssertions build() {
        return this.internal;
    }
}
