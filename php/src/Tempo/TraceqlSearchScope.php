<?php

namespace Grafana\Foundation\Tempo;

/**
 * static fields are pre-set in the UI, dynamic fields are added by the user
 */
final class TraceqlSearchScope implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, TraceqlSearchScope>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function intrinsic(): self
    {
        if (!isset(self::$instances["Intrinsic"])) {
            self::$instances["Intrinsic"] = new self("intrinsic");
        }

        return self::$instances["Intrinsic"];
    }

    public static function unscoped(): self
    {
        if (!isset(self::$instances["Unscoped"])) {
            self::$instances["Unscoped"] = new self("unscoped");
        }

        return self::$instances["Unscoped"];
    }

    public static function event(): self
    {
        if (!isset(self::$instances["Event"])) {
            self::$instances["Event"] = new self("event");
        }

        return self::$instances["Event"];
    }

    public static function instrumentation(): self
    {
        if (!isset(self::$instances["Instrumentation"])) {
            self::$instances["Instrumentation"] = new self("instrumentation");
        }

        return self::$instances["Instrumentation"];
    }

    public static function link(): self
    {
        if (!isset(self::$instances["Link"])) {
            self::$instances["Link"] = new self("link");
        }

        return self::$instances["Link"];
    }

    public static function resource(): self
    {
        if (!isset(self::$instances["Resource"])) {
            self::$instances["Resource"] = new self("resource");
        }

        return self::$instances["Resource"];
    }

    public static function span(): self
    {
        if (!isset(self::$instances["Span"])) {
            self::$instances["Span"] = new self("span");
        }

        return self::$instances["Span"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "intrinsic") {
            return self::intrinsic();
        }

        if ($value === "unscoped") {
            return self::unscoped();
        }

        if ($value === "event") {
            return self::event();
        }

        if ($value === "instrumentation") {
            return self::instrumentation();
        }

        if ($value === "link") {
            return self::link();
        }

        if ($value === "resource") {
            return self::resource();
        }

        if ($value === "span") {
            return self::span();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum TraceqlSearchScope");
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

