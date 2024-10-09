<?php

namespace Grafana\Foundation\Accesspolicy;

final class AccessRuleConverter
{
    public static function convert(\Grafana\Foundation\Accesspolicy\AccessRule $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Accesspolicy\AccessRuleBuilder())',
        ];
            if ($input->kind !== "" && $input->kind !== "*") {
    
        
    $buffer = 'kind(';
        $arg0 =\var_export($input->kind, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'verb(';
        switch (true) {
            case is_string($input->verb):
                $disjunctionverb =\var_export($input->verb, true);
                $arg0 = $disjunctionverb;
                break;
            default:
                throw new \ValueError('disjunction branch not handled');
        }
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->target !== null && $input->target !== "") {
    
        
    $buffer = 'target(';
        $arg0 =\var_export($input->target, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

