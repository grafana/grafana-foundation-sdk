<?php

namespace Grafana\Foundation\Alertingv1beta1;

final class RuleRefConverter
{
    public static function convert(\Grafana\Foundation\Alertingv1beta1\RuleRef $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Alertingv1beta1\RuleRefBuilder())',
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

