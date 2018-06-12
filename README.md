## 前置条件
    启动文件位于/opt/goworkspace/src/github.com/whatisoop/fabric/examples/dev/resources目录下
    docker-compose up
    peer channel create -o orderer.example.com:7050 -c demochannel -f ./demochannel.tx  
    peer channel join -b demochannel.block 

## couchdb增加视图查询接口
### 测试数据
    peer chaincode install -n viewcc -v 1.0 -p github.com/hyperledger/fabric/examples/chaincode/go/chaincode_view
    peer chaincode instantiate -o orderer.example.com:7050 -C demochannel -n viewcc -v 1.0 -c '{"Args":["init"]}'  
    peer chaincode invoke -C demochannel -n viewcc -v 1.0 -c '{"Args":["add","{\"title\":\"Hello Blockchain\",\"body\":\"Welcome to blockchain world!\",        \"date\":\"2018/11/10 13:12:20\"}"]}'  
    peer chaincode invoke -C demochannel -n viewcc -v 1.0 -c '{"Args":["add","{\"title\":\"Biking\",\"body\":\"My biggest hobby is mountainbiking. The other day...\",\"date\":\"2009/01/30 18:04:11\"}"]}'  
    peer chaincode invoke -C demochannel -n viewcc -v 1.0 -c '{"Args":["add","{\"title\":\"Bought a Cat\",\"body\":\"I went to the the pet store earlier and brought home a little kitty...\",\"date\":\"2009/01/30 18:04:11\"}"]}'  
    peer chaincode query -C demochannel -n viewcc -v 1.0 -c '{"Args":["get","6378806b97cc4d5753c0f8f85007f77850e160bdd92ecfbe07feb9cca29206b7"]}'  
### 创建view
    curl -H "Content-Type:application/json" -X PUT http://localhost:5984/demochannel/_design/demodesigndoc2 -d '{"views": {"demoview1": {"map":"function(doc){if(doc.data.date && doc.data.title) {emit(doc.data.title, doc); }}"}},"language": "javascript"}'         
### 查询view
    curl -H "Content-Type:application/json" localhost:5984/demochannel/_design/demodesigndoc2/_view/demoview1
    curl -H "Content-Type:application/json" localhost:5984/demochannel/_design/demodesigndoc2/_view/demoview1?key=%22Biking%22 
    curl -H "Content-Type:application/json" localhost:5984/demochannel/_design/demodesigndoc2/_view/demoview1?key=%22Hello%20World%22
### 通过chaincode查询view
    peer chaincode invoke -C demochannel -n viewcc -v 1.0 -c '{"Args":["searchByView","{\"designDocName\":\"demodesigndoc2\",\"viewName\":\"demoview1\",\"key\":\"Biking\"}"]}'
    peer chaincode invoke -C demochannel -n viewcc -v 1.0 -c '{"Args":["searchByView","{\"designDocName\":\"demodesigndoc2\",\"viewName\":\"demoview1\",\"key\":\"Hello Blockchain\"}"]}'

## 增加chaincode_token
### 功能列表
参考 https://cloud.tencent.com/developer/article/1066053 
1. token系统  
    1）发行token。用户可以发行多个token，属性有：token的名称、token的代号、token的总量、token的发行者、token的状态（锁仓、正常）  
      token的代号不能重复。token的发行者必须存在。
    2）转账。
      冻结的账户不能转账。锁仓的token不能转账。余额不足不能转账。不存在的token不能转账。  
    3）余额查询。可以指定查询指定账户的某个token余额，可以查询指定账户的所有token的余额。    
    4）token锁仓/解锁
    5）token列表。展示1)中发行的token信息。
2. 账户系统  
    1）账户属性包括账户名称、账户状态（冻结、正常）、账户余额  
    2）创建账户  
       账户名称不可重复。账户状态默认是正常。  
    3）更新账户状态（冻结/解冻）  
    4）查询账户：根据账户名称返回1）中创建的账户信息。  
### 测试数据
    peer chaincode install -n token -v 1.0 -p github.com/hyperledger/fabric/examples/chaincode/go/chaincode_token
    peer chaincode instantiate -C demochannel -n token -v 1.0 -c '{"Args":[""]}'
    peer chaincode invoke -C demochannel -n token -v 1.0 -c '{"function":"createAccount","Args":["coinbase"]}' 
    peer chaincode invoke -C demochannel -n token -v 1.0 -c '{"function":"createAccount","Args":["whatisoop"]}' 
    peer chaincode invoke -C demochannel -n token -v 1.0 -c '{"function":"showAccount","Args":["coinbase"]}'
    peer chaincode invoke -C demochannel -n token -v 1.0 -c '{"function":"balanceAll","Args":["coinbase"]}'
    peer chaincode invoke -C demochannel -n token -v 1.0 -c '{"function":"balance","Args":["coinbase","NKC"]}'
    peer chaincode invoke -C demochannel -n token -v 1.0 -c '{"function":"frozenAccount","Args":["coinbase","true"]}'
    peer chaincode invoke -C demochannel -n token -v 1.0 -c '{"function":"createToken","Args":["Apple Token","apple","1000000","coinbase"]}'
    peer chaincode invoke -C demochannel -n token -v 1.0 -c '{"function":"createToken","Args":["Orange Token","orange","5000000","coinbase"]}'
    peer chaincode invoke -C demochannel -n token -v 1.0 -c '{"function":"setLock","Args":["apple","true"]}'
    peer chaincode invoke -C demochannel -n token -v 1.0 -c '{"function":"setLock","Args":["apple","false"]}'
    peer chaincode invoke -C demochannel -n token -v 1.0 -c '{"function":"transferToken","Args":["coinbase","whatisoop","orange","300"]}'		

