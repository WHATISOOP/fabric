## couchdb增加视图查询接口
### peer channel create -o orderer.example.com:7050 -c demochannel -f ./demochannel.tx  
### peer channel join -b demochannel.block  
### peer chaincode install -n viewcc -v 1.0 -p github.com/hyperledger/fabric/examples/chaincode/go/chaincode_view
### peer chaincode instantiate -o orderer.example.com:7050 -C demochannel -n viewcc -v 1.0 -c '{"Args":["init"]}'  
### peer chaincode invoke -C demochannel -n viewcc -v 1.0 -c '{"Args":["add","{\"title\":\"Hello Blockchain\",\"body\":\"Welcome to blockchain world!\",\"date\":\"2018/11/10 13:12:20\"}"]}'  
### peer chaincode invoke -C demochannel -n viewcc -v 1.0 -c '{"Args":["add","{\"title\":\"Biking\",\"body\":\"My biggest hobby is mountainbiking. The other day...\",\"date\":\"2009/01/30 18:04:11\"}"]}'  
### peer chaincode invoke -C demochannel -n viewcc -v 1.0 -c '{"Args":["add","{\"title\":\"Bought a Cat\",\"body\":\"I went to the the pet store earlier and brought home a little kitty...\",\"date\":\"2009/01/30 18:04:11\"}"]}'  
### peer chaincode query -C demochannel -n viewcc -v 1.0 -c '{"Args":["get","6378806b97cc4d5753c0f8f85007f77850e160bdd92ecfbe07feb9cca29206b7"]}'  
## 创建view
### curl -H "Content-Type:application/json" -X PUT http://localhost:5984/demochannel/_design/demodesigndoc2 -d '{"views": {"demoview1": {"map":"function(doc){if(doc.data.date && doc.data.title) {emit(doc.data.title, doc); }}"}},"language": "javascript"}'         
## 查询view
### curl -H "Content-Type:application/json" localhost:5984/demochannel/_design/demodesigndoc2/_view/demoview1
### curl -H "Content-Type:application/json" localhost:5984/demochannel/_design/demodesigndoc2/_view/demoview1?key=%22Biking%22 
### curl -H "Content-Type:application/json" localhost:5984/demochannel/_design/demodesigndoc2/_view/demoview1?key=%22Hello%20World%22
## 通过chaincode查询view
### peer chaincode invoke -C demochannel -n viewcc -v 1.0 -c '{"Args":["searchByView","{\"designDocName\":\"demodesigndoc2\",\"viewName\":\"demoview1\",\"key\":\"Biking\"}"]}'
### peer chaincode invoke -C demochannel -n viewcc -v 1.0 -c '{"Args":["searchByView","{\"designDocName\":\"demodesigndoc2\",\"viewName\":\"demoview1\",\"key\":\"Hello Blockchain\"}"]}'