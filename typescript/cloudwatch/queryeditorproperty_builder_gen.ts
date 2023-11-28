// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as cloudwatch from '../cloudwatch';

export class QueryEditorPropertyBuilder implements cog.OptionsBuilder<cloudwatch.QueryEditorProperty> {
    private readonly internal: cloudwatch.QueryEditorProperty;

    constructor() {
        this.internal = cloudwatch.defaultQueryEditorProperty();
    }

    build(): cloudwatch.QueryEditorProperty {
        return this.internal;
    }

    type(type: cloudwatch.QueryEditorPropertyType): this {
        this.internal.type = type;
        return this;
    }

    name(name: string): this {
        this.internal.name = name;
        return this;
    }
}
