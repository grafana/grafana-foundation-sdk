<?php

namespace Grafana\Foundation\Googlecloudmonitoring;

final class PromQLQueryConverter
{
    public static function convert(\Grafana\Foundation\Googlecloudmonitoring\PromQLQuery $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Googlecloudmonitoring\PromQLQueryBuilder())',
        ];
            if ($input->projectName !== "") {
    
        
    $buffer = 'projectName(';
        $arg0 =\var_export($input->projectName, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->expr !== "") {
    
        
    $buffer = 'expr(';
        $arg0 =\var_export($input->expr, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->step !== "") {
    
        
    $buffer = 'step(';
        $arg0 =\var_export($input->step, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

