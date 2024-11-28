<?php

namespace Grafana\Foundation\Dashboard;

final class DatasourceVariableConverter
{
    public static function convert(\Grafana\Foundation\Dashboard\VariableModel $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboard\DatasourceVariableBuilder('.\var_export($input->name, true).'))',
        ];
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
            if ($input->hide !== null) {
    
        
    $buffer = 'hide(';
        $arg0 ='\Grafana\Foundation\Dashboard\VariableHide::fromValue("'.$input->hide.'")';
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
    
        
    $buffer = 'type(';
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
            if ($input->current !== null) {
    
        
    $buffer = 'current(';
        $arg0 ='(new \Grafana\Foundation\Dashboard\VariableOption('.(($input->current->selected !== null) ? 'selected: '.\var_export($input->current->selected, true).', ' : '').'text: '.\var_export($input->current->text, true).',value: '.\var_export($input->current->value, true).',))';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->multi !== null && $input->multi !== false) {
    
        
    $buffer = 'multi(';
        $arg0 =\var_export($input->multi, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->allowCustomValue !== null && $input->allowCustomValue !== true) {
    
        
    $buffer = 'allowCustomValue(';
        $arg0 =\var_export($input->allowCustomValue, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->includeAll !== null && $input->includeAll !== false) {
    
        
    $buffer = 'includeAll(';
        $arg0 =\var_export($input->includeAll, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->allValue !== null && $input->allValue !== "") {
    
        
    $buffer = 'allValue(';
        $arg0 =\var_export($input->allValue, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->regex !== null && $input->regex !== "") {
    
        
    $buffer = 'regex(';
        $arg0 =\var_export($input->regex, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

