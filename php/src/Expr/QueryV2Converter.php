<?php

namespace Grafana\Foundation\Expr;

final class QueryV2Converter
{
    public static function convert(\Grafana\Foundation\Dashboardv2\DataQueryKind $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Expr\QueryV2Builder())',
        ];
            if ($input->version !== "" && $input->version !== "v0") {
    
        
    $buffer = 'version(';
        $arg0 =\var_export($input->version, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->labels !== null) {
    
        
    $buffer = 'labels(';
        $arg0 = "[";
        foreach ($input->labels as $key => $arg1) {
            $tmplabelsarg1 =\var_export($arg1, true);
            $arg0 .= "\t".var_export($key, true)." => $tmplabelsarg1,";
        }
        $arg0 .= "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->datasource !== null) {
    
        
    $buffer = 'datasource(';
        $arg0 = \Grafana\Foundation\Dashboardv2\Dashboardv2DataQueryKindDatasourceConverter::convert($input->datasource);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Expr\TypeMath) {
    
        
    $buffer = 'typeMath(';
        $arg0 = \Grafana\Foundation\Expr\TypeMathConverter::convert($input->spec);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Expr\TypeReduce) {
    
        
    $buffer = 'typeReduce(';
        $arg0 = \Grafana\Foundation\Expr\TypeReduceConverter::convert($input->spec);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Expr\TypeResample) {
    
        
    $buffer = 'typeResample(';
        $arg0 = \Grafana\Foundation\Expr\TypeResampleConverter::convert($input->spec);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Expr\TypeClassicConditions) {
    
        
    $buffer = 'typeClassicConditions(';
        $arg0 = \Grafana\Foundation\Expr\TypeClassicConditionsConverter::convert($input->spec);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Expr\TypeThreshold) {
    
        
    $buffer = 'typeThreshold(';
        $arg0 = \Grafana\Foundation\Expr\TypeThresholdConverter::convert($input->spec);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Expr\TypeSql) {
    
        
    $buffer = 'typeSql(';
        $arg0 = \Grafana\Foundation\Expr\TypeSqlConverter::convert($input->spec);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

