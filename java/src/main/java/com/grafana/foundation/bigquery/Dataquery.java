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
    public String dataqueryName() {
        return "grafana-bigquery-datasource";
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
    public Builder dataset(String dataset) {
    this.internal.dataset = dataset;
        return this;
    }
    
    public Builder table(String table) {
    this.internal.table = table;
        return this;
    }
    
    public Builder project(String project) {
    this.internal.project = project;
        return this;
    }
    
    public Builder format(QueryFormat format) {
    this.internal.format = format;
        return this;
    }
    
    public Builder rawQuery(Boolean rawQuery) {
    this.internal.rawQuery = rawQuery;
        return this;
    }
    
    public Builder rawSql(String rawSql) {
    this.internal.rawSql = rawSql;
        return this;
    }
    
    public Builder location(String location) {
    this.internal.location = location;
        return this;
    }
    
    public Builder partitioned(Boolean partitioned) {
    this.internal.partitioned = partitioned;
        return this;
    }
    
    public Builder partitionedField(String partitionedField) {
    this.internal.partitionedField = partitionedField;
        return this;
    }
    
    public Builder convertToUTC(Boolean convertToUTC) {
    this.internal.convertToUTC = convertToUTC;
        return this;
    }
    
    public Builder sharded(Boolean sharded) {
    this.internal.sharded = sharded;
        return this;
    }
    
    public Builder queryPriority(QueryPriority queryPriority) {
    this.internal.queryPriority = queryPriority;
        return this;
    }
    
    public Builder timeShift(String timeShift) {
    this.internal.timeShift = timeShift;
        return this;
    }
    
    public Builder editorMode(EditorMode editorMode) {
    this.internal.editorMode = editorMode;
        return this;
    }
    
    public Builder sql(com.grafana.foundation.cog.Builder<SQLExpression> sql) {
    this.internal.sql = sql.build();
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
