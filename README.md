
# b2b-reputation-system
基于区块链的电子元器件数据发布及交易平台，并通过交易数据来实现买家卖家信誉机制。

#1.简介：
该智能合约实现一个简单的交易平台，然后发展成一个信誉系统。
买家可以在平台发布一些求购需求，包括需要的器件以及数量。卖家通过平台可以看到这些需求，并且对这些需求给出自己的报价，多个卖家可以对同一个需求进行报价，最后由买家在收到的报价中选择并确认最后的交易。
通过对接bom2buy，可以打通用户、器件、卖家数据，进行更好的数据利用和分析。

#2.主要实现如下的功能：
## Phase 1 (Chaincode)
* 买家发布需求
* 卖家报价
* 买家确认交易
* 获取买家交易记录
* 获取卖家交易记录
* 获取所有交易记录
* 获取某个器件的交易记录
* 新增买家
* 新增卖家
* TBD

## Phase 2 (Dashboard & Bom2buy)
* 对接Bom2buy用户、器件等数据
* 相似型号报价
* 后台界面
* TBD

## Phase 3 (Enhancement)
* 交易信息私密化
* 买家信誉评分系统
* 卖家信誉评分系统
* TBD

#3.主要函数
```
init：初始化
invoke：调用合约内部的函数
query：查询相关的信息
createUser：新增买家
createDistributor：新增分销商(卖家)

publishInquiry：用户发布采购需求（交易行为）
offerInquiry：分销商报价（交易行为）
finalizeInquiry：用户接受报价（交易行为）

getUsers：获取所有的用户信息
getDistributors：获取所有的分销商信息
getTransactions：获取所有的交易记录
getUserById：获取某个用户信息
getDistributorById：获取某个分销商信息
getTransactionsByPartNumber：获取某器件的交易记录
writeUser：修改用户信息
writeDistributor：修改分销商信息
writeTransaction：写入交易信息
```
#4.数据结构设计 (Phase 1)
```
User 用户/买家

ID：编号
Name：名称

```
```
Distributor 分销商/卖家

ID：编号
Name：名称

```
```
Inquiry 询价
ID：交易 ID
InquiryID：询价ID，卖家给出的Offer，以及最后的交易确认都基于这个ID。
PartNumber：器件号
QtyNeeded: 需要的数量
QtyOffered: 卖家给出的数量
UnitPriceOffered: 卖家给出的价格
FromType：发送方角色 // User:0, Distributor:1
FromID：发送方 ID
ToType：接收方角色 // User:0, Distributor:1, Public:2
ToID：接收方 ID // When ToType = 2, ToID should be empty
Time：交易时间
```
#5.接口设计
```
init
request:
args[0] Version
args[1] 1.0
response:
{"AppName":"b2b-reputation-system","Version":"1.0"}
```
```
createUser
request:
args[0] 用户名
response:
{"Name":"XXX","ID":"XX"}
```
```
createDistributor
request:
args[0] 分销商名
response:
{"Name":"XXX","ID":"XX"}
```
```
publishInquiry
request 参数:
args[0] PartNumber：器件号
args[1] QtyNeeded: 需要的数量
args[2] FromType：发送方角色 // User:0, Distributor:1
args[3] FromID：发送方 ID
response:

{"Status":"SUCCESS","ID":"1","InquiryID":"1","TrasactionID":"123","Time":"2018-10-16T02:00:00Z"}
```
```
offerInquiry
request:
args[0] InquiryID
args[1] QtyOffered
args[2] UnitPriceOffered
args[3] FromType
args[4] FromID
response:
{"Status":"SUCCESS","ID":"1","InquiryID":"1","TrasactionID":"123","OfferID":"123","Time":"2018-10-16T02:00:00Z"}
```
```
finalizeInquiry
request:
args[0] OfferID
args[1] FromType
args[2] FromID
response:
{"Status":"SUCCESS","ID":"1","InquiryID":"1","TrasactionID":"123","OfferID":"123","Time":"2018-10-16T02:00:00Z"}
```
```
getTransactions
response
[{"FromType":"XX","FromID":"XX","ToType":"XX","ToID":"XX","Time":"XX","Number":"XX","ID":"XX"},{"FromType":"XX","FromID":"XX","ToType":"XX","ToID":"XX","Time":"XX","Number":"XX","ID":"XX"},...]
```
