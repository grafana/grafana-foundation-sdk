<?php

namespace Grafana\Foundation\Alerting;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alerting\NotificationSettings>
 */
class NotificationSettingsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Alerting\NotificationSettings $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Alerting\NotificationSettings();
    }

    /**
     * @return \Grafana\Foundation\Alerting\NotificationSettings
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * @param array<string> $groupBy
     */
    public function groupBy(array $groupBy): static
    {
        $this->internal->groupBy = $groupBy;
    
        return $this;
    }
    public function groupInterval(string $groupInterval): static
    {
        $this->internal->groupInterval = $groupInterval;
    
        return $this;
    }
    public function groupWait(string $groupWait): static
    {
        $this->internal->groupWait = $groupWait;
    
        return $this;
    }
    /**
     * @param array<string> $muteTimeIntervals
     */
    public function muteTimeIntervals(array $muteTimeIntervals): static
    {
        $this->internal->muteTimeIntervals = $muteTimeIntervals;
    
        return $this;
    }
    public function receiver(string $receiver): static
    {
        $this->internal->receiver = $receiver;
    
        return $this;
    }
    public function repeatInterval(string $repeatInterval): static
    {
        $this->internal->repeatInterval = $repeatInterval;
    
        return $this;
    }

}
