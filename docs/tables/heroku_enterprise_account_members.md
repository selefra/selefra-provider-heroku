# Table: heroku_enterprise_account_members

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| enterprise_account | json | X | √ |  | 
| id | string | √ | √ |  | 
| identity_provider | json | X | √ |  | 
| permissions | json | X | √ |  | 
| two_factor_authentication | bool | X | √ |  | 
| user | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


