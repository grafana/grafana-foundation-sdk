// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.testdata;

import java.util.List;

public class ResultAssertionsBuilder implements com.grafana.foundation.cog.Builder<ResultAssertions> {
    protected final ResultAssertions internal;
    
    public ResultAssertionsBuilder() {
        this.internal = new ResultAssertions();
    }
    public ResultAssertionsBuilder maxFrames(Long maxFrames) {
        this.internal.maxFrames = maxFrames;
        return this;
    }
    
    public ResultAssertionsBuilder type(ResultAssertionsType type) {
        this.internal.type = type;
        return this;
    }
    
    public ResultAssertionsBuilder typeVersion(List<Long> typeVersion) {
        this.internal.typeVersion = typeVersion;
        return this;
    }
    public ResultAssertions build() {
        return this.internal;
    }
}
