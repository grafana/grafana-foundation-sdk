<?php

namespace Grafana\Foundation\Expr;

final class TypeClassicConditionsConverter
{
    public static function convert(\Grafana\Foundation\Expr\TypeClassicConditions $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Expr\TypeClassicConditionsBuilder())',
        ];
            if (count($input->conditions) >= 1) {
    
        
    $buffer = 'conditions(';
        $tmparg0 = [];
        foreach ($input->conditions as $arg1) {
        $tmpconditionsarg1 = \Grafana\Foundation\Expr\ExprTypeClassicConditionsConditionsConverter::convert($arg1);
        $tmparg0[] = $tmpconditionsarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
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
            if ($input->hide !== null) {
    
        
    $buffer = 'hide(';
        $arg0 =\var_export($input->hide, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->intervalMs !== null) {
    
        
    $buffer = 'intervalMs(';
        $arg0 =\var_export($input->intervalMs, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->maxDataPoints !== null) {
    
        
    $buffer = 'maxDataPoints(';
        $arg0 =\var_export($input->maxDataPoints, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->queryType !== null && $input->queryType !== "") {
    
        
    $buffer = 'queryType(';
        $arg0 =\var_export($input->queryType, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->refId !== "") {
    
        
    $buffer = 'refId(';
        $arg0 =\var_export($input->refId, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->resultAssertions !== null) {
    
        
    $buffer = 'resultAssertions(';
        $arg0 = \Grafana\Foundation\Expr\ExprTypeClassicConditionsResultAssertionsConverter::convert($input->resultAssertions);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->timeRange !== null) {
    
        
    $buffer = 'timeRange(';
        $arg0 = \Grafana\Foundation\Expr\ExprTypeClassicConditionsTimeRangeConverter::convert($input->timeRange);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

