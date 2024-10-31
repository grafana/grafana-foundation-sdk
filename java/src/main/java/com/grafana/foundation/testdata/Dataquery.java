// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.testdata;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.List;
import com.grafana.foundation.dashboard.DataSourceRef;

public class Dataquery implements com.grafana.foundation.cog.variants.Dataquery {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("alias")
    public String alias;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("scenarioId")
    public TestDataQueryType scenarioId;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("stringInput")
    public String stringInput;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("stream")
    public StreamingQuery stream;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("pulseWave")
    public PulseWaveQuery pulseWave;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("sim")
    public SimulationQuery sim;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("csvWave")
    public List<CSVWave> csvWave;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("labels")
    public String labels;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("lines")
    public Long lines;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("levelColumn")
    public Boolean levelColumn;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("channel")
    public String channel;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("nodes")
    public NodesQuery nodes;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("csvFileName")
    public String csvFileName;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("csvContent")
    public String csvContent;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("rawFrameContent")
    public String rawFrameContent;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("seriesCount")
    public Integer seriesCount;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("usa")
    public USAQuery usa;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("errorType")
    public DataqueryErrorType errorType;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("spanCount")
    public Integer spanCount;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("points")
    public List<List<StringOrInt64>> points;
    // Drop percentage (the chance we will lose a point 0-100)
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("dropPercent")
    public Double dropPercent;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("flamegraphDiff")
    public Boolean flamegraphDiff;
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
        return "testdata";
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<com.grafana.foundation.cog.variants.Dataquery> {
        protected final Dataquery internal;
        
        public Builder() {
            this.internal = new Dataquery();
        this.scenarioId(TestDataQueryType.RANDOM_WALK);
        }
    public Builder alias(String alias) {
    this.internal.alias = alias;
        return this;
    }
    
    public Builder scenarioId(TestDataQueryType scenarioId) {
    this.internal.scenarioId = scenarioId;
        return this;
    }
    
    public Builder stringInput(String stringInput) {
    this.internal.stringInput = stringInput;
        return this;
    }
    
    public Builder stream(com.grafana.foundation.cog.Builder<StreamingQuery> stream) {
    this.internal.stream = stream.build();
        return this;
    }
    
    public Builder pulseWave(com.grafana.foundation.cog.Builder<PulseWaveQuery> pulseWave) {
    this.internal.pulseWave = pulseWave.build();
        return this;
    }
    
    public Builder sim(com.grafana.foundation.cog.Builder<SimulationQuery> sim) {
    this.internal.sim = sim.build();
        return this;
    }
    
    public Builder csvWave(com.grafana.foundation.cog.Builder<List<CSVWave>> csvWave) {
    this.internal.csvWave = csvWave.build();
        return this;
    }
    
    public Builder labels(String labels) {
    this.internal.labels = labels;
        return this;
    }
    
    public Builder lines(Long lines) {
    this.internal.lines = lines;
        return this;
    }
    
    public Builder levelColumn(Boolean levelColumn) {
    this.internal.levelColumn = levelColumn;
        return this;
    }
    
    public Builder channel(String channel) {
    this.internal.channel = channel;
        return this;
    }
    
    public Builder nodes(com.grafana.foundation.cog.Builder<NodesQuery> nodes) {
    this.internal.nodes = nodes.build();
        return this;
    }
    
    public Builder csvFileName(String csvFileName) {
    this.internal.csvFileName = csvFileName;
        return this;
    }
    
    public Builder csvContent(String csvContent) {
    this.internal.csvContent = csvContent;
        return this;
    }
    
    public Builder rawFrameContent(String rawFrameContent) {
    this.internal.rawFrameContent = rawFrameContent;
        return this;
    }
    
    public Builder seriesCount(Integer seriesCount) {
    this.internal.seriesCount = seriesCount;
        return this;
    }
    
    public Builder usa(com.grafana.foundation.cog.Builder<USAQuery> usa) {
    this.internal.usa = usa.build();
        return this;
    }
    
    public Builder errorType(DataqueryErrorType errorType) {
    this.internal.errorType = errorType;
        return this;
    }
    
    public Builder spanCount(Integer spanCount) {
    this.internal.spanCount = spanCount;
        return this;
    }
    
    public Builder points(List<List<StringOrInt64>> points) {
    this.internal.points = points;
        return this;
    }
    
    public Builder dropPercent(Double dropPercent) {
    this.internal.dropPercent = dropPercent;
        return this;
    }
    
    public Builder flamegraphDiff(Boolean flamegraphDiff) {
    this.internal.flamegraphDiff = flamegraphDiff;
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
