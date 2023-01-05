# Table: heroku_pipeline_deployments

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| slug | json | X | √ |  | 
| status | string | X | √ |  | 
| version | int | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| current | bool | X | √ |  | 
| description | string | X | √ |  | 
| id | string | √ | √ |  | 
| output_stream_url | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| addon_plan_names | string_array | X | √ |  | 
| app | json | X | √ |  | 
| updated_at | timestamp | X | √ |  | 
| user | json | X | √ |  | 


