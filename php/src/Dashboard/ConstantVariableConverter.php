<?php

namespace Grafana\Foundation\Dashboard;

final class ConstantVariableConverter
{
    public static function convert(\Grafana\Foundation\Dashboard\VariableModel $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboard\ConstantVariableBuilder('.\var_export($input->name, true).'))',
        ];
            if ($input->id !== "" && $input->id !== "00000000-0000-0000-0000-000000000000") {
    
        
    $buffer = 'id(';
        $arg0 =\var_export($input->id, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->name !== "") {
    
        
    $buffer = 'name(';
        $arg0 =\var_export($input->name, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->label !== null && $input->label !== "") {
    
        
    $buffer = 'label(';
        $arg0 =\var_export($input->label, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->description !== null && $input->description !== "") {
    
        
    $buffer = 'description(';
        $arg0 =\var_export($input->description, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->query !== null) {
    
        
    $buffer = 'value(';
        switch (true) {
            case is_string($input->query):
                $disjunctionquery =\var_export($input->query, true);
                $arg0 = $disjunctionquery;
                break;
            case is_array($input->query):
                $disjunctionquery = "[";
        foreach ($input->query as $key => $arg1) {
            $tmpqueryarg1 =\var_export($arg1, true);
            $disjunctionquery .= "\t".var_export($key, true)." => $tmpqueryarg1,";
        }
        $disjunctionquery .= "]";
                $arg0 = $disjunctionquery;
                break;
            default:
                throw new \ValueError('disjunction branch not handled');
        }
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->allFormat !== null && $input->allFormat !== "") {
    
        
    $buffer = 'allFormat(';
        $arg0 =\var_export($input->allFormat, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

