// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as elasticsearch from '../elasticsearch';

export class TermsSettingsBuilder implements cog.Builder<elasticsearch.TermsSettings> {
    protected readonly internal: elasticsearch.TermsSettings;

    constructor() {
        this.internal = elasticsearch.defaultTermsSettings();
    }

    build(): elasticsearch.TermsSettings {
        return this.internal;
    }

    order(order: elasticsearch.TermsOrder): this {
        this.internal.order = order;
        return this;
    }

    size(size: string): this {
        this.internal.size = size;
        return this;
    }

    min_doc_count(min_doc_count: string): this {
        this.internal.min_doc_count = min_doc_count;
        return this;
    }

    orderBy(orderBy: string): this {
        this.internal.orderBy = orderBy;
        return this;
    }

    missing(missing: string): this {
        this.internal.missing = missing;
        return this;
    }
}
