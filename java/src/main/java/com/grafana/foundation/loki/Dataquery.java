// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.loki;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.grafana.foundation.dashboard.DataSourceRef;

public class Dataquery implements com.grafana.foundation.cog.variants.Dataquery {
    // The LogQL query.
    @JsonProperty("expr")
    public String expr;
    // Used to override the name of the series.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("legendFormat")
    public String legendFormat;
    // Used to limit the number of log rows returned.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("maxLines")
    public Long maxLines;
    // @deprecated, now use step.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("resolution")
    public Long resolution;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("editorMode")
    public QueryEditorMode editorMode;
    // @deprecated, now use queryType.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("range")
    public Boolean range;
    // @deprecated, now use queryType.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("instant")
    public Boolean instant;
    // Used to set step value for range queries.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("step")
    public String step;
    // A unique identifier for the query within the list of targets.
    // In server side expressions, the refId is used as a variable name to identify results.
    // By default, the UI will assign A->Z; however setting meaningful names may be useful.
    @JsonProperty("refId")
    public String refId;
    // true if query is disabled (ie should not be returned to the dashboard)
    // Note this does not always imply that the query should not be executed since
    // the results from a hidden query may be used as the input to other queries (SSE etc)
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("hide")
    public Boolean hide;
    // Specify the query flavor
    // TODO make this required and give it a default
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("queryType")
    public String queryType;
    // For mixed data sources the selected datasource is on the query level.
    // For non mixed scenarios this is undefined.
    // TODO find a better way to do this ^ that's friendly to schema
    // TODO this shouldn't be unknown but DataSourceRef | null
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("datasource")
    public DataSourceRef datasource;
    public String dataqueryName() {
        return "loki";
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<com.grafana.foundation.cog.variants.Dataquery> {
        protected final Dataquery internal;
        
        public Builder() {
            this.internal = new Dataquery();
        }
    public Builder expr(String expr) {
    this.internal.expr = expr;
        return this;
    }
    
    public Builder legendFormat(String legendFormat) {
    this.internal.legendFormat = legendFormat;
        return this;
    }
    
    public Builder maxLines(Long maxLines) {
    this.internal.maxLines = maxLines;
        return this;
    }
    
    public Builder resolution(Long resolution) {
    this.internal.resolution = resolution;
        return this;
    }
    
    public Builder editorMode(QueryEditorMode editorMode) {
    this.internal.editorMode = editorMode;
        return this;
    }
    
    public Builder range(Boolean range) {
    this.internal.range = range;
        return this;
    }
    
    public Builder instant(Boolean instant) {
    this.internal.instant = instant;
        return this;
    }
    
    public Builder step(String step) {
    this.internal.step = step;
        return this;
    }
    
    public Builder refId(String refId) {
    this.internal.refId = refId;
        return this;
    }
    
    public Builder hide(Boolean hide) {
    this.internal.hide = hide;
        return this;
    }
    
    public Builder queryType(String queryType) {
    this.internal.queryType = queryType;
        return this;
    }
    
    public Builder datasource(DataSourceRef datasource) {
    this.internal.datasource = datasource;
        return this;
    }
    public Dataquery build() {
            return this.internal;
        }
    }
}
