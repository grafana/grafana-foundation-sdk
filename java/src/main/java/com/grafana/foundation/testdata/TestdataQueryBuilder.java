// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.testdata;

import com.grafana.foundation.dashboardv2beta1.DataQueryKind;
import java.util.Map;
import com.grafana.foundation.dashboardv2beta1.Dashboardv2beta1DataQueryKindDatasource;
import java.util.List;
import java.util.LinkedList;

public class TestdataQueryBuilder implements com.grafana.foundation.cog.Builder<DataQueryKind> {
    protected final DataQueryKind internal;
    
    public TestdataQueryBuilder() {
        this.internal = new DataQueryKind();
        this.internal.kind = "DataQuery";
        this.internal.group = "testdata";
    }
    public TestdataQueryBuilder version(String version) {
        this.internal.version = version;
        return this;
    }
    
    public TestdataQueryBuilder labels(Map<String, String> labels) {
        this.internal.labels = labels;
        return this;
    }
    
    public TestdataQueryBuilder datasource(com.grafana.foundation.cog.Builder<Dashboardv2beta1DataQueryKindDatasource> datasource) {
    Dashboardv2beta1DataQueryKindDatasource datasourceResource = datasource.build();
        this.internal.datasource = datasourceResource;
        return this;
    }
    
    public TestdataQueryBuilder alias(String alias) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.testdata.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).alias = alias;
        return this;
    }
    
    public TestdataQueryBuilder channel(String channel) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.testdata.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).channel = channel;
        return this;
    }
    
    public TestdataQueryBuilder csvContent(String csvContent) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.testdata.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).csvContent = csvContent;
        return this;
    }
    
    public TestdataQueryBuilder csvFileName(String csvFileName) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.testdata.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).csvFileName = csvFileName;
        return this;
    }
    
    public TestdataQueryBuilder csvWave(List<com.grafana.foundation.cog.Builder<CSVWave>> csvWave) {
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
    
    public TestdataQueryBuilder dropPercent(Double dropPercent) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.testdata.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).dropPercent = dropPercent;
        return this;
    }
    
    public TestdataQueryBuilder errorSource(DataqueryErrorSource errorSource) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.testdata.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).errorSource = errorSource;
        return this;
    }
    
    public TestdataQueryBuilder errorType(DataqueryErrorType errorType) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.testdata.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).errorType = errorType;
        return this;
    }
    
    public TestdataQueryBuilder flamegraphDiff(Boolean flamegraphDiff) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.testdata.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).flamegraphDiff = flamegraphDiff;
        return this;
    }
    
    public TestdataQueryBuilder hide(Boolean hide) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.testdata.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).hide = hide;
        return this;
    }
    
    public TestdataQueryBuilder intervalMs(Double intervalMs) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.testdata.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).intervalMs = intervalMs;
        return this;
    }
    
    public TestdataQueryBuilder dataqueryLabels(String labels) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.testdata.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).labels = labels;
        return this;
    }
    
    public TestdataQueryBuilder levelColumn(Boolean levelColumn) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.testdata.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).levelColumn = levelColumn;
        return this;
    }
    
    public TestdataQueryBuilder lines(Long lines) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.testdata.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).lines = lines;
        return this;
    }
    
    public TestdataQueryBuilder max(Double max) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.testdata.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).max = max;
        return this;
    }
    
    public TestdataQueryBuilder maxDataPoints(Long maxDataPoints) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.testdata.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).maxDataPoints = maxDataPoints;
        return this;
    }
    
    public TestdataQueryBuilder min(Double min) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.testdata.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).min = min;
        return this;
    }
    
    public TestdataQueryBuilder nodes(com.grafana.foundation.cog.Builder<NodesQuery> nodes) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.testdata.DataqueryBuilder().build();
		}
    NodesQuery nodesResource = nodes.build();
        ((Dataquery) this.internal.spec).nodes = nodesResource;
        return this;
    }
    
    public TestdataQueryBuilder noise(Double noise) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.testdata.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).noise = noise;
        return this;
    }
    
    public TestdataQueryBuilder points(List<List<Object>> points) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.testdata.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).points = points;
        return this;
    }
    
    public TestdataQueryBuilder pulseWave(com.grafana.foundation.cog.Builder<PulseWaveQuery> pulseWave) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.testdata.DataqueryBuilder().build();
		}
    PulseWaveQuery pulseWaveResource = pulseWave.build();
        ((Dataquery) this.internal.spec).pulseWave = pulseWaveResource;
        return this;
    }
    
    public TestdataQueryBuilder queryType(String queryType) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.testdata.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).queryType = queryType;
        return this;
    }
    
    public TestdataQueryBuilder rawFrameContent(String rawFrameContent) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.testdata.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).rawFrameContent = rawFrameContent;
        return this;
    }
    
    public TestdataQueryBuilder refId(String refId) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.testdata.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).refId = refId;
        return this;
    }
    
    public TestdataQueryBuilder resultAssertions(com.grafana.foundation.cog.Builder<ResultAssertions> resultAssertions) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.testdata.DataqueryBuilder().build();
		}
    ResultAssertions resultAssertionsResource = resultAssertions.build();
        ((Dataquery) this.internal.spec).resultAssertions = resultAssertionsResource;
        return this;
    }
    
    public TestdataQueryBuilder scenarioId(DataqueryScenarioId scenarioId) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.testdata.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).scenarioId = scenarioId;
        return this;
    }
    
    public TestdataQueryBuilder seriesCount(Long seriesCount) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.testdata.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).seriesCount = seriesCount;
        return this;
    }
    
    public TestdataQueryBuilder sim(com.grafana.foundation.cog.Builder<SimulationQuery> sim) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.testdata.DataqueryBuilder().build();
		}
    SimulationQuery simResource = sim.build();
        ((Dataquery) this.internal.spec).sim = simResource;
        return this;
    }
    
    public TestdataQueryBuilder spanCount(Long spanCount) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.testdata.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).spanCount = spanCount;
        return this;
    }
    
    public TestdataQueryBuilder spread(Double spread) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.testdata.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).spread = spread;
        return this;
    }
    
    public TestdataQueryBuilder startValue(Double startValue) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.testdata.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).startValue = startValue;
        return this;
    }
    
    public TestdataQueryBuilder stream(com.grafana.foundation.cog.Builder<StreamingQuery> stream) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.testdata.DataqueryBuilder().build();
		}
    StreamingQuery streamResource = stream.build();
        ((Dataquery) this.internal.spec).stream = streamResource;
        return this;
    }
    
    public TestdataQueryBuilder stringInput(String stringInput) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.testdata.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).stringInput = stringInput;
        return this;
    }
    
    public TestdataQueryBuilder timeRange(com.grafana.foundation.cog.Builder<TimeRange> timeRange) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.testdata.DataqueryBuilder().build();
		}
    TimeRange timeRangeResource = timeRange.build();
        ((Dataquery) this.internal.spec).timeRange = timeRangeResource;
        return this;
    }
    
    public TestdataQueryBuilder usa(com.grafana.foundation.cog.Builder<USAQuery> usa) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.testdata.DataqueryBuilder().build();
		}
    USAQuery usaResource = usa.build();
        ((Dataquery) this.internal.spec).usa = usaResource;
        return this;
    }
    
    public TestdataQueryBuilder withNil(Boolean withNil) {
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
