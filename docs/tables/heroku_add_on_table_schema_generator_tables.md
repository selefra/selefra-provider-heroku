# Table: heroku_add_on_table_schema_generator_tables

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| created_at | timestamp | X | √ |  | 
| human_name | string | X | √ |  | 
| id | string | √ | √ |  | 
| supports_multiple_installations | bool | X | √ |  | 
| supports_sharing | bool | X | √ |  | 
| cli_plugin_name | string | X | √ |  | 
| name | string | X | √ |  | 
| state | string | X | √ |  | 
| updated_at | timestamp | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


