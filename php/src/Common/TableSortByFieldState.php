<?php

namespace Grafana\Foundation\Common;

/**
 * Sort by field state
 */
class TableSortByFieldState implements \JsonSerializable
{
    /**
     * Sets the display name of the field to sort by
     */
    public string $displayName;

    /**
     * Flag used to indicate descending sort order
     */
    public ?bool $desc;

    /**
     * @param string|null $displayName
     * @param bool|null $desc
     */
    public function __construct(?string $displayName = null, ?bool $desc = null)
    {
        $this->displayName = $displayName ?: "";
        $this->desc = $desc;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{displayName?: string, desc?: bool} $inputData */
        $data = $inputData;
        return new self(
            displayName: $data["displayName"] ?? null,
            desc: $data["desc"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->displayName = $this->displayName;
        if (isset($this->desc)) {
            $data->desc = $this->desc;
        }
        return $data;
    }
}
