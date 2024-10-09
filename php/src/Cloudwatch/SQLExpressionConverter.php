<?php

namespace Grafana\Foundation\Cloudwatch;

final class SQLExpressionConverter
{
    public static function convert(\Grafana\Foundation\Cloudwatch\SQLExpression $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Cloudwatch\SQLExpressionBuilder())',
        ];
            if ($input->select !== null) {
    
        
    $buffer = 'select(';
        $arg0 = \Grafana\Foundation\Cloudwatch\QueryEditorFunctionExpressionConverter::convert($input->select);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->from !== null) {
    
        
    $buffer = 'from(';
        switch (true) {
            case $input->from instanceof \Grafana\Foundation\Cloudwatch\QueryEditorPropertyExpression:
                $disjunctionfrom = \Grafana\Foundation\Cloudwatch\QueryEditorPropertyExpressionConverter::convert($input->from);
                $arg0 = $disjunctionfrom;
                break;
            case $input->from instanceof \Grafana\Foundation\Cloudwatch\QueryEditorFunctionExpression:
                $disjunctionfrom = \Grafana\Foundation\Cloudwatch\QueryEditorFunctionExpressionConverter::convert($input->from);
                $arg0 = $disjunctionfrom;
                break;
            default:
                throw new \ValueError('disjunction branch not handled');
        }
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->where !== null) {
    
        
    $buffer = 'where(';
        $arg0 = \Grafana\Foundation\Cloudwatch\QueryEditorArrayExpressionConverter::convert($input->where);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->groupBy !== null) {
    
        
    $buffer = 'groupBy(';
        $arg0 = \Grafana\Foundation\Cloudwatch\QueryEditorArrayExpressionConverter::convert($input->groupBy);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->orderBy !== null) {
    
        
    $buffer = 'orderBy(';
        $arg0 = \Grafana\Foundation\Cloudwatch\QueryEditorFunctionExpressionConverter::convert($input->orderBy);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->orderByDirection !== null && $input->orderByDirection !== "") {
    
        
    $buffer = 'orderByDirection(';
        $arg0 =\var_export($input->orderByDirection, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->limit !== null) {
    
        
    $buffer = 'limit(';
        $arg0 =\var_export($input->limit, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

