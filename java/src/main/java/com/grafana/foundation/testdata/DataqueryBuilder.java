// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.testdata;

import java.util.List;
import java.util.LinkedList;
import com.grafana.foundation.dashboard.DataSourceRef;

public class DataqueryBuilder implements com.grafana.foundation.cog.Builder<com.grafana.foundation.cog.variants.Dataquery> {
    protected final Dataquery internal;
    
    public DataqueryBuilder() {
        this.internal = new Dataquery();
    }
    public DataqueryBuilder alias(String alias) {
        this.internal.alias = alias;
        return this;
    }
    
    public DataqueryBuilder channel(String channel) {
        this.internal.channel = channel;
        return this;
    }
    
    public DataqueryBuilder csvContent(String csvContent) {
        this.internal.csvContent = csvContent;
        return this;
    }
    
    public DataqueryBuilder csvFileName(String csvFileName) {
        this.internal.csvFileName = csvFileName;
        return this;
    }
    
    public DataqueryBuilder csvWave(List<com.grafana.foundation.cog.Builder<CSVWave>> csvWave) {
        List<CSVWave> csvWaveResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<CSVWave> r1 : csvWave) {
                CSVWave csvWaveDepth1 = r1.build();
                csvWaveResources.add(csvWaveDepth1); 
        }
        this.internal.csvWave = csvWaveResources;
        return this;
    }
    
    public DataqueryBuilder datasource(DataSourceRef datasource) {
        this.internal.datasource = datasource;
        return this;
    }
    
    public DataqueryBuilder dropPercent(Double dropPercent) {
        this.internal.dropPercent = dropPercent;
        return this;
    }
    
    public DataqueryBuilder errorSource(DataqueryErrorSource errorSource) {
        this.internal.errorSource = errorSource;
        return this;
    }
    
    public DataqueryBuilder errorType(DataqueryErrorType errorType) {
        this.internal.errorType = errorType;
        return this;
    }
    
    public DataqueryBuilder flamegraphDiff(Boolean flamegraphDiff) {
        this.internal.flamegraphDiff = flamegraphDiff;
        return this;
    }
    
    public DataqueryBuilder hide(Boolean hide) {
        this.internal.hide = hide;
        return this;
    }
    
    public DataqueryBuilder intervalMs(Double intervalMs) {
        this.internal.intervalMs = intervalMs;
        return this;
    }
    
    public DataqueryBuilder labels(String labels) {
        this.internal.labels = labels;
        return this;
    }
    
    public DataqueryBuilder levelColumn(Boolean levelColumn) {
        this.internal.levelColumn = levelColumn;
        return this;
    }
    
    public DataqueryBuilder lines(Long lines) {
        this.internal.lines = lines;
        return this;
    }
    
    public DataqueryBuilder max(Double max) {
        this.internal.max = max;
        return this;
    }
    
    public DataqueryBuilder maxDataPoints(Long maxDataPoints) {
        this.internal.maxDataPoints = maxDataPoints;
        return this;
    }
    
    public DataqueryBuilder min(Double min) {
        this.internal.min = min;
        return this;
    }
    
    public DataqueryBuilder nodes(com.grafana.foundation.cog.Builder<NodesQuery> nodes) {
    NodesQuery nodesResource = nodes.build();
        this.internal.nodes = nodesResource;
        return this;
    }
    
    public DataqueryBuilder noise(Double noise) {
        this.internal.noise = noise;
        return this;
    }
    
    public DataqueryBuilder points(List<List<Object>> points) {
        this.internal.points = points;
        return this;
    }
    
    public DataqueryBuilder pulseWave(com.grafana.foundation.cog.Builder<PulseWaveQuery> pulseWave) {
    PulseWaveQuery pulseWaveResource = pulseWave.build();
        this.internal.pulseWave = pulseWaveResource;
        return this;
    }
    
    public DataqueryBuilder queryType(String queryType) {
        this.internal.queryType = queryType;
        return this;
    }
    
    public DataqueryBuilder rawFrameContent(String rawFrameContent) {
        this.internal.rawFrameContent = rawFrameContent;
        return this;
    }
    
    public DataqueryBuilder refId(String refId) {
        this.internal.refId = refId;
        return this;
    }
    
    public DataqueryBuilder resultAssertions(com.grafana.foundation.cog.Builder<ResultAssertions> resultAssertions) {
    ResultAssertions resultAssertionsResource = resultAssertions.build();
        this.internal.resultAssertions = resultAssertionsResource;
        return this;
    }
    
    public DataqueryBuilder scenarioId(DataqueryScenarioId scenarioId) {
        this.internal.scenarioId = scenarioId;
        return this;
    }
    
    public DataqueryBuilder seriesCount(Long seriesCount) {
        this.internal.seriesCount = seriesCount;
        return this;
    }
    
    public DataqueryBuilder sim(com.grafana.foundation.cog.Builder<SimulationQuery> sim) {
    SimulationQuery simResource = sim.build();
        this.internal.sim = simResource;
        return this;
    }
    
    public DataqueryBuilder spanCount(Long spanCount) {
        this.internal.spanCount = spanCount;
        return this;
    }
    
    public DataqueryBuilder spread(Double spread) {
        this.internal.spread = spread;
        return this;
    }
    
    public DataqueryBuilder startValue(Double startValue) {
        this.internal.startValue = startValue;
        return this;
    }
    
    public DataqueryBuilder stream(com.grafana.foundation.cog.Builder<StreamingQuery> stream) {
    StreamingQuery streamResource = stream.build();
        this.internal.stream = streamResource;
        return this;
    }
    
    public DataqueryBuilder stringInput(String stringInput) {
        this.internal.stringInput = stringInput;
        return this;
    }
    
    public DataqueryBuilder timeRange(com.grafana.foundation.cog.Builder<TimeRange> timeRange) {
    TimeRange timeRangeResource = timeRange.build();
        this.internal.timeRange = timeRangeResource;
        return this;
    }
    
    public DataqueryBuilder usa(com.grafana.foundation.cog.Builder<USAQuery> usa) {
    USAQuery usaResource = usa.build();
        this.internal.usa = usaResource;
        return this;
    }
    
    public DataqueryBuilder withNil(Boolean withNil) {
        this.internal.withNil = withNil;
        return this;
    }
    public Dataquery build() {
        return this.internal;
    }
}
