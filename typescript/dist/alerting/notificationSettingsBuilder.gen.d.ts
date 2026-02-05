import * as cog from '../cog';
import * as alerting from '../alerting';
export declare class NotificationSettingsBuilder implements cog.Builder<alerting.NotificationSettings> {
    protected readonly internal: alerting.NotificationSettings;
    constructor();
    /**
     * Builds the object.
     */
    build(): alerting.NotificationSettings;
    groupBy(groupBy: string[]): this;
    groupInterval(groupInterval: string): this;
    groupWait(groupWait: string): this;
    muteTimeIntervals(muteTimeIntervals: string[]): this;
    receiver(receiver: string): this;
    repeatInterval(repeatInterval: string): this;
}
