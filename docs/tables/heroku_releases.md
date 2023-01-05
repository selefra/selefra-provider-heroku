# Table: heroku_releases

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| addon_plan_names | string_array | X | √ |  | 
| app | json | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| current | bool | X | √ |  | 
| description | string | X | √ |  | 
| output_stream_url | string | X | √ |  | 
| slug | json | X | √ |  | 
| status | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| id | string | √ | √ |  | 
| updated_at | timestamp | X | √ |  | 
| user | json | X | √ |  | 
| version | int | X | √ |  | 


