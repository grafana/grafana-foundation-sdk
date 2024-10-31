---
title: <span class="badge object-type-enum"></span> RoleBindingSubjectKind
---
# <span class="badge object-type-enum"></span> RoleBindingSubjectKind

## Definition

```php
final class RoleBindingSubjectKind implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, RoleBindingSubjectKind>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
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
        if ($value === "Team") {
            return self::team();
        }

        if ($value === "User") {
            return self::user();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum RoleBindingSubjectKind");
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
