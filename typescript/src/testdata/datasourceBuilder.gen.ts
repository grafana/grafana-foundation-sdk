// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as testdata from '../testdata';

export class DatasourceBuilder implements cog.Builder<testdata.datasource> {
    protected readonly internal: testdata.datasource;

    constructor() {
        this.internal = testdata.defaultDatasource();
    }

    build(): testdata.datasource {
        return this.internal;
    }

    // The datasource plugin type
    type(type: string): this {
        this.internal.type = type;
        return this;
    }

    // Datasource UID
    uid(uid: string): this {
        this.internal.uid = uid;
        return this;
    }
}
