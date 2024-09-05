// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.tempo;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;
import java.util.List;

public class TempoQuery implements com.grafana.foundation.cog.variants.Dataquery {
    // A unique identifier for the query within the list of targets.
    // In server side expressions, the refId is used as a variable name to identify results.
    // By default, the UI will assign A->Z; however setting meaningful names may be useful.
    @JsonProperty("refId")
    public String refId;
    // true if query is disabled (ie should not be returned to the dashboard)
    // Note this does not always imply that the query should not be executed since
    // the results from a hidden query may be used as the input to other queries (SSE etc)
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("hide")
    public Boolean hide;
    // Specify the query flavor
    // TODO make this required and give it a default
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("queryType")
    public String queryType;
    // TraceQL query or trace ID
    @JsonProperty("query")
    public String query;
    // @deprecated Logfmt query to filter traces by their tags. Example: http.status_code=200 error=true
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("search")
    public String search;
    // @deprecated Query traces by service name
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("serviceName")
    public String serviceName;
    // @deprecated Query traces by span name
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("spanName")
    public String spanName;
    // @deprecated Define the minimum duration to select traces. Use duration format, for example: 1.2s, 100ms
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("minDuration")
    public String minDuration;
    // @deprecated Define the maximum duration to select traces. Use duration format, for example: 1.2s, 100ms
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("maxDuration")
    public String maxDuration;
    // Filters to be included in a PromQL query to select data for the service graph. Example: {client="app",service="app"}
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("serviceMapQuery")
    public String serviceMapQuery;
    // Use service.namespace in addition to service.name to uniquely identify a service.
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("serviceMapIncludeNamespace")
    public Boolean serviceMapIncludeNamespace;
    // Defines the maximum number of traces that are returned from Tempo
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("limit")
    public Long limit;
    // For mixed data sources the selected datasource is on the query level.
    // For non mixed scenarios this is undefined.
    // TODO find a better way to do this ^ that's friendly to schema
    // TODO this shouldn't be unknown but DataSourceRef | null
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("datasource")
    public Object datasource;
    @JsonProperty("filters")
    public List<TraceqlFilter> filters;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<TempoQuery> {
        private final TempoQuery internal;
        
        public Builder() {
            this.internal = new TempoQuery();
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
    
    public Builder query(String query) {
    this.internal.query = query;
        return this;
    }
    
    public Builder search(String search) {
    this.internal.search = search;
        return this;
    }
    
    public Builder serviceName(String serviceName) {
    this.internal.serviceName = serviceName;
        return this;
    }
    
    public Builder spanName(String spanName) {
    this.internal.spanName = spanName;
        return this;
    }
    
    public Builder minDuration(String minDuration) {
    this.internal.minDuration = minDuration;
        return this;
    }
    
    public Builder maxDuration(String maxDuration) {
    this.internal.maxDuration = maxDuration;
        return this;
    }
    
    public Builder serviceMapQuery(String serviceMapQuery) {
    this.internal.serviceMapQuery = serviceMapQuery;
        return this;
    }
    
    public Builder serviceMapIncludeNamespace(Boolean serviceMapIncludeNamespace) {
    this.internal.serviceMapIncludeNamespace = serviceMapIncludeNamespace;
        return this;
    }
    
    public Builder limit(Long limit) {
    this.internal.limit = limit;
        return this;
    }
    
    public Builder datasource(Object datasource) {
    this.internal.datasource = datasource;
        return this;
    }
    
    public Builder filters(com.grafana.foundation.cog.Builder<List<TraceqlFilter>> filters) {
    this.internal.filters = filters.build();
        return this;
    }
    public TempoQuery build() {
            return this.internal;
        }
    }
}
