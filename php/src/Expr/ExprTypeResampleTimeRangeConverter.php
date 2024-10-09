<?php

namespace Grafana\Foundation\Expr;

final class ExprTypeResampleTimeRangeConverter
{
    public static function convert(\Grafana\Foundation\Expr\ExprTypeResampleTimeRange $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Expr\ExprTypeResampleTimeRangeBuilder())',
        ];
            if ($input->from !== "" && $input->from !== "now-6h") {
    
        
    $buffer = 'from(';
        $arg0 =\var_export($input->from, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->to !== "" && $input->to !== "now") {
    
        
    $buffer = 'to(';
        $arg0 =\var_export($input->to, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

