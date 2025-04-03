<?php

namespace Grafana\Foundation\Alerting;

final class RecordRuleConverter
{
    public static function convert(\Grafana\Foundation\Alerting\RecordRule $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Alerting\RecordRuleBuilder())',
        ];
            if ($input->from !== "") {
    
        
    $buffer = 'from(';
        $arg0 =\var_export($input->from, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->metric !== "") {
    
        
    $buffer = 'metric(';
        $arg0 =\var_export($input->metric, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

