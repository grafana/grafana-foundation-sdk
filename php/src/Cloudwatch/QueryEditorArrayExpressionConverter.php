<?php

namespace Grafana\Foundation\Cloudwatch;

final class QueryEditorArrayExpressionConverter
{
    public static function convert(\Grafana\Foundation\Cloudwatch\QueryEditorArrayExpression $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Cloudwatch\QueryEditorArrayExpressionBuilder())',
        ];
            
    
        {
    $buffer = 'type(';
        $arg0 ='\Grafana\Foundation\Cloudwatch\QueryEditorArrayExpressionType::fromValue("'.$input->type.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if (count($input->expressions) >= 1) {
    
        
    $buffer = 'expressions(';
        $tmparg0 = [];
        foreach ($input->expressions as $arg1) {
        switch (true) {
            case $arg1 instanceof \Grafana\Foundation\Cloudwatch\QueryEditorArrayExpression:
                $disjunctionarg1 = \Grafana\Foundation\Cloudwatch\QueryEditorArrayExpressionConverter::convert($arg1);
                $tmpexpressionsarg1 = $disjunctionarg1;
                break;
            case $arg1 instanceof \Grafana\Foundation\Cloudwatch\QueryEditorPropertyExpression:
                $disjunctionarg1 = \Grafana\Foundation\Cloudwatch\QueryEditorPropertyExpressionConverter::convert($arg1);
                $tmpexpressionsarg1 = $disjunctionarg1;
                break;
            case $arg1 instanceof \Grafana\Foundation\Cloudwatch\QueryEditorGroupByExpression:
                $disjunctionarg1 = \Grafana\Foundation\Cloudwatch\QueryEditorGroupByExpressionConverter::convert($arg1);
                $tmpexpressionsarg1 = $disjunctionarg1;
                break;
            case $arg1 instanceof \Grafana\Foundation\Cloudwatch\QueryEditorFunctionExpression:
                $disjunctionarg1 = \Grafana\Foundation\Cloudwatch\QueryEditorFunctionExpressionConverter::convert($arg1);
                $tmpexpressionsarg1 = $disjunctionarg1;
                break;
            case $arg1 instanceof \Grafana\Foundation\Cloudwatch\QueryEditorFunctionParameterExpression:
                $disjunctionarg1 = \Grafana\Foundation\Cloudwatch\QueryEditorFunctionParameterExpressionConverter::convert($arg1);
                $tmpexpressionsarg1 = $disjunctionarg1;
                break;
            case $arg1 instanceof \Grafana\Foundation\Cloudwatch\QueryEditorOperatorExpression:
                $disjunctionarg1 = \Grafana\Foundation\Cloudwatch\QueryEditorOperatorExpressionConverter::convert($arg1);
                $tmpexpressionsarg1 = $disjunctionarg1;
                break;
            default:
                throw new \ValueError('disjunction branch not handled');
        }
        $tmparg0[] = $tmpexpressionsarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

