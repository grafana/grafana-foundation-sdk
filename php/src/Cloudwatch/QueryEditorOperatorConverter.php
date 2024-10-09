<?php

namespace Grafana\Foundation\Cloudwatch;

final class QueryEditorOperatorConverter
{
    public static function convert(\Grafana\Foundation\Cloudwatch\QueryEditorOperator $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Cloudwatch\QueryEditorOperatorBuilder())',
        ];
            if ($input->name !== null && $input->name !== "") {
    
        
    $buffer = 'name(';
        $arg0 =\var_export($input->name, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->value !== null) {
    
        
    $buffer = 'value(';
        switch (true) {
            case is_string($input->value):
                $disjunctionvalue =\var_export($input->value, true);
                $arg0 = $disjunctionvalue;
                break;
            case is_bool($input->value):
                $disjunctionvalue =\var_export($input->value, true);
                $arg0 = $disjunctionvalue;
                break;
            case is_int($input->value):
                $disjunctionvalue =\var_export($input->value, true);
                $arg0 = $disjunctionvalue;
                break;
            case is_array($input->value):
                $tmpdisjunctionvalue = [];
        foreach ($input->value as $arg1) {
        switch (true) {
            case is_string($arg1):
                $disjunctionarg1 =\var_export($arg1, true);
                $tmpvaluearg1 = $disjunctionarg1;
                break;
            case is_bool($arg1):
                $disjunctionarg1 =\var_export($arg1, true);
                $tmpvaluearg1 = $disjunctionarg1;
                break;
            case is_int($arg1):
                $disjunctionarg1 =\var_export($arg1, true);
                $tmpvaluearg1 = $disjunctionarg1;
                break;
            default:
                throw new \ValueError('disjunction branch not handled');
        }
        $tmpdisjunctionvalue[] = $tmpvaluearg1;
        }
        $disjunctionvalue = "[" . implode(", \n", $tmpdisjunctionvalue) . "]";
                $arg0 = $disjunctionvalue;
                break;
            default:
                throw new \ValueError('disjunction branch not handled');
        }
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

