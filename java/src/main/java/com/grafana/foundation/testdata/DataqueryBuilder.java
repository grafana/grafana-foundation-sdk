// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.testdata;

import java.util.List;
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
    
    public DataqueryBuilder csvWave(com.grafana.foundation.cog.Builder<List<CSVWave>> csvWave) {
    this.internal.csvWave = csvWave.build();
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
    this.internal.nodes = nodes.build();
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
    this.internal.pulseWave = pulseWave.build();
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
    this.internal.resultAssertions = resultAssertions.build();
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
    this.internal.sim = sim.build();
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
    this.internal.stream = stream.build();
        return this;
    }
    
    public DataqueryBuilder stringInput(String stringInput) {
    this.internal.stringInput = stringInput;
        return this;
    }
    
    public DataqueryBuilder timeRange(com.grafana.foundation.cog.Builder<TimeRange> timeRange) {
    this.internal.timeRange = timeRange.build();
        return this;
    }
    
    public DataqueryBuilder usa(com.grafana.foundation.cog.Builder<USAQuery> usa) {
    this.internal.usa = usa.build();
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
