<?php

namespace Grafana\Foundation\Dashboardlist;

class Options implements \JsonSerializable
{
    public bool $keepTime;

    public bool $includeVars;

    public bool $showStarred;

    public bool $showRecentlyViewed;

    public bool $showSearch;

    public bool $showHeadings;

    public int $maxItems;

    public string $query;

    /**
     * @var array<string>
     */
    public array $tags;

    /**
     * folderId is deprecated, and migrated to folderUid on panel init
     */
    public ?int $folderId;

    public ?string $folderUID;

    /**
     * @param bool|null $keepTime
     * @param bool|null $includeVars
     * @param bool|null $showStarred
     * @param bool|null $showRecentlyViewed
     * @param bool|null $showSearch
     * @param bool|null $showHeadings
     * @param int|null $maxItems
     * @param string|null $query
     * @param array<string>|null $tags
     * @param int|null $folderId
     * @param string|null $folderUID
     */
    public function __construct(?bool $keepTime = null, ?bool $includeVars = null, ?bool $showStarred = null, ?bool $showRecentlyViewed = null, ?bool $showSearch = null, ?bool $showHeadings = null, ?int $maxItems = null, ?string $query = null, ?array $tags = null, ?int $folderId = null, ?string $folderUID = null)
    {
        $this->keepTime = $keepTime ?: false;
        $this->includeVars = $includeVars ?: false;
        $this->showStarred = $showStarred ?: true;
        $this->showRecentlyViewed = $showRecentlyViewed ?: false;
        $this->showSearch = $showSearch ?: false;
        $this->showHeadings = $showHeadings ?: true;
        $this->maxItems = $maxItems ?: 10;
        $this->query = $query ?: "";
        $this->tags = $tags ?: [];
        $this->folderId = $folderId;
        $this->folderUID = $folderUID;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{keepTime?: bool, includeVars?: bool, showStarred?: bool, showRecentlyViewed?: bool, showSearch?: bool, showHeadings?: bool, maxItems?: int, query?: string, tags?: array<string>, folderId?: int, folderUID?: string} $inputData */
        $data = $inputData;
        return new self(
            keepTime: $data["keepTime"] ?? null,
            includeVars: $data["includeVars"] ?? null,
            showStarred: $data["showStarred"] ?? null,
            showRecentlyViewed: $data["showRecentlyViewed"] ?? null,
            showSearch: $data["showSearch"] ?? null,
            showHeadings: $data["showHeadings"] ?? null,
            maxItems: $data["maxItems"] ?? null,
            query: $data["query"] ?? null,
            tags: $data["tags"] ?? null,
            folderId: $data["folderId"] ?? null,
            folderUID: $data["folderUID"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "keepTime" => $this->keepTime,
            "includeVars" => $this->includeVars,
            "showStarred" => $this->showStarred,
            "showRecentlyViewed" => $this->showRecentlyViewed,
            "showSearch" => $this->showSearch,
            "showHeadings" => $this->showHeadings,
            "maxItems" => $this->maxItems,
            "query" => $this->query,
            "tags" => $this->tags,
        ];
        if (isset($this->folderId)) {
            $data["folderId"] = $this->folderId;
        }
        if (isset($this->folderUID)) {
            $data["folderUID"] = $this->folderUID;
        }
        return $data;
    }
}
