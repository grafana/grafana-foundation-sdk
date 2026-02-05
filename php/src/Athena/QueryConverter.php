<?php

namespace Grafana\Foundation\Athena;

final class QueryConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\DataQueryKind $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Athena\QueryBuilder())',
        ];
            if ($input->version !== "" && $input->version !== "v0") {
    
        
    $buffer = 'version(';
        $arg0 =\var_export($input->version, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->datasource !== null) {
    
        
    $buffer = 'datasource(';
        $arg0 = \Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1DataQueryKindDatasourceConverter::convert($input->datasource);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Athena\Dataquery) {
    
        
    $buffer = 'format(';
        $arg0 ='\Grafana\Foundation\Athena\FormatOptions::fromValue("'.$input->spec->format.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Athena\Dataquery) {
    
        
    $buffer = 'connectionArgs(';
        $arg0 = \Grafana\Foundation\Athena\ConnectionArgsConverter::convert($input->spec->connectionArgs);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Athena\Dataquery && $input->spec->table !== null && $input->spec->table !== "") {
    
        
    $buffer = 'table(';
        $arg0 =\var_export($input->spec->table, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Athena\Dataquery && $input->spec->column !== null && $input->spec->column !== "") {
    
        
    $buffer = 'column(';
        $arg0 =\var_export($input->spec->column, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Athena\Dataquery && $input->spec->queryID !== null && $input->spec->queryID !== "") {
    
        
    $buffer = 'queryID(';
        $arg0 =\var_export($input->spec->queryID, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Athena\Dataquery && $input->spec->refId !== null && $input->spec->refId !== "") {
    
        
    $buffer = 'refId(';
        $arg0 =\var_export($input->spec->refId, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Athena\Dataquery && $input->spec->hide !== null) {
    
        
    $buffer = 'hide(';
        $arg0 =\var_export($input->spec->hide, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Athena\Dataquery && $input->spec->queryType !== null && $input->spec->queryType !== "") {
    
        
    $buffer = 'queryType(';
        $arg0 =\var_export($input->spec->queryType, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Athena\Dataquery && $input->spec->rawSQL !== "") {
    
        
    $buffer = 'rawSQL(';
        $arg0 =\var_export($input->spec->rawSQL, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

