@ECHO OFF

:QUERY

for /f "delims=" %%x in (serviceKey.txt) do set Keytext=%%x

@echo. 
c:\curl\curl.exe -X POST --insecure --header "Content-Type: application/json" --header "Accept: application/json" -d "{\"jsonrpc\": \"2.0\",\"method\": \"query\",\"params\": {\"type\": 1,\"chaincodeID\": {\"name\": \"%Keytext:~52,128%\"},\"ctorMsg\": {\"function\": \"authenticate\",\"args\": [\"David\",\"Password1\"]},\"secureContext\": \"user_type1_2\"},\"id\": 1}" "https://7993495d15b440378dff6a89aaa0b616-vp0.us.blockchain.ibm.com:5002/chaincode"
@echo. 

pause

goto :QUERY