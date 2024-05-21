// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as alerting from '../alerting';

export class NotificationSettingsBuilder implements cog.Builder<alerting.NotificationSettings> {
    protected readonly internal: alerting.NotificationSettings;

    constructor() {
        this.internal = alerting.defaultNotificationSettings();
    }

    build(): alerting.NotificationSettings {
        return this.internal;
    }

    groupBy(groupBy: string[]): this {
        this.internal.group_by = groupBy;
        return this;
    }

    groupInterval(groupInterval: string): this {
        this.internal.group_interval = groupInterval;
        return this;
    }

    groupWait(groupWait: string): this {
        this.internal.group_wait = groupWait;
        return this;
    }

    muteTimeIntervals(muteTimeIntervals: string[]): this {
        this.internal.mute_time_intervals = muteTimeIntervals;
        return this;
    }

    receiver(receiver: string): this {
        this.internal.receiver = receiver;
        return this;
    }

    repeatInterval(repeatInterval: string): this {
        this.internal.repeat_interval = repeatInterval;
        return this;
    }
}
