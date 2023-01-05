# Table: heroku_enterprise_accounts

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| updated_at | timestamp | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| created_at | timestamp | X | √ |  | 
| id | string | √ | √ |  | 
| identity_provider | json | X | √ |  | 
| name | string | X | √ |  | 
| permissions | string_array | X | √ |  | 
| trial | bool | X | √ |  | 


