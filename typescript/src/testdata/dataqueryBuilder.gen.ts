// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as testdata from '../testdata';
import * as dashboard from '../dashboard';

export class DataqueryBuilder implements cog.Builder<cog.Dataquery> {
    protected readonly internal: testdata.dataquery;

    constructor() {
        this.internal = testdata.defaultDataquery();
    }

    /**
     * Builds the object.
     */
    build(): testdata.dataquery {
        return this.internal;
    }

    alias(alias: string): this {
        this.internal.alias = alias;
        return this;
    }

    scenarioId(scenarioId: testdata.TestDataQueryType): this {
        this.internal.scenarioId = scenarioId;
        return this;
    }

    stringInput(stringInput: string): this {
        this.internal.stringInput = stringInput;
        return this;
    }

    stream(stream: cog.Builder<testdata.StreamingQuery>): this {
        const streamResource = stream.build();
        this.internal.stream = streamResource;
        return this;
    }

    pulseWave(pulseWave: cog.Builder<testdata.PulseWaveQuery>): this {
        const pulseWaveResource = pulseWave.build();
        this.internal.pulseWave = pulseWaveResource;
        return this;
    }

    sim(sim: cog.Builder<testdata.SimulationQuery>): this {
        const simResource = sim.build();
        this.internal.sim = simResource;
        return this;
    }

    csvWave(csvWave: cog.Builder<testdata.CSVWave>[]): this {
        const csvWaveResources = csvWave.map(builder1 => builder1.build());
        this.internal.csvWave = csvWaveResources;
        return this;
    }

    labels(labels: string): this {
        this.internal.labels = labels;
        return this;
    }

    lines(lines: number): this {
        this.internal.lines = lines;
        return this;
    }

    levelColumn(levelColumn: boolean): this {
        this.internal.levelColumn = levelColumn;
        return this;
    }

    channel(channel: string): this {
        this.internal.channel = channel;
        return this;
    }

    nodes(nodes: cog.Builder<testdata.NodesQuery>): this {
        const nodesResource = nodes.build();
        this.internal.nodes = nodesResource;
        return this;
    }

    csvFileName(csvFileName: string): this {
        this.internal.csvFileName = csvFileName;
        return this;
    }

    csvContent(csvContent: string): this {
        this.internal.csvContent = csvContent;
        return this;
    }

    rawFrameContent(rawFrameContent: string): this {
        this.internal.rawFrameContent = rawFrameContent;
        return this;
    }

    seriesCount(seriesCount: number): this {
        this.internal.seriesCount = seriesCount;
        return this;
    }

    usa(usa: cog.Builder<testdata.USAQuery>): this {
        const usaResource = usa.build();
        this.internal.usa = usaResource;
        return this;
    }

    errorType(errorType: "server_panic" | "frontend_exception" | "frontend_observable"): this {
        this.internal.errorType = errorType;
        return this;
    }

    spanCount(spanCount: number): this {
        this.internal.spanCount = spanCount;
        return this;
    }

    points(points: (string | number)[][]): this {
        this.internal.points = points;
        return this;
    }

    // Drop percentage (the chance we will lose a point 0-100)
    dropPercent(dropPercent: number): this {
        this.internal.dropPercent = dropPercent;
        return this;
    }

    // A unique identifier for the query within the list of targets.
    // In server side expressions, the refId is used as a variable name to identify results.
    // By default, the UI will assign A->Z; however setting meaningful names may be useful.
    refId(refId: string): this {
        this.internal.refId = refId;
        return this;
    }

    // true if query is disabled (ie should not be returned to the dashboard)
    // Note this does not always imply that the query should not be executed since
    // the results from a hidden query may be used as the input to other queries (SSE etc)
    hide(hide: boolean): this {
        this.internal.hide = hide;
        return this;
    }

    // Specify the query flavor
    // TODO make this required and give it a default
    queryType(queryType: string): this {
        this.internal.queryType = queryType;
        return this;
    }

    // For mixed data sources the selected datasource is on the query level.
    // For non mixed scenarios this is undefined.
    // TODO find a better way to do this ^ that's friendly to schema
    // TODO this shouldn't be unknown but DataSourceRef | null
    datasource(datasource: dashboard.DataSourceRef): this {
        this.internal.datasource = datasource;
        return this;
    }
}
