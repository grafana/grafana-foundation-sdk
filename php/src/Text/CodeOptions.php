<?php

namespace Grafana\Foundation\Text;

class CodeOptions implements \JsonSerializable
{
    /**
     * The language passed to monaco code editor
     */
    public \Grafana\Foundation\Text\CodeLanguage $language;

    public bool $showLineNumbers;

    public bool $showMiniMap;

    /**
     * @param \Grafana\Foundation\Text\CodeLanguage|null $language
     * @param bool|null $showLineNumbers
     * @param bool|null $showMiniMap
     */
    public function __construct(?\Grafana\Foundation\Text\CodeLanguage $language = null, ?bool $showLineNumbers = null, ?bool $showMiniMap = null)
    {
        $this->language = $language ?: \Grafana\Foundation\Text\CodeLanguage::plaintext();
        $this->showLineNumbers = $showLineNumbers ?: false;
        $this->showMiniMap = $showMiniMap ?: false;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{language?: string, showLineNumbers?: bool, showMiniMap?: bool} $inputData */
        $data = $inputData;
        return new self(
            language: isset($data["language"]) ? (function($input) { return \Grafana\Foundation\Text\CodeLanguage::fromValue($input); })($data["language"]) : null,
            showLineNumbers: $data["showLineNumbers"] ?? null,
            showMiniMap: $data["showMiniMap"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->language = $this->language;
        $data->showLineNumbers = $this->showLineNumbers;
        $data->showMiniMap = $this->showMiniMap;
        return $data;
    }
}
