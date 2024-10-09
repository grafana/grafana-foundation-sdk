<?php

namespace Grafana\Foundation\Role;

final class RoleConverter
{
    public static function convert(\Grafana\Foundation\Role\Role $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Role\RoleBuilder())',
        ];
            if ($input->name !== "") {
    
        
    $buffer = 'name(';
        $arg0 =\var_export($input->name, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->displayName !== null && $input->displayName !== "") {
    
        
    $buffer = 'displayName(';
        $arg0 =\var_export($input->displayName, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->groupName !== null && $input->groupName !== "") {
    
        
    $buffer = 'groupName(';
        $arg0 =\var_export($input->groupName, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->description !== null && $input->description !== "") {
    
        
    $buffer = 'description(';
        $arg0 =\var_export($input->description, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->hidden !== false) {
    
        
    $buffer = 'hidden(';
        $arg0 =\var_export($input->hidden, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

