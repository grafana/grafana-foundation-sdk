// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.tempo;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonSetter;
import com.fasterxml.jackson.annotation.Nulls;
import java.util.List;
import com.grafana.foundation.dashboard.DataSourceRef;

public class TempoQuery implements com.grafana.foundation.cog.variants.Dataquery {
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
    // TraceQL query or trace ID
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("query")
    public String query;
    // @deprecated Logfmt query to filter traces by their tags. Example: http.status_code=200 error=true
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("search")
    public String search;
    // @deprecated Query traces by service name
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("serviceName")
    public String serviceName;
    // @deprecated Query traces by span name
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("spanName")
    public String spanName;
    // @deprecated Define the minimum duration to select traces. Use duration format, for example: 1.2s, 100ms
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("minDuration")
    public String minDuration;
    // @deprecated Define the maximum duration to select traces. Use duration format, for example: 1.2s, 100ms
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("maxDuration")
    public String maxDuration;
    // Filters to be included in a PromQL query to select data for the service graph. Example: {client="app",service="app"}. Providing multiple values will produce union of results for each filter, using PromQL OR operator internally.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("serviceMapQuery")
    public StringOrArrayOfString serviceMapQuery;
    // Use service.namespace in addition to service.name to uniquely identify a service.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("serviceMapIncludeNamespace")
    public Boolean serviceMapIncludeNamespace;
    // Defines the maximum number of traces that are returned from Tempo
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("limit")
    public Long limit;
    // Defines the maximum number of spans per spanset that are returned from Tempo
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("spss")
    public Long spss;
    @JsonSetter(nulls = Nulls.AS_EMPTY)
    @JsonProperty("filters")
    public List<TraceqlFilter> filters;
    // Filters that are used to query the metrics summary
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("groupBy")
    public List<TraceqlFilter> groupBy;
    // For mixed data sources the selected datasource is on the query level.
    // For non mixed scenarios this is undefined.
    // TODO find a better way to do this ^ that's friendly to schema
    // TODO this shouldn't be unknown but DataSourceRef | null
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("datasource")
    public DataSourceRef datasource;
    // The type of the table that is used to display the search results
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("tableType")
    public SearchTableType tableType;
    public String dataqueryName() {
        return "tempo";
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<TempoQuery> {
        protected final TempoQuery internal;
        
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
    
    public Builder serviceMapQuery(StringOrArrayOfString serviceMapQuery) {
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
    
    public Builder spss(Long spss) {
    this.internal.spss = spss;
        return this;
    }
    
    public Builder filters(com.grafana.foundation.cog.Builder<List<TraceqlFilter>> filters) {
    this.internal.filters = filters.build();
        return this;
    }
    
    public Builder groupBy(com.grafana.foundation.cog.Builder<List<TraceqlFilter>> groupBy) {
    this.internal.groupBy = groupBy.build();
        return this;
    }
    
    public Builder datasource(DataSourceRef datasource) {
    this.internal.datasource = datasource;
        return this;
    }
    
    public Builder tableType(SearchTableType tableType) {
    this.internal.tableType = tableType;
        return this;
    }
    public TempoQuery build() {
            return this.internal;
        }
    }
}
