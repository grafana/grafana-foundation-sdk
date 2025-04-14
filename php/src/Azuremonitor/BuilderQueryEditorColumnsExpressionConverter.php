<?php

namespace Grafana\Foundation\Azuremonitor;

final class BuilderQueryEditorColumnsExpressionConverter
{
    public static function convert(\Grafana\Foundation\Azuremonitor\BuilderQueryEditorColumnsExpression $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Azuremonitor\BuilderQueryEditorColumnsExpressionBuilder())',
        ];
            if ($input->columns !== null && count($input->columns) >= 1) {
    
        
    $buffer = 'columns(';
        $tmparg0 = [];
        foreach ($input->columns as $arg1) {
        $tmpcolumnsarg1 =\var_export($arg1, true);
        $tmparg0[] = $tmpcolumnsarg1;
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

