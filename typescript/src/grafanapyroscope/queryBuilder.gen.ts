// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
import * as grafanapyroscope from '../grafanapyroscope';

export class QueryBuilder implements cog.Builder<dashboardv2beta1.DataQueryKind> {
    protected readonly internal: dashboardv2beta1.DataQueryKind;

    constructor() {
        this.internal = dashboardv2beta1.defaultDataQueryKind();
        this.internal.kind = "DataQuery";
        this.internal.group = "grafanapyroscope";
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.DataQueryKind {
        return this.internal;
    }

    version(version: string): this {
        this.internal.version = version;
        return this;
    }

    // New type for datasource reference
    // Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.
    datasource(ref: {
	name?: string;
}): this {
        this.internal.datasource = ref;
        return this;
    }

    // Specifies the query label selectors.
    labelSelector(labelSelector: string): this {
        if (!this.internal.spec) {
            this.internal.spec = grafanapyroscope.defaultDataquery();
        }
        this.internal.spec.labelSelector = labelSelector;
        return this;
    }

    // Specifies the query span selectors.
    spanSelector(spanSelector: string[]): this {
        if (!this.internal.spec) {
            this.internal.spec = grafanapyroscope.defaultDataquery();
        }
        this.internal.spec.spanSelector = spanSelector;
        return this;
    }

    // Specifies the type of profile to query.
    profileTypeId(profileTypeId: string): this {
        if (!this.internal.spec) {
            this.internal.spec = grafanapyroscope.defaultDataquery();
        }
        this.internal.spec.profileTypeId = profileTypeId;
        return this;
    }

    // Allows to group the results.
    groupBy(groupBy: string[]): this {
        if (!this.internal.spec) {
            this.internal.spec = grafanapyroscope.defaultDataquery();
        }
        this.internal.spec.groupBy = groupBy;
        return this;
    }

    // Sets the maximum number of time series.
    limit(limit: number): this {
        if (!this.internal.spec) {
            this.internal.spec = grafanapyroscope.defaultDataquery();
        }
        this.internal.spec.limit = limit;
        return this;
    }

    // Sets the maximum number of nodes in the flamegraph.
    maxNodes(maxNodes: number): this {
        if (!this.internal.spec) {
            this.internal.spec = grafanapyroscope.defaultDataquery();
        }
        this.internal.spec.maxNodes = maxNodes;
        return this;
    }

    // A unique identifier for the query within the list of targets.
    // In server side expressions, the refId is used as a variable name to identify results.
    // By default, the UI will assign A->Z; however setting meaningful names may be useful.
    refId(refId: string): this {
        if (!this.internal.spec) {
            this.internal.spec = grafanapyroscope.defaultDataquery();
        }
        this.internal.spec.refId = refId;
        return this;
    }

    // If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
    hide(hide: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = grafanapyroscope.defaultDataquery();
        }
        this.internal.spec.hide = hide;
        return this;
    }

    // Specify the query flavor
    // TODO make this required and give it a default
    queryType(queryType: string): this {
        if (!this.internal.spec) {
            this.internal.spec = grafanapyroscope.defaultDataquery();
        }
        this.internal.spec.queryType = queryType;
        return this;
    }
}

