import * as cog from '../cog';
import * as alerting from '../alerting';
export declare class NotificationTemplateBuilder implements cog.Builder<alerting.NotificationTemplate> {
    protected readonly internal: alerting.NotificationTemplate;
    constructor();
    /**
     * Builds the object.
     */
    build(): alerting.NotificationTemplate;
    name(name: string): this;
    provenance(provenance: alerting.Provenance): this;
    template(template: string): this;
    version(version: string): this;
}
