---
title: <span class="badge object-type-enum"></span> RoleRefKind
---
# <span class="badge object-type-enum"></span> RoleRefKind

## Definition

```php
final class RoleRefKind implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, RoleRefKind>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function role(): self
    {
        if (!isset(self::$instances["Role"])) {
            self::$instances["Role"] = new self("Role");
        }

        return self::$instances["Role"];
    }

    public static function builtinRole(): self
    {
        if (!isset(self::$instances["BuiltinRole"])) {
            self::$instances["BuiltinRole"] = new self("BuiltinRole");
        }

        return self::$instances["BuiltinRole"];
    }

    public static function team(): self
    {
        if (!isset(self::$instances["Team"])) {
            self::$instances["Team"] = new self("Team");
        }

        return self::$instances["Team"];
    }

    public static function user(): self
    {
        if (!isset(self::$instances["User"])) {
            self::$instances["User"] = new self("User");
        }

        return self::$instances["User"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "Role") {
            return self::role();
        }

        if ($value === "BuiltinRole") {
            return self::builtinRole();
        }

        if ($value === "Team") {
            return self::team();
        }

        if ($value === "User") {
            return self::user();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum RoleRefKind");
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

```
