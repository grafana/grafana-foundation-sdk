<?php

namespace Grafana\Foundation\Expr;

final class ExprTypeReduceSettingsConverter
{
    public static function convert(\Grafana\Foundation\Expr\ExprTypeReduceSettings $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Expr\ExprTypeReduceSettingsBuilder())',
        ];
            
    
        {
    $buffer = 'mode(';
        $arg0 ='\Grafana\Foundation\Expr\TypeReduceMode::fromValue("'.$input->mode.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->replaceWithValue !== null) {
    
        
    $buffer = 'replaceWithValue(';
        $arg0 =\var_export($input->replaceWithValue, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

