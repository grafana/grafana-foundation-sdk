<?php

namespace Grafana\Foundation\Azuremonitor;

final class BuilderQueryEditorWhereExpressionArrayConverter
{
    public static function convert(\Grafana\Foundation\Azuremonitor\BuilderQueryEditorWhereExpressionArray $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Azuremonitor\BuilderQueryEditorWhereExpressionArrayBuilder())',
        ];
            if (count($input->expressions) >= 1) {
    
        
    $buffer = 'expressions(';
        $tmparg0 = [];
        foreach ($input->expressions as $arg1) {
        $tmpexpressionsarg1 = \Grafana\Foundation\Azuremonitor\BuilderQueryEditorWhereExpressionConverter::convert($arg1);
        $tmparg0[] = $tmpexpressionsarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'type(';
        $arg0 ='\Grafana\Foundation\Azuremonitor\BuilderQueryEditorExpressionType::fromValue("'.$input->type.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    

        return \implode("\n\t->", $calls);
    }
}

