// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as azuremonitor from '../azuremonitor';

export class BuilderQueryEditorGroupByExpressionBuilder implements cog.Builder<azuremonitor.BuilderQueryEditorGroupByExpression> {
    protected readonly internal: azuremonitor.BuilderQueryEditorGroupByExpression;

    constructor() {
        this.internal = azuremonitor.defaultBuilderQueryEditorGroupByExpression();
    }

    /**
     * Builds the object.
     */
    build(): azuremonitor.BuilderQueryEditorGroupByExpression {
        return this.internal;
    }

    property(property: cog.Builder<azuremonitor.BuilderQueryEditorProperty>): this {
        const propertyResource = property.build();
        this.internal.property = propertyResource;
        return this;
    }

    interval(interval: cog.Builder<azuremonitor.BuilderQueryEditorProperty>): this {
        const intervalResource = interval.build();
        this.internal.interval = intervalResource;
        return this;
    }

    focus(focus: boolean): this {
        this.internal.focus = focus;
        return this;
    }

    type(type: azuremonitor.BuilderQueryEditorExpressionType): this {
        this.internal.type = type;
        return this;
    }
}

