<?php

namespace Grafana\Foundation\Alerting;

class RuleGroup implements \JsonSerializable
{
    public ?string $folderUid;

    /**
     * The interval, in seconds, at which all rules in the group are evaluated.
     * If a group contains many rules, the rules are evaluated sequentially.
     */
    public int $interval;

    /**
     * @var array<\Grafana\Foundation\Alerting\Rule>|null
     */
    public ?array $rules;

    public ?string $title;

    /**
     * @param string|null $folderUid
     * @param int|null $interval
     * @param array<\Grafana\Foundation\Alerting\Rule>|null $rules
     * @param string|null $title
     */
    public function __construct(?string $folderUid = null, ?int $interval = null, ?array $rules = null, ?string $title = null)
    {
        $this->folderUid = $folderUid;
        $this->interval = $interval ?: 0;
        $this->rules = $rules;
        $this->title = $title;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{folderUid?: string, interval?: int, rules?: array<mixed>, title?: string} $inputData */
        $data = $inputData;
        return new self(
            folderUid: $data["folderUid"] ?? null,
            interval: $data["interval"] ?? null,
            rules: array_filter(array_map((function($input) {
    	/** @var array{annotations?: array<string, string>, condition?: string, data?: array<mixed>, execErrState?: string, folderUID?: string, for?: string, id?: int, isPaused?: bool, labels?: array<string, string>, noDataState?: string, notification_settings?: mixed, orgID?: int, provenance?: string, ruleGroup?: string, title?: string, uid?: string, updated?: string} */
    $val = $input;
    	return \Grafana\Foundation\Alerting\Rule::fromArray($val);
    }), $data["rules"] ?? [])),
            title: $data["title"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "interval" => $this->interval,
        ];
        if (isset($this->folderUid)) {
            $data["folderUid"] = $this->folderUid;
        }
        if (isset($this->rules)) {
            $data["rules"] = $this->rules;
        }
        if (isset($this->title)) {
            $data["title"] = $this->title;
        }
        return $data;
    }
}
