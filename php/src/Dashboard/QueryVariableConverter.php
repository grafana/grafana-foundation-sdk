<?php

namespace Grafana\Foundation\Dashboard;

final class QueryVariableConverter
{
    public static function convert(\Grafana\Foundation\Dashboard\VariableModel $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboard\QueryVariableBuilder('.\var_export($input->name, true).'))',
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
    
        
    $buffer = 'query(';
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
            if ($input->datasource !== null) {
    
        
    $buffer = 'datasource(';
        $arg0 ='(new \Grafana\Foundation\Dashboard\DataSourceRef('.(($input->datasource->type !== null) ? 'type: '.\var_export($input->datasource->type, true).', ' : '').''.(($input->datasource->uid !== null) ? 'uid: '.\var_export($input->datasource->uid, true).', ' : '').'))';
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
            if ($input->options !== null && count($input->options) >= 1) {
    
        
    $buffer = 'options(';
        $tmparg0 = [];
        foreach ($input->options as $arg1) {
        $tmpoptionsarg1 ='(new \Grafana\Foundation\Dashboard\VariableOption('.(($arg1->selected !== null) ? 'selected: '.\var_export($arg1->selected, true).', ' : '').'text: '.\var_export($arg1->text, true).',value: '.\var_export($arg1->value, true).',))';
        $tmparg0[] = $tmpoptionsarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->refresh !== null) {
    
        
    $buffer = 'refresh(';
        $arg0 ='\Grafana\Foundation\Dashboard\VariableRefresh::fromValue("'.$input->refresh.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->sort !== null) {
    
        
    $buffer = 'sort(';
        $arg0 ='\Grafana\Foundation\Dashboard\VariableSort::fromValue("'.$input->sort.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

