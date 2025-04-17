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

    /**
     * @var array<string, string>|null
     */
    public ?array $labels;

    public \Grafana\Foundation\Alerting\RuleNoDataState $noDataState;

    public int $orgID;

    public string $provenance;

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
     * @param array<string, string>|null $labels
     * @param \Grafana\Foundation\Alerting\RuleNoDataState|null $noDataState
     * @param int|null $orgID
     * @param string|null $provenance
     * @param string|null $ruleGroup
     * @param string|null $title
     * @param string|null $uid
     * @param string|null $updated
     */
    public function __construct(?array $annotations = null, ?string $condition = null, ?array $data = null, ?\Grafana\Foundation\Alerting\RuleExecErrState $execErrState = null, ?string $folderUID = null, ?string $for = null, ?int $id = null, ?bool $isPaused = null, ?array $labels = null, ?\Grafana\Foundation\Alerting\RuleNoDataState $noDataState = null, ?int $orgID = null, ?string $provenance = null, ?string $ruleGroup = null, ?string $title = null, ?string $uid = null, ?string $updated = null)
    {
        $this->annotations = $annotations;
        $this->condition = $condition ?: "";
        $this->data = $data ?: [];
        $this->execErrState = $execErrState ?: \Grafana\Foundation\Alerting\RuleExecErrState::OK();
        $this->folderUID = $folderUID ?: "";
        $this->for = $for ?: "";
        $this->id = $id;
        $this->isPaused = $isPaused;
        $this->labels = $labels;
        $this->noDataState = $noDataState ?: \Grafana\Foundation\Alerting\RuleNoDataState::Alerting();
        $this->orgID = $orgID ?: 0;
        $this->provenance = $provenance ?: "";
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
        /** @var array{annotations?: array<string, string>, condition?: string, data?: array<mixed>, execErrState?: string, folderUID?: string, for?: string, id?: int, isPaused?: bool, labels?: array<string, string>, noDataState?: string, orgID?: int, provenance?: string, ruleGroup?: string, title?: string, uid?: string, updated?: string} $inputData */
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
            labels: $data["labels"] ?? null,
            noDataState: isset($data["noDataState"]) ? (function($input) { return \Grafana\Foundation\Alerting\RuleNoDataState::fromValue($input); })($data["noDataState"]) : null,
            orgID: $data["orgID"] ?? null,
            provenance: $data["provenance"] ?? null,
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
        if (isset($this->labels)) {
            $data->labels = $this->labels;
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
