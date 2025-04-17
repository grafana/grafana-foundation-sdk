<?php

namespace Grafana\Foundation\Annotationslist;

class Options implements \JsonSerializable
{
    public bool $onlyFromThisDashboard;

    public bool $onlyInTimeRange;

    /**
     * @var array<string>
     */
    public array $tags;

    public int $limit;

    public bool $showUser;

    public bool $showTime;

    public bool $showTags;

    public bool $navigateToPanel;

    public string $navigateBefore;

    public string $navigateAfter;

    /**
     * @param bool|null $onlyFromThisDashboard
     * @param bool|null $onlyInTimeRange
     * @param array<string>|null $tags
     * @param int|null $limit
     * @param bool|null $showUser
     * @param bool|null $showTime
     * @param bool|null $showTags
     * @param bool|null $navigateToPanel
     * @param string|null $navigateBefore
     * @param string|null $navigateAfter
     */
    public function __construct(?bool $onlyFromThisDashboard = null, ?bool $onlyInTimeRange = null, ?array $tags = null, ?int $limit = null, ?bool $showUser = null, ?bool $showTime = null, ?bool $showTags = null, ?bool $navigateToPanel = null, ?string $navigateBefore = null, ?string $navigateAfter = null)
    {
        $this->onlyFromThisDashboard = $onlyFromThisDashboard ?: false;
        $this->onlyInTimeRange = $onlyInTimeRange ?: false;
        $this->tags = $tags ?: [];
        $this->limit = $limit ?: 10;
        $this->showUser = $showUser ?: true;
        $this->showTime = $showTime ?: true;
        $this->showTags = $showTags ?: true;
        $this->navigateToPanel = $navigateToPanel ?: true;
        $this->navigateBefore = $navigateBefore ?: "10m";
        $this->navigateAfter = $navigateAfter ?: "10m";
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{onlyFromThisDashboard?: bool, onlyInTimeRange?: bool, tags?: array<string>, limit?: int, showUser?: bool, showTime?: bool, showTags?: bool, navigateToPanel?: bool, navigateBefore?: string, navigateAfter?: string} $inputData */
        $data = $inputData;
        return new self(
            onlyFromThisDashboard: $data["onlyFromThisDashboard"] ?? null,
            onlyInTimeRange: $data["onlyInTimeRange"] ?? null,
            tags: $data["tags"] ?? null,
            limit: $data["limit"] ?? null,
            showUser: $data["showUser"] ?? null,
            showTime: $data["showTime"] ?? null,
            showTags: $data["showTags"] ?? null,
            navigateToPanel: $data["navigateToPanel"] ?? null,
            navigateBefore: $data["navigateBefore"] ?? null,
            navigateAfter: $data["navigateAfter"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->onlyFromThisDashboard = $this->onlyFromThisDashboard;
        $data->onlyInTimeRange = $this->onlyInTimeRange;
        $data->tags = $this->tags;
        $data->limit = $this->limit;
        $data->showUser = $this->showUser;
        $data->showTime = $this->showTime;
        $data->showTags = $this->showTags;
        $data->navigateToPanel = $this->navigateToPanel;
        $data->navigateBefore = $this->navigateBefore;
        $data->navigateAfter = $this->navigateAfter;
        return $data;
    }
}
