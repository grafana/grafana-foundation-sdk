<?php

namespace Grafana\Foundation\Alertingv0alpha1;

final class RuleRefConverter
{
    public static function convert(\Grafana\Foundation\Alertingv0alpha1\RuleRef $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Alertingv0alpha1\RuleRefBuilder())',
        ];
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

