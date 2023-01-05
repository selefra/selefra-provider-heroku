# Table: heroku_app_webhooks

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| app | json | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| id | string | √ | √ |  | 
| include | string_array | X | √ |  | 
| level | string | X | √ |  | 
| updated_at | timestamp | X | √ |  | 
| url | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


