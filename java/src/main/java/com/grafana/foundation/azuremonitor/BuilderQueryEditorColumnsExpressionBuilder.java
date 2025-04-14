// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;

import java.util.List;

public class BuilderQueryEditorColumnsExpressionBuilder implements com.grafana.foundation.cog.Builder<BuilderQueryEditorColumnsExpression> {
    protected final BuilderQueryEditorColumnsExpression internal;
    
    public BuilderQueryEditorColumnsExpressionBuilder() {
        this.internal = new BuilderQueryEditorColumnsExpression();
    }
    public BuilderQueryEditorColumnsExpressionBuilder columns(List<String> columns) {
        this.internal.columns = columns;
        return this;
    }
    
    public BuilderQueryEditorColumnsExpressionBuilder type(BuilderQueryEditorExpressionType type) {
        this.internal.type = type;
        return this;
    }
    public BuilderQueryEditorColumnsExpression build() {
        return this.internal;
    }
}
