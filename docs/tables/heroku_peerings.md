# Table: heroku_peerings

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| aws_account_id | string | X | √ |  | 
| aws_region | string | X | √ |  | 
| aws_vpc_id | string | X | √ |  | 
| status | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| cidr_blocks | string_array | X | √ |  | 
| expires | timestamp | X | √ |  | 
| pcx_id | string | X | √ |  | 
| type | string | X | √ |  | 


