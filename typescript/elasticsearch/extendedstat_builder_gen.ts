// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as elasticsearch from '../elasticsearch';

export class ExtendedStatBuilder implements cog.OptionsBuilder<elasticsearch.ExtendedStat> {
    private readonly internal: elasticsearch.ExtendedStat;

    constructor() {
        this.internal = elasticsearch.defaultExtendedStat();
    }

    build(): elasticsearch.ExtendedStat {
        return this.internal;
    }

    label(label: string): this {
        this.internal.label = label;
        return this;
    }

    value(value: elasticsearch.ExtendedStatMetaType): this {
        this.internal.value = value;
        return this;
    }
}
