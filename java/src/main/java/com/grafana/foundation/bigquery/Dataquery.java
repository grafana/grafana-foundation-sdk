// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.bigquery;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.grafana.foundation.dashboard.DataSourceRef;

// Manually converted from https://github.com/grafana/google-bigquery-datasource/blob/18680e42ba557791d109c7c540c2c3f2647592f0/src/types.ts#L75
public class Dataquery implements com.grafana.foundation.cog.variants.Dataquery {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("dataset")
    public String dataset;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("table")
    public String table;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("project")
    public String project;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("format")
    public QueryFormat format;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("rawQuery")
    public Boolean rawQuery;
    @JsonProperty("rawSql")
    public String rawSql;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("location")
    public String location;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("partitioned")
    public Boolean partitioned;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("partitionedField")
    public String partitionedField;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("convertToUTC")
    public Boolean convertToUTC;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("sharded")
    public Boolean sharded;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("queryPriority")
    public QueryPriority queryPriority;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("timeShift")
    public String timeShift;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("editorMode")
    public EditorMode editorMode;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("sql")
    public SQLExpression sql;
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
    public Dataquery() {
        this.format = QueryFormat.TIMESERIES;
        this.rawSql = "";
        this.refId = "";
    }
    public Dataquery(String dataset,String table,String project,QueryFormat format,Boolean rawQuery,String rawSql,String location,Boolean partitioned,String partitionedField,Boolean convertToUTC,Boolean sharded,QueryPriority queryPriority,String timeShift,EditorMode editorMode,SQLExpression sql,String refId,Boolean hide,String queryType,DataSourceRef datasource) {
        this.dataset = dataset;
        this.table = table;
        this.project = project;
        this.format = format;
        this.rawQuery = rawQuery;
        this.rawSql = rawSql;
        this.location = location;
        this.partitioned = partitioned;
        this.partitionedField = partitionedField;
        this.convertToUTC = convertToUTC;
        this.sharded = sharded;
        this.queryPriority = queryPriority;
        this.timeShift = timeShift;
        this.editorMode = editorMode;
        this.sql = sql;
        this.refId = refId;
        this.hide = hide;
        this.queryType = queryType;
        this.datasource = datasource;
    }
    public String dataqueryName() {
        return "grafana-bigquery-datasource";
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
