// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as datasource from '../datasource';
import * as common from '../common';

export class DataqueryBuilder implements cog.Builder<cog.Dataquery> {
    protected readonly internal: datasource.Dataquery;

    constructor() {
        this.internal = datasource.defaultDataquery();
    }

    /**
     * Builds the object.
     */
    build(): datasource.Dataquery {
        return this.internal;
    }

    // Panel ID from wich the queries will be reused.
    panelId(panelId: number): this {
        this.internal.panelId = panelId;
        return this;
    }

    // A unique identifier for the query within the list of targets.
    // In server side expressions, the refId is used as a variable name to identify results.
    // By default, the UI will assign A->Z; however setting meaningful names may be useful.
    refId(refId: string): this {
        this.internal.refId = refId;
        return this;
    }

    // If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
    hide(hide: boolean): this {
        this.internal.hide = hide;
        return this;
    }

    withTransforms(withTransforms: boolean): this {
        this.internal.withTransforms = withTransforms;
        return this;
    }

    // For mixed data sources the selected datasource is on the query level.
    // For non mixed scenarios this is undefined.
    // TODO find a better way to do this ^ that's friendly to schema
    // TODO this shouldn't be unknown but DataSourceRef | null
    datasource(ref: common.DataSourceRef): this {
        this.internal.datasource = ref;
        return this;
    }
}

