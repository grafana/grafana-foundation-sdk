<?php

namespace Grafana\Foundation\Common;

/**
 * TODO docs
 */
final class VisibilityMode implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, VisibilityMode>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function auto(): self
    {
        if (!isset(self::$instances["Auto"])) {
            self::$instances["Auto"] = new self("auto");
        }

        return self::$instances["Auto"];
    }

    public static function never(): self
    {
        if (!isset(self::$instances["Never"])) {
            self::$instances["Never"] = new self("never");
        }

        return self::$instances["Never"];
    }

    public static function always(): self
    {
        if (!isset(self::$instances["Always"])) {
            self::$instances["Always"] = new self("always");
        }

        return self::$instances["Always"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "auto") {
            return self::auto();
        }

        if ($value === "never") {
            return self::never();
        }

        if ($value === "always") {
            return self::always();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum VisibilityMode");
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

