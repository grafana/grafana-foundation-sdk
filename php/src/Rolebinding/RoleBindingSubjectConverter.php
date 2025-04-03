<?php

namespace Grafana\Foundation\Rolebinding;

final class RoleBindingSubjectConverter
{
    public static function convert(\Grafana\Foundation\Rolebinding\RoleBindingSubject $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Rolebinding\RoleBindingSubjectBuilder())',
        ];
            
    
        {
    $buffer = 'kind(';
        $arg0 ='\Grafana\Foundation\Rolebinding\RoleBindingSubjectKind::fromValue("'.$input->kind.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->name !== "") {
    
        
    $buffer = 'name(';
        $arg0 =\var_export($input->name, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

