// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as testdata from '../testdata';

export class DatasourceBuilder implements cog.Builder<testdata.Datasource> {
    protected readonly internal: testdata.Datasource;

    constructor() {
        this.internal = testdata.defaultDatasource();
    }

    build(): testdata.Datasource {
        return this.internal;
    }

    // The apiserver version
    apiVersion(apiVersion: string): this {
        this.internal.apiVersion = apiVersion;
        return this;
    }

    // The datasource plugin type
    type(type: string): this {
        this.internal.type = type;
        return this;
    }

    // Datasource UID (NOTE: name in k8s)
    uid(uid: string): this {
        this.internal.uid = uid;
        return this;
    }
}
