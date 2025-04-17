<?php

namespace Grafana\Foundation\Folder;

/**
 * TODO:
 * common metadata will soon support setting the parent folder in the metadata
 */
class Folder implements \JsonSerializable
{
    /**
     * Unique folder id. (will be k8s name)
     */
    public string $uid;

    /**
     * Folder title
     */
    public string $title;

    /**
     * Description of the folder.
     */
    public ?string $description;

    /**
     * @param string|null $uid
     * @param string|null $title
     * @param string|null $description
     */
    public function __construct(?string $uid = null, ?string $title = null, ?string $description = null)
    {
        $this->uid = $uid ?: "";
        $this->title = $title ?: "";
        $this->description = $description;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{uid?: string, title?: string, description?: string} $inputData */
        $data = $inputData;
        return new self(
            uid: $data["uid"] ?? null,
            title: $data["title"] ?? null,
            description: $data["description"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->uid = $this->uid;
        $data->title = $this->title;
        if (isset($this->description)) {
            $data->description = $this->description;
        }
        return $data;
    }
}
