<?php

namespace Grafana\Foundation\Expr;

final class ExprTypeClassicConditionsConditionsReducerConverter
{
    public static function convert(\Grafana\Foundation\Expr\ExprTypeClassicConditionsConditionsReducer $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Expr\ExprTypeClassicConditionsConditionsReducerBuilder())',
        ];
            if ($input->type !== "") {
    
        
    $buffer = 'type(';
        $arg0 =\var_export($input->type, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

