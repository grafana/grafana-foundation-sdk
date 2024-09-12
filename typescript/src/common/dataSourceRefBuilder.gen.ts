// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as common from '../common';

export class DataSourceRefBuilder implements cog.Builder<common.DataSourceRef> {
    protected readonly internal: common.DataSourceRef;

    constructor() {
        this.internal = common.defaultDataSourceRef();
    }

    build(): common.DataSourceRef {
        return this.internal;
    }

    // The plugin type-id
    type(type: string): this {
        this.internal.type = type;
        return this;
    }

    // Specific datasource instance
    uid(uid: string): this {
        this.internal.uid = uid;
        return this;
    }

    // Datasource API version
    apiVersion(apiVersion: string): this {
        this.internal.apiVersion = apiVersion;
        return this;
    }
}
