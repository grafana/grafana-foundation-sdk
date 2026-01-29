// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboard from '../dashboard';

// Ref to a DataSource instance
export class DataSourceRefBuilder implements cog.Builder<dashboard.DataSourceRef> {
    protected readonly internal: dashboard.DataSourceRef;

    constructor() {
        this.internal = dashboard.defaultDataSourceRef();
    }

    /**
     * Builds the object.
     */
    build(): dashboard.DataSourceRef {
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
}

