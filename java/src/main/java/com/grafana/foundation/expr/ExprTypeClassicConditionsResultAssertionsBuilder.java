// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.expr;

import java.util.List;

public class ExprTypeClassicConditionsResultAssertionsBuilder implements com.grafana.foundation.cog.Builder<ExprTypeClassicConditionsResultAssertions> {
    protected final ExprTypeClassicConditionsResultAssertions internal;
    
    public ExprTypeClassicConditionsResultAssertionsBuilder() {
        this.internal = new ExprTypeClassicConditionsResultAssertions();
    }
    public ExprTypeClassicConditionsResultAssertionsBuilder maxFrames(Long maxFrames) {
        this.internal.maxFrames = maxFrames;
        return this;
    }
    
    public ExprTypeClassicConditionsResultAssertionsBuilder type(ExprTypeClassicConditionsResultAssertionsType type) {
        this.internal.type = type;
        return this;
    }
    
    public ExprTypeClassicConditionsResultAssertionsBuilder typeVersion(List<Long> typeVersion) {
        this.internal.typeVersion = typeVersion;
        return this;
    }
    public ExprTypeClassicConditionsResultAssertions build() {
        return this.internal;
    }
}
