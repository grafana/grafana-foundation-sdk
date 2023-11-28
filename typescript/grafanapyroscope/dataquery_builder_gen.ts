// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as grafanapyroscope from '../grafanapyroscope';

export class DataqueryBuilder implements cog.OptionsBuilder<cog.Dataquery> {
    private readonly internal: grafanapyroscope.dataquery;

    constructor() {
        this.internal = grafanapyroscope.defaultDataquery();
    }

    build(): grafanapyroscope.dataquery {
        return this.internal;
    }

    labelSelector(labelSelector: string): this {
        this.internal.labelSelector = labelSelector;
        return this;
    }

    spanSelector(spanSelector: string[]): this {
        this.internal.spanSelector = spanSelector;
        return this;
    }

    profileTypeId(profileTypeId: string): this {
        this.internal.profileTypeId = profileTypeId;
        return this;
    }

    groupBy(groupBy: string[]): this {
        this.internal.groupBy = groupBy;
        return this;
    }

    maxNodes(maxNodes: number): this {
        this.internal.maxNodes = maxNodes;
        return this;
    }

    refId(refId: string): this {
        this.internal.refId = refId;
        return this;
    }

    hide(hide: boolean): this {
        this.internal.hide = hide;
        return this;
    }

    queryType(queryType: string): this {
        this.internal.queryType = queryType;
        return this;
    }

    datasource(datasource: any): this {
        this.internal.datasource = datasource;
        return this;
    }
}
