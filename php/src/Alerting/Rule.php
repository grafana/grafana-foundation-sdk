<?php

namespace Grafana\Foundation\Alerting;

class Rule implements \JsonSerializable
{
    /**
     * @var array<string, string>|null
     */
    public ?array $annotations;

    public string $condition;

    /**
     * @var array<\Grafana\Foundation\Alerting\Query>
     */
    public array $data;

    public \Grafana\Foundation\Alerting\RuleExecErrState $execErrState;

    public string $folderUID;

    /**
     * The amount of time, in seconds, for which the rule must be breached for the rule to be considered to be Firing.
     * Before this time has elapsed, the rule is only considered to be Pending.
     */
    public string $for;

    public ?int $id;

    public ?bool $isPaused;

    public ?string $keepFiringFor;

    /**
     * @var array<string, string>|null
     */
    public ?array $labels;

    public ?int $missingSeriesEvalsToResolve;

    public \Grafana\Foundation\Alerting\RuleNoDataState $noDataState;

    public ?\Grafana\Foundation\Alerting\NotificationSettings $notificationSettings;

    public int $orgID;

    public string $provenance;

    public ?\Grafana\Foundation\Alerting\RecordRule $record;

    public string $ruleGroup;

    public string $title;

    public ?string $uid;

    public ?string $updated;

    /**
     * @param array<string, string>|null $annotations
     * @param string|null $condition
     * @param array<\Grafana\Foundation\Alerting\Query>|null $data
     * @param \Grafana\Foundation\Alerting\RuleExecErrState|null $execErrState
     * @param string|null $folderUID
     * @param string|null $for
     * @param int|null $id
     * @param bool|null $isPaused
     * @param string|null $keepFiringFor
     * @param array<string, string>|null $labels
     * @param int|null $missingSeriesEvalsToResolve
     * @param \Grafana\Foundation\Alerting\RuleNoDataState|null $noDataState
     * @param \Grafana\Foundation\Alerting\NotificationSettings|null $notificationSettings
     * @param int|null $orgID
     * @param string|null $provenance
     * @param \Grafana\Foundation\Alerting\RecordRule|null $record
     * @param string|null $ruleGroup
     * @param string|null $title
     * @param string|null $uid
     * @param string|null $updated
     */
    public function __construct(?array $annotations = null, ?string $condition = null, ?array $data = null, ?\Grafana\Foundation\Alerting\RuleExecErrState $execErrState = null, ?string $folderUID = null, ?string $for = null, ?int $id = null, ?bool $isPaused = null, ?string $keepFiringFor = null, ?array $labels = null, ?int $missingSeriesEvalsToResolve = null, ?\Grafana\Foundation\Alerting\RuleNoDataState $noDataState = null, ?\Grafana\Foundation\Alerting\NotificationSettings $notificationSettings = null, ?int $orgID = null, ?string $provenance = null, ?\Grafana\Foundation\Alerting\RecordRule $record = null, ?string $ruleGroup = null, ?string $title = null, ?string $uid = null, ?string $updated = null)
    {
        $this->annotations = $annotations;
        $this->condition = $condition ?: "";
        $this->data = $data ?: [];
        $this->execErrState = $execErrState ?: \Grafana\Foundation\Alerting\RuleExecErrState::OK();
        $this->folderUID = $folderUID ?: "";
        $this->for = $for ?: "";
        $this->id = $id;
        $this->isPaused = $isPaused;
        $this->keepFiringFor = $keepFiringFor;
        $this->labels = $labels;
        $this->missingSeriesEvalsToResolve = $missingSeriesEvalsToResolve;
        $this->noDataState = $noDataState ?: \Grafana\Foundation\Alerting\RuleNoDataState::Alerting();
        $this->notificationSettings = $notificationSettings;
        $this->orgID = $orgID ?: 0;
        $this->provenance = $provenance ?: "";
        $this->record = $record;
        $this->ruleGroup = $ruleGroup ?: "";
        $this->title = $title ?: "";
        $this->uid = $uid;
        $this->updated = $updated;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{annotations?: array<string, string>, condition?: string, data?: array<mixed>, execErrState?: string, folderUID?: string, for?: string, id?: int, isPaused?: bool, keep_firing_for?: string, labels?: array<string, string>, missingSeriesEvalsToResolve?: int, noDataState?: string, notification_settings?: mixed, orgID?: int, provenance?: string, record?: mixed, ruleGroup?: string, title?: string, uid?: string, updated?: string} $inputData */
        $data = $inputData;
        return new self(
            annotations: $data["annotations"] ?? null,
            condition: $data["condition"] ?? null,
            data: array_filter(array_map((function($input) {
    	/** @var array{datasourceUid?: string, model?: mixed, queryType?: string, refId?: string, relativeTimeRange?: mixed} */
    $val = $input;
    	return \Grafana\Foundation\Alerting\Query::fromArray($val);
    }), $data["data"] ?? [])),
            execErrState: isset($data["execErrState"]) ? (function($input) { return \Grafana\Foundation\Alerting\RuleExecErrState::fromValue($input); })($data["execErrState"]) : null,
            folderUID: $data["folderUID"] ?? null,
            for: $data["for"] ?? null,
            id: $data["id"] ?? null,
            isPaused: $data["isPaused"] ?? null,
            keepFiringFor: $data["keep_firing_for"] ?? null,
            labels: $data["labels"] ?? null,
            missingSeriesEvalsToResolve: $data["missingSeriesEvalsToResolve"] ?? null,
            noDataState: isset($data["noDataState"]) ? (function($input) { return \Grafana\Foundation\Alerting\RuleNoDataState::fromValue($input); })($data["noDataState"]) : null,
            notificationSettings: isset($data["notification_settings"]) ? (function($input) {
    	/** @var array{group_by?: array<string>, group_interval?: string, group_wait?: string, mute_time_intervals?: array<string>, receiver?: string, repeat_interval?: string} */
    $val = $input;
    	return \Grafana\Foundation\Alerting\NotificationSettings::fromArray($val);
    })($data["notification_settings"]) : null,
            orgID: $data["orgID"] ?? null,
            provenance: $data["provenance"] ?? null,
            record: isset($data["record"]) ? (function($input) {
    	/** @var array{from?: string, metric?: string, target_datasource_uid?: string} */
    $val = $input;
    	return \Grafana\Foundation\Alerting\RecordRule::fromArray($val);
    })($data["record"]) : null,
            ruleGroup: $data["ruleGroup"] ?? null,
            title: $data["title"] ?? null,
            uid: $data["uid"] ?? null,
            updated: $data["updated"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->condition = $this->condition;
        $data->data = $this->data;
        $data->execErrState = $this->execErrState;
        $data->folderUID = $this->folderUID;
        $data->for = $this->for;
        $data->noDataState = $this->noDataState;
        $data->orgID = $this->orgID;
        $data->provenance = $this->provenance;
        $data->ruleGroup = $this->ruleGroup;
        $data->title = $this->title;
        if (isset($this->annotations)) {
            $data->annotations = $this->annotations;
        }
        if (isset($this->id)) {
            $data->id = $this->id;
        }
        if (isset($this->isPaused)) {
            $data->isPaused = $this->isPaused;
        }
        if (isset($this->keepFiringFor)) {
            $data->keep_firing_for = $this->keepFiringFor;
        }
        if (isset($this->labels)) {
            $data->labels = $this->labels;
        }
        if (isset($this->missingSeriesEvalsToResolve)) {
            $data->missingSeriesEvalsToResolve = $this->missingSeriesEvalsToResolve;
        }
        if (isset($this->notificationSettings)) {
            $data->notification_settings = $this->notificationSettings;
        }
        if (isset($this->record)) {
            $data->record = $this->record;
        }
        if (isset($this->uid)) {
            $data->uid = $this->uid;
        }
        if (isset($this->updated)) {
            $data->updated = $this->updated;
        }
        return $data;
    }
}
