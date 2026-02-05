<?php

namespace Grafana\Foundation\Expr;

final class ExprTypeClassicConditionsConditionsQueryConverter
{
    public static function convert(\Grafana\Foundation\Expr\ExprTypeClassicConditionsConditionsQuery $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Expr\ExprTypeClassicConditionsConditionsQueryBuilder())',
        ];
            if (count($input->params) >= 1) {
    
        
    $buffer = 'params(';
        $tmparg0 = [];
        foreach ($input->params as $arg1) {
        $tmpparamsarg1 =\var_export($arg1, true);
        $tmparg0[] = $tmpparamsarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

