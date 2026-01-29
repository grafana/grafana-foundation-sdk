// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';

// Annotation event field mapping. Defines how to map a data frame field to an annotation event field.
export class AnnotationEventFieldMappingBuilder implements cog.Builder<dashboardv2beta1.AnnotationEventFieldMapping> {
    protected readonly internal: dashboardv2beta1.AnnotationEventFieldMapping;

    constructor() {
        this.internal = dashboardv2beta1.defaultAnnotationEventFieldMapping();
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.AnnotationEventFieldMapping {
        return this.internal;
    }

    // Source type for the field value
    source(source: string): this {
        this.internal.source = source;
        return this;
    }

    // Constant value to use when source is "text"
    value(value: string): this {
        this.internal.value = value;
        return this;
    }

    // Regular expression to apply to the field value
    regex(regex: string): this {
        this.internal.regex = regex;
        return this;
    }
}

