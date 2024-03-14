// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as parca from '../parca';

export class DataqueryBuilder implements cog.Builder<cog.Dataquery> {
    private readonly internal: parca.dataquery;

    constructor() {
        this.internal = parca.defaultDataquery();
    }

    build(): parca.dataquery {
        return this.internal;
    }

    labelSelector(labelSelector: string): this {
        this.internal.labelSelector = labelSelector;
        return this;
    }

    profileTypeId(profileTypeId: string): this {
        this.internal.profileTypeId = profileTypeId;
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
