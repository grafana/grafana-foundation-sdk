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
    // The type of the table that is used to display the search results
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("tableType")
    public SearchTableType tableType;
    // For metric queries, the step size to use
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("step")
    public String step;
    // For metric queries, how many exemplars to request, 0 means no exemplars
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("exemplars")
    public Long exemplars;
    // For mixed data sources the selected datasource is on the query level.
    // For non mixed scenarios this is undefined.
    // TODO find a better way to do this ^ that's friendly to schema
    // TODO this shouldn't be unknown but DataSourceRef | null
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("datasource")
    public DataSourceRef datasource;
    // For metric queries, whether to run instant or range queries
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("metricsQueryType")
    public MetricsQueryType metricsQueryType;
    public TempoQuery() {
    }
    
    public TempoQuery(String refId,Boolean hide,String queryType,String query,String search,String serviceName,String spanName,String minDuration,String maxDuration,StringOrArrayOfString serviceMapQuery,Boolean serviceMapIncludeNamespace,Long limit,Long spss,List<TraceqlFilter> filters,List<TraceqlFilter> groupBy,SearchTableType tableType,String step,Long exemplars,DataSourceRef datasource,MetricsQueryType metricsQueryType) {
        this.refId = refId;
        this.hide = hide;
        this.queryType = queryType;
        this.query = query;
        this.search = search;
        this.serviceName = serviceName;
        this.spanName = spanName;
        this.minDuration = minDuration;
        this.maxDuration = maxDuration;
        this.serviceMapQuery = serviceMapQuery;
        this.serviceMapIncludeNamespace = serviceMapIncludeNamespace;
        this.limit = limit;
        this.spss = spss;
        this.filters = filters;
        this.groupBy = groupBy;
        this.tableType = tableType;
        this.step = step;
        this.exemplars = exemplars;
        this.datasource = datasource;
        this.metricsQueryType = metricsQueryType;
    }
    public String dataqueryName() {
        return "tempo";
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
