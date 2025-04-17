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

    public bool $showFolderNames;

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
     * @param bool|null $showFolderNames
     * @param int|null $maxItems
     * @param string|null $query
     * @param array<string>|null $tags
     * @param int|null $folderId
     * @param string|null $folderUID
     */
    public function __construct(?bool $keepTime = null, ?bool $includeVars = null, ?bool $showStarred = null, ?bool $showRecentlyViewed = null, ?bool $showSearch = null, ?bool $showHeadings = null, ?bool $showFolderNames = null, ?int $maxItems = null, ?string $query = null, ?array $tags = null, ?int $folderId = null, ?string $folderUID = null)
    {
        $this->keepTime = $keepTime ?: false;
        $this->includeVars = $includeVars ?: false;
        $this->showStarred = $showStarred ?: true;
        $this->showRecentlyViewed = $showRecentlyViewed ?: false;
        $this->showSearch = $showSearch ?: false;
        $this->showHeadings = $showHeadings ?: true;
        $this->showFolderNames = $showFolderNames ?: true;
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
        /** @var array{keepTime?: bool, includeVars?: bool, showStarred?: bool, showRecentlyViewed?: bool, showSearch?: bool, showHeadings?: bool, showFolderNames?: bool, maxItems?: int, query?: string, tags?: array<string>, folderId?: int, folderUID?: string} $inputData */
        $data = $inputData;
        return new self(
            keepTime: $data["keepTime"] ?? null,
            includeVars: $data["includeVars"] ?? null,
            showStarred: $data["showStarred"] ?? null,
            showRecentlyViewed: $data["showRecentlyViewed"] ?? null,
            showSearch: $data["showSearch"] ?? null,
            showHeadings: $data["showHeadings"] ?? null,
            showFolderNames: $data["showFolderNames"] ?? null,
            maxItems: $data["maxItems"] ?? null,
            query: $data["query"] ?? null,
            tags: $data["tags"] ?? null,
            folderId: $data["folderId"] ?? null,
            folderUID: $data["folderUID"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->keepTime = $this->keepTime;
        $data->includeVars = $this->includeVars;
        $data->showStarred = $this->showStarred;
        $data->showRecentlyViewed = $this->showRecentlyViewed;
        $data->showSearch = $this->showSearch;
        $data->showHeadings = $this->showHeadings;
        $data->showFolderNames = $this->showFolderNames;
        $data->maxItems = $this->maxItems;
        $data->query = $this->query;
        $data->tags = $this->tags;
        if (isset($this->folderId)) {
            $data->folderId = $this->folderId;
        }
        if (isset($this->folderUID)) {
            $data->folderUID = $this->folderUID;
        }
        return $data;
    }
}
