// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboard from '../dashboard';

/**
 * Annotation event field mapping. Defines how to map a data frame field to an annotation event field.
 */
export class AnnotationEventFieldMappingBuilder implements cog.Builder<dashboard.AnnotationEventFieldMapping> {
    protected readonly internal: dashboard.AnnotationEventFieldMapping;

    constructor() {
        this.internal = dashboard.defaultAnnotationEventFieldMapping();
    }

    /**
     * Builds the object.
     */
    build(): dashboard.AnnotationEventFieldMapping {
        return this.internal;
    }

    // Source type for the field value.
    source(source: dashboard.AnnotationEventFieldSource): this {
        this.internal.source = source;
        return this;
    }

    // Constant value to use when source is "text".
    value(value: string): this {
        this.internal.value = value;
        return this;
    }

    // Regular expression to apply to the field value.
    regex(regex: string): this {
        this.internal.regex = regex;
        return this;
    }
}

