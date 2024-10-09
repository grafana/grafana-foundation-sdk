<?php

namespace Grafana\Foundation\Accesspolicy;

final class AccessPolicyConverter
{
    public static function convert(\Grafana\Foundation\Accesspolicy\AccessPolicy $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Accesspolicy\AccessPolicyBuilder())',
        ];
            
    
        {
    $buffer = 'scope(';
        $arg0 = \Grafana\Foundation\Accesspolicy\ResourceRefConverter::convert($input->scope);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'role(';
        $arg0 = \Grafana\Foundation\Accesspolicy\RoleRefConverter::convert($input->role);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if (count($input->rules) >= 1) {
    foreach ($input->rules as $item) {
        
    $buffer = 'rules(';
        $arg0 = \Grafana\Foundation\Accesspolicy\AccessRuleConverter::convert($item);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    }
    }

        return \implode("\n\t->", $calls);
    }
}

