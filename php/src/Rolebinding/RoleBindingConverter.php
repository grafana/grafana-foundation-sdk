<?php

namespace Grafana\Foundation\Rolebinding;

final class RoleBindingConverter
{
    public static function convert(\Grafana\Foundation\Rolebinding\RoleBinding $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Rolebinding\RoleBindingBuilder())',
        ];
            
    
        {
    $buffer = 'role(';
        switch (true) {
            case $input->role instanceof \Grafana\Foundation\Rolebinding\BuiltinRoleRef:
                $disjunctionrole = \Grafana\Foundation\Rolebinding\BuiltinRoleRefConverter::convert($input->role);
                $arg0 = $disjunctionrole;
                break;
            case $input->role instanceof \Grafana\Foundation\Rolebinding\CustomRoleRef:
                $disjunctionrole = \Grafana\Foundation\Rolebinding\CustomRoleRefConverter::convert($input->role);
                $arg0 = $disjunctionrole;
                break;
            default:
                throw new \ValueError('disjunction branch not handled');
        }
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'subject(';
        $arg0 = \Grafana\Foundation\Rolebinding\RoleBindingSubjectConverter::convert($input->subject);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    

        return \implode("\n\t->", $calls);
    }
}

