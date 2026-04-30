// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2 from '../dashboardv2';

export class DataLinkBuilder implements cog.Builder<dashboardv2.DataLink> {
    protected readonly internal: dashboardv2.DataLink;

    constructor() {
        this.internal = dashboardv2.defaultDataLink();
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2.DataLink {
        return this.internal;
    }

    title(title: string): this {
        this.internal.title = title;
        return this;
    }

    url(url: string): this {
        this.internal.url = url;
        return this;
    }

    targetBlank(targetBlank: boolean): this {
        this.internal.targetBlank = targetBlank;
        return this;
    }
}

