// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.testdata;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.List;

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
    @JsonInclude(JsonInclude.Include.NON_NULL)
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
    public Object datasource;
    public Dataquery() {
        this.scenarioId = TestDataQueryType.RANDOM_WALK;
    }
    public Dataquery(String alias,TestDataQueryType scenarioId,String stringInput,StreamingQuery stream,PulseWaveQuery pulseWave,SimulationQuery sim,List<CSVWave> csvWave,String labels,Long lines,Boolean levelColumn,String channel,NodesQuery nodes,String csvFileName,String csvContent,String rawFrameContent,Integer seriesCount,USAQuery usa,DataqueryErrorType errorType,Integer spanCount,List<List<StringOrInt64>> points,Double dropPercent,Boolean flamegraphDiff,String refId,Boolean hide,String queryType,Object datasource) {
        this.alias = alias;
        this.scenarioId = scenarioId;
        this.stringInput = stringInput;
        this.stream = stream;
        this.pulseWave = pulseWave;
        this.sim = sim;
        this.csvWave = csvWave;
        this.labels = labels;
        this.lines = lines;
        this.levelColumn = levelColumn;
        this.channel = channel;
        this.nodes = nodes;
        this.csvFileName = csvFileName;
        this.csvContent = csvContent;
        this.rawFrameContent = rawFrameContent;
        this.seriesCount = seriesCount;
        this.usa = usa;
        this.errorType = errorType;
        this.spanCount = spanCount;
        this.points = points;
        this.dropPercent = dropPercent;
        this.flamegraphDiff = flamegraphDiff;
        this.refId = refId;
        this.hide = hide;
        this.queryType = queryType;
        this.datasource = datasource;
    }
    public String dataqueryName() {
        return "testdata";
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
