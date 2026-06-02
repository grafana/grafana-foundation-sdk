<?php

namespace Grafana\Foundation\Common;

/**
 * TODO docs
 * Note: "hidden" needs to remain as an option for plugins compatibility
 */
final class LegendDisplayMode implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, LegendDisplayMode>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function list(): self
    {
        if (!isset(self::$instances["List"])) {
            self::$instances["List"] = new self("list");
        }

        return self::$instances["List"];
    }

    public static function table(): self
    {
        if (!isset(self::$instances["Table"])) {
            self::$instances["Table"] = new self("table");
        }

        return self::$instances["Table"];
    }

    public static function hidden(): self
    {
        if (!isset(self::$instances["Hidden"])) {
            self::$instances["Hidden"] = new self("hidden");
        }

        return self::$instances["Hidden"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "list") {
            return self::list();
        }

        if ($value === "table") {
            return self::table();
        }

        if ($value === "hidden") {
            return self::hidden();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum LegendDisplayMode");
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

