// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.prometheus;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.grafana.foundation.dashboard.DataSourceRef;

public class Dataquery implements com.grafana.foundation.cog.variants.Dataquery {
    // The actual expression/query that will be evaluated by Prometheus
    @JsonProperty("expr")
    public String expr;
    // Returns only the latest value that Prometheus has scraped for the requested time series
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("instant")
    public Boolean instant;
    // Returns a Range vector, comprised of a set of time series containing a range of data points over time for each time series
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("range")
    public Boolean range;
    // Execute an additional query to identify interesting raw samples relevant for the given expr
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("exemplar")
    public Boolean exemplar;
    // Specifies which editor is being used to prepare the query. It can be "code" or "builder"
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("editorMode")
    public QueryEditorMode editorMode;
    // Query format to determine how to display data points in panel. It can be "time_series", "table", "heatmap"
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("format")
    public PromQueryFormat format;
    // Series name override or template. Ex. {{hostname}} will be replaced with label value for hostname
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("legendFormat")
    public String legendFormat;
    // @deprecated Used to specify how many times to divide max data points by. We use max data points under query options
    // See https://github.com/grafana/grafana/issues/48081
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("intervalFactor")
    public Double intervalFactor;
    // A unique identifier for the query within the list of targets.
    // In server side expressions, the refId is used as a variable name to identify results.
    // By default, the UI will assign A->Z; however setting meaningful names may be useful.
    @JsonProperty("refId")
    public String refId;
    // If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
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
    // An additional lower limit for the step parameter of the Prometheus query and for the
    // `$__interval` and `$__rate_interval` variables.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("interval")
    public String interval;
    public String dataqueryName() {
        return "prometheus";
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
    
    public Builder instant() {
    this.internal.instant = true;
    this.internal.range = false;
        return this;
    }
    
    public Builder range() {
    this.internal.range = true;
    this.internal.instant = false;
        return this;
    }
    
    public Builder exemplar(Boolean exemplar) {
    this.internal.exemplar = exemplar;
        return this;
    }
    
    public Builder editorMode(QueryEditorMode editorMode) {
    this.internal.editorMode = editorMode;
        return this;
    }
    
    public Builder format(PromQueryFormat format) {
    this.internal.format = format;
        return this;
    }
    
    public Builder legendFormat(String legendFormat) {
    this.internal.legendFormat = legendFormat;
        return this;
    }
    
    public Builder intervalFactor(Double intervalFactor) {
    this.internal.intervalFactor = intervalFactor;
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
    
    public Builder interval(String interval) {
    this.internal.interval = interval;
        return this;
    }
    
    public Builder rangeAndInstant() {
    this.internal.range = true;
    this.internal.instant = true;
        return this;
    }
    public Dataquery build() {
            return this.internal;
        }
    }
}
