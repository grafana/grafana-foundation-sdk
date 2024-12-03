// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as athena from '../athena';

export class ConnectionArgsBuilder implements cog.Builder<athena.ConnectionArgs> {
    protected readonly internal: athena.ConnectionArgs;

    constructor() {
        this.internal = athena.defaultConnectionArgs();
    }

    /**
     * Builds the object.
     */
    build(): athena.ConnectionArgs {
        return this.internal;
    }

    region(region: string): this {
        this.internal.region = region;
        return this;
    }

    catalog(catalog: string): this {
        this.internal.catalog = catalog;
        return this;
    }

    database(database: string): this {
        this.internal.database = database;
        return this;
    }

    resultReuseEnabled(resultReuseEnabled: boolean): this {
        this.internal.resultReuseEnabled = resultReuseEnabled;
        return this;
    }

    resultReuseMaxAgeInMinutes(resultReuseMaxAgeInMinutes: number): this {
        this.internal.resultReuseMaxAgeInMinutes = resultReuseMaxAgeInMinutes;
        return this;
    }
}
