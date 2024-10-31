// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as grafanapyroscope from '../grafanapyroscope';
import * as dashboard from '../dashboard';

export class DataqueryBuilder implements cog.Builder<cog.Dataquery> {
    protected readonly internal: grafanapyroscope.dataquery;

    constructor() {
        this.internal = grafanapyroscope.defaultDataquery();
    }

    /**
     * Builds the object.
     */
    build(): grafanapyroscope.dataquery {
        return this.internal;
    }

    // Specifies the query label selectors.
    labelSelector(labelSelector: string): this {
        this.internal.labelSelector = labelSelector;
        return this;
    }

    // Specifies the type of profile to query.
    profileTypeId(profileTypeId: string): this {
        this.internal.profileTypeId = profileTypeId;
        return this;
    }

    // Allows to group the results.
    groupBy(groupBy: string[]): this {
        this.internal.groupBy = groupBy;
        return this;
    }

    // Sets the maximum number of nodes in the flamegraph.
    maxNodes(maxNodes: number): this {
        this.internal.maxNodes = maxNodes;
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
