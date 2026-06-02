<?php

namespace Grafana\Foundation\Common;

/**
 * TODO docs
 */
final class TooltipDisplayMode implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, TooltipDisplayMode>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function single(): self
    {
        if (!isset(self::$instances["Single"])) {
            self::$instances["Single"] = new self("single");
        }

        return self::$instances["Single"];
    }

    public static function multi(): self
    {
        if (!isset(self::$instances["Multi"])) {
            self::$instances["Multi"] = new self("multi");
        }

        return self::$instances["Multi"];
    }

    public static function none(): self
    {
        if (!isset(self::$instances["None"])) {
            self::$instances["None"] = new self("none");
        }

        return self::$instances["None"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "single") {
            return self::single();
        }

        if ($value === "multi") {
            return self::multi();
        }

        if ($value === "none") {
            return self::none();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum TooltipDisplayMode");
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

