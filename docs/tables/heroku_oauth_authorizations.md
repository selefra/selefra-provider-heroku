# Table: heroku_oauth_authorizations

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| access_token | json | X | √ |  | 
| id | string | √ | √ |  | 
| scope | string_array | X | √ |  | 
| user | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| client | json | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| grant | json | X | √ |  | 
| refresh_token | json | X | √ |  | 
| updated_at | timestamp | X | √ |  | 


