// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.bigquery;

import java.util.List;
import java.util.LinkedList;

public class QueryEditorFunctionExpressionBuilder implements com.grafana.foundation.cog.Builder<QueryEditorFunctionExpression> {
    protected final QueryEditorFunctionExpression internal;
    
    public QueryEditorFunctionExpressionBuilder() {
        this.internal = new QueryEditorFunctionExpression();
    }
    public QueryEditorFunctionExpressionBuilder name(String name) {
        this.internal.name = name;
        return this;
    }
    
    public QueryEditorFunctionExpressionBuilder parameters(List<com.grafana.foundation.cog.Builder<QueryEditorFunctionParameterExpression>> parameters) {
        List<QueryEditorFunctionParameterExpression> parametersResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<QueryEditorFunctionParameterExpression> r1 : parameters) {
                QueryEditorFunctionParameterExpression parametersDepth1 = r1.build();
                parametersResources.add(parametersDepth1); 
        }
        this.internal.parameters = parametersResources;
        return this;
    }
    public QueryEditorFunctionExpression build() {
        return this.internal;
    }
}
