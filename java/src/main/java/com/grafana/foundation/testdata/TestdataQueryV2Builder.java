// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.testdata;

import com.grafana.foundation.dashboardv2.DataQueryKind;
import java.util.Map;
import com.grafana.foundation.dashboardv2.Dashboardv2DataQueryKindDatasource;
import java.util.List;
import java.util.LinkedList;

public class TestdataQueryV2Builder implements com.grafana.foundation.cog.Builder<DataQueryKind> {
    protected final DataQueryKind internal;
    
    public TestdataQueryV2Builder() {
        this.internal = new DataQueryKind();
        this.internal.kind = "DataQuery";
        this.internal.group = "testdata";
    }
    public TestdataQueryV2Builder version(String version) {
        this.internal.version = version;
        return this;
    }
    
    public TestdataQueryV2Builder labels(Map<String, String> labels) {
        this.internal.labels = labels;
        return this;
    }
    
    public TestdataQueryV2Builder datasource(com.grafana.foundation.cog.Builder<Dashboardv2DataQueryKindDatasource> datasource) {
    Dashboardv2DataQueryKindDatasource datasourceResource = datasource.build();
        this.internal.datasource = datasourceResource;
        return this;
    }
    
    public TestdataQueryV2Builder alias(String alias) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.testdata.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).alias = alias;
        return this;
    }
    
    public TestdataQueryV2Builder channel(String channel) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.testdata.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).channel = channel;
        return this;
    }
    
    public TestdataQueryV2Builder csvContent(String csvContent) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.testdata.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).csvContent = csvContent;
        return this;
    }
    
    public TestdataQueryV2Builder csvFileName(String csvFileName) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.testdata.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).csvFileName = csvFileName;
        return this;
    }
    
    public TestdataQueryV2Builder csvWave(List<com.grafana.foundation.cog.Builder<CSVWave>> csvWave) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.testdata.DataqueryBuilder().build();
		}
        List<CSVWave> csvWaveResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<CSVWave> r1 : csvWave) {
                CSVWave csvWaveDepth1 = r1.build();
                csvWaveResources.add(csvWaveDepth1); 
        }
        ((Dataquery) this.internal.spec).csvWave = csvWaveResources;
        return this;
    }
    
    public TestdataQueryV2Builder dropPercent(Double dropPercent) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.testdata.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).dropPercent = dropPercent;
        return this;
    }
    
    public TestdataQueryV2Builder errorSource(DataqueryErrorSource errorSource) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.testdata.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).errorSource = errorSource;
        return this;
    }
    
    public TestdataQueryV2Builder errorType(DataqueryErrorType errorType) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.testdata.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).errorType = errorType;
        return this;
    }
    
    public TestdataQueryV2Builder flamegraphDiff(Boolean flamegraphDiff) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.testdata.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).flamegraphDiff = flamegraphDiff;
        return this;
    }
    
    public TestdataQueryV2Builder hide(Boolean hide) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.testdata.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).hide = hide;
        return this;
    }
    
    public TestdataQueryV2Builder intervalMs(Double intervalMs) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.testdata.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).intervalMs = intervalMs;
        return this;
    }
    
    public TestdataQueryV2Builder dataqueryLabels(String labels) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.testdata.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).labels = labels;
        return this;
    }
    
    public TestdataQueryV2Builder levelColumn(Boolean levelColumn) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.testdata.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).levelColumn = levelColumn;
        return this;
    }
    
    public TestdataQueryV2Builder lines(Long lines) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.testdata.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).lines = lines;
        return this;
    }
    
    public TestdataQueryV2Builder max(Double max) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.testdata.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).max = max;
        return this;
    }
    
    public TestdataQueryV2Builder maxDataPoints(Long maxDataPoints) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.testdata.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).maxDataPoints = maxDataPoints;
        return this;
    }
    
    public TestdataQueryV2Builder min(Double min) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.testdata.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).min = min;
        return this;
    }
    
    public TestdataQueryV2Builder nodes(com.grafana.foundation.cog.Builder<NodesQuery> nodes) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.testdata.DataqueryBuilder().build();
		}
    NodesQuery nodesResource = nodes.build();
        ((Dataquery) this.internal.spec).nodes = nodesResource;
        return this;
    }
    
    public TestdataQueryV2Builder noise(Double noise) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.testdata.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).noise = noise;
        return this;
    }
    
    public TestdataQueryV2Builder points(List<List<Object>> points) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.testdata.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).points = points;
        return this;
    }
    
    public TestdataQueryV2Builder pulseWave(com.grafana.foundation.cog.Builder<PulseWaveQuery> pulseWave) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.testdata.DataqueryBuilder().build();
		}
    PulseWaveQuery pulseWaveResource = pulseWave.build();
        ((Dataquery) this.internal.spec).pulseWave = pulseWaveResource;
        return this;
    }
    
    public TestdataQueryV2Builder queryType(String queryType) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.testdata.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).queryType = queryType;
        return this;
    }
    
    public TestdataQueryV2Builder rawFrameContent(String rawFrameContent) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.testdata.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).rawFrameContent = rawFrameContent;
        return this;
    }
    
    public TestdataQueryV2Builder refId(String refId) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.testdata.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).refId = refId;
        return this;
    }
    
    public TestdataQueryV2Builder resultAssertions(com.grafana.foundation.cog.Builder<ResultAssertions> resultAssertions) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.testdata.DataqueryBuilder().build();
		}
    ResultAssertions resultAssertionsResource = resultAssertions.build();
        ((Dataquery) this.internal.spec).resultAssertions = resultAssertionsResource;
        return this;
    }
    
    public TestdataQueryV2Builder scenarioId(DataqueryScenarioId scenarioId) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.testdata.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).scenarioId = scenarioId;
        return this;
    }
    
    public TestdataQueryV2Builder seriesCount(Long seriesCount) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.testdata.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).seriesCount = seriesCount;
        return this;
    }
    
    public TestdataQueryV2Builder sim(com.grafana.foundation.cog.Builder<SimulationQuery> sim) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.testdata.DataqueryBuilder().build();
		}
    SimulationQuery simResource = sim.build();
        ((Dataquery) this.internal.spec).sim = simResource;
        return this;
    }
    
    public TestdataQueryV2Builder spanCount(Long spanCount) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.testdata.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).spanCount = spanCount;
        return this;
    }
    
    public TestdataQueryV2Builder spread(Double spread) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.testdata.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).spread = spread;
        return this;
    }
    
    public TestdataQueryV2Builder startValue(Double startValue) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.testdata.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).startValue = startValue;
        return this;
    }
    
    public TestdataQueryV2Builder stream(com.grafana.foundation.cog.Builder<StreamingQuery> stream) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.testdata.DataqueryBuilder().build();
		}
    StreamingQuery streamResource = stream.build();
        ((Dataquery) this.internal.spec).stream = streamResource;
        return this;
    }
    
    public TestdataQueryV2Builder stringInput(String stringInput) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.testdata.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).stringInput = stringInput;
        return this;
    }
    
    public TestdataQueryV2Builder timeRange(com.grafana.foundation.cog.Builder<TimeRange> timeRange) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.testdata.DataqueryBuilder().build();
		}
    TimeRange timeRangeResource = timeRange.build();
        ((Dataquery) this.internal.spec).timeRange = timeRangeResource;
        return this;
    }
    
    public TestdataQueryV2Builder usa(com.grafana.foundation.cog.Builder<USAQuery> usa) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.testdata.DataqueryBuilder().build();
		}
    USAQuery usaResource = usa.build();
        ((Dataquery) this.internal.spec).usa = usaResource;
        return this;
    }
    
    public TestdataQueryV2Builder withNil(Boolean withNil) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.testdata.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).withNil = withNil;
        return this;
    }
    public DataQueryKind build() {
        return this.internal;
    }
}
