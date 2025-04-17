<?php

namespace Grafana\Foundation\Alerting;

/**
 * A Route is a node that contains definitions of how to handle alerts. This is modified
 * from the upstream alertmanager in that it adds the ObjectMatchers property.
 */
class NotificationPolicy implements \JsonSerializable
{
    /**
     * @var array<string>|null
     */
    public ?array $activeTimeIntervals;

    public ?bool $continue;

    /**
     * @var array<string>|null
     */
    public ?array $groupBy;

    public ?string $groupInterval;

    public ?string $groupWait;

    /**
     * Deprecated. Remove before v1.0 release.
     * @var array<string, string>|null
     */
    public ?array $match;

    /**
     * @var array<string, string>
     */
    public array $matchRe;

    /**
     * Matchers is a slice of Matchers that is sortable, implements Stringer, and
     * provides a Matches method to match a LabelSet against all Matchers in the
     * slice. Note that some users of Matchers might require it to be sorted.
     * @var array<\Grafana\Foundation\Alerting\Matcher>
     */
    public array $matchers;

    /**
     * @var array<string>|null
     */
    public ?array $muteTimeIntervals;

    /**
     * @var array<array<string>>
     */
    public array $objectMatchers;

    public string $provenance;

    public ?string $receiver;

    public ?string $repeatInterval;

    /**
     * @var array<\Grafana\Foundation\Alerting\NotificationPolicy>|null
     */
    public ?array $routes;

    /**
     * @param array<string>|null $activeTimeIntervals
     * @param bool|null $continue
     * @param array<string>|null $groupBy
     * @param string|null $groupInterval
     * @param string|null $groupWait
     * @param array<string, string>|null $match
     * @param array<string, string>|null $matchRe
     * @param array<\Grafana\Foundation\Alerting\Matcher>|null $matchers
     * @param array<string>|null $muteTimeIntervals
     * @param array<array<string>>|null $objectMatchers
     * @param string|null $provenance
     * @param string|null $receiver
     * @param string|null $repeatInterval
     * @param array<\Grafana\Foundation\Alerting\NotificationPolicy>|null $routes
     */
    public function __construct(?array $activeTimeIntervals = null, ?bool $continue = null, ?array $groupBy = null, ?string $groupInterval = null, ?string $groupWait = null, ?array $match = null, ?array $matchRe = null, ?array $matchers = null, ?array $muteTimeIntervals = null, ?array $objectMatchers = null, ?string $provenance = null, ?string $receiver = null, ?string $repeatInterval = null, ?array $routes = null)
    {
        $this->activeTimeIntervals = $activeTimeIntervals;
        $this->continue = $continue;
        $this->groupBy = $groupBy;
        $this->groupInterval = $groupInterval;
        $this->groupWait = $groupWait;
        $this->match = $match;
        $this->matchRe = $matchRe ?: [];
        $this->matchers = $matchers ?: [];
        $this->muteTimeIntervals = $muteTimeIntervals;
        $this->objectMatchers = $objectMatchers ?: [];
        $this->provenance = $provenance ?: "";
        $this->receiver = $receiver;
        $this->repeatInterval = $repeatInterval;
        $this->routes = $routes;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{active_time_intervals?: array<string>, continue?: bool, group_by?: array<string>, group_interval?: string, group_wait?: string, match?: array<string, string>, match_re?: array<string, string>, matchers?: array<mixed>, mute_time_intervals?: array<string>, object_matchers?: array<array<string>>, provenance?: string, receiver?: string, repeat_interval?: string, routes?: array<mixed>} $inputData */
        $data = $inputData;
        return new self(
            activeTimeIntervals: $data["active_time_intervals"] ?? null,
            continue: $data["continue"] ?? null,
            groupBy: $data["group_by"] ?? null,
            groupInterval: $data["group_interval"] ?? null,
            groupWait: $data["group_wait"] ?? null,
            match: $data["match"] ?? null,
            matchRe: $data["match_re"] ?? null,
            matchers: array_filter(array_map((function($input) {
    	/** @var array{Name?: string, Type?: string, Value?: string} */
    $val = $input;
    	return \Grafana\Foundation\Alerting\Matcher::fromArray($val);
    }), $data["matchers"] ?? [])),
            muteTimeIntervals: $data["mute_time_intervals"] ?? null,
            objectMatchers: $data["object_matchers"] ?? null,
            provenance: $data["provenance"] ?? null,
            receiver: $data["receiver"] ?? null,
            repeatInterval: $data["repeat_interval"] ?? null,
            routes: array_filter(array_map((function($input) {
    	/** @var array{active_time_intervals?: array<string>, continue?: bool, group_by?: array<string>, group_interval?: string, group_wait?: string, match?: array<string, string>, match_re?: array<string, string>, matchers?: array<mixed>, mute_time_intervals?: array<string>, object_matchers?: array<array<string>>, provenance?: string, receiver?: string, repeat_interval?: string, routes?: array<mixed>} */
    $val = $input;
    	return \Grafana\Foundation\Alerting\NotificationPolicy::fromArray($val);
    }), $data["routes"] ?? [])),
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->match_re = $this->matchRe;
        $data->matchers = $this->matchers;
        $data->object_matchers = $this->objectMatchers;
        $data->provenance = $this->provenance;
        if (isset($this->activeTimeIntervals)) {
            $data->active_time_intervals = $this->activeTimeIntervals;
        }
        if (isset($this->continue)) {
            $data->continue = $this->continue;
        }
        if (isset($this->groupBy)) {
            $data->group_by = $this->groupBy;
        }
        if (isset($this->groupInterval)) {
            $data->group_interval = $this->groupInterval;
        }
        if (isset($this->groupWait)) {
            $data->group_wait = $this->groupWait;
        }
        if (isset($this->match)) {
            $data->match = $this->match;
        }
        if (isset($this->muteTimeIntervals)) {
            $data->mute_time_intervals = $this->muteTimeIntervals;
        }
        if (isset($this->receiver)) {
            $data->receiver = $this->receiver;
        }
        if (isset($this->repeatInterval)) {
            $data->repeat_interval = $this->repeatInterval;
        }
        if (isset($this->routes)) {
            $data->routes = $this->routes;
        }
        return $data;
    }
}
