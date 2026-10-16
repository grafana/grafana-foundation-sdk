<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class NotebookLayoutItemSpecSource implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, NotebookLayoutItemSpecSource>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function assistant(): self
    {
        if (!isset(self::$instances["Assistant"])) {
            self::$instances["Assistant"] = new self("assistant");
        }

        return self::$instances["Assistant"];
    }

    public static function user(): self
    {
        if (!isset(self::$instances["User"])) {
            self::$instances["User"] = new self("user");
        }

        return self::$instances["User"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "assistant") {
            return self::assistant();
        }

        if ($value === "user") {
            return self::user();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum NotebookLayoutItemSpecSource");
    }

    public function jsonSerialize(): string
    {
        return $this->value;
    }

    public function __toString(): string
    {
        return $this->value;
    }
}

