// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.bigquery;


public class QueryEditorPropertyBuilder implements com.grafana.foundation.cog.Builder<QueryEditorProperty> {
    protected final QueryEditorProperty internal;
    
    public QueryEditorPropertyBuilder() {
        this.internal = new QueryEditorProperty();
    }
    public QueryEditorPropertyBuilder type(QueryEditorPropertyType type) {
        this.internal.type = type;
        return this;
    }
    
    public QueryEditorPropertyBuilder name(String name) {
        this.internal.name = name;
        return this;
    }
    public QueryEditorProperty build() {
        return this.internal;
    }
}
