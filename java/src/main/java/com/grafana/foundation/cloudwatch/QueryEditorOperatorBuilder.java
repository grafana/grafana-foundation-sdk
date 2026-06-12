// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.cloudwatch;


/**
 * TS type is QueryEditorOperator<T extends QueryEditorOperatorValueType>, extended in veneer
 */
public class QueryEditorOperatorBuilder implements com.grafana.foundation.cog.Builder<QueryEditorOperator> {
    protected final QueryEditorOperator internal;
    
    public QueryEditorOperatorBuilder() {
        this.internal = new QueryEditorOperator();
    }
    public QueryEditorOperatorBuilder name(String name) {
        this.internal.name = name;
        return this;
    }
    
    public QueryEditorOperatorBuilder value(QueryEditorOperatorValueType value) {
        this.internal.value = value;
        return this;
    }
    public QueryEditorOperator build() {
        return this.internal;
    }
}
