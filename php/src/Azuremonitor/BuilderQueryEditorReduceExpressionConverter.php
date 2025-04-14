<?php

namespace Grafana\Foundation\Azuremonitor;

final class BuilderQueryEditorReduceExpressionConverter
{
    public static function convert(\Grafana\Foundation\Azuremonitor\BuilderQueryEditorReduceExpression $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Azuremonitor\BuilderQueryEditorReduceExpressionBuilder())',
        ];
            if ($input->property !== null) {
    
        
    $buffer = 'property(';
        $arg0 = \Grafana\Foundation\Azuremonitor\BuilderQueryEditorPropertyConverter::convert($input->property);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->reduce !== null) {
    
        
    $buffer = 'reduce(';
        $arg0 = \Grafana\Foundation\Azuremonitor\BuilderQueryEditorPropertyConverter::convert($input->reduce);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->parameters !== null && count($input->parameters) >= 1) {
    
        
    $buffer = 'parameters(';
        $tmparg0 = [];
        foreach ($input->parameters as $arg1) {
        $tmpparametersarg1 = \Grafana\Foundation\Azuremonitor\BuilderQueryEditorFunctionParameterExpressionConverter::convert($arg1);
        $tmparg0[] = $tmpparametersarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->focus !== null) {
    
        
    $buffer = 'focus(';
        $arg0 =\var_export($input->focus, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

