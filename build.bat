@ECHO OFF
cd c:\gogit

:build
cls
ECHO Building.....
go build ./myHL/

IF NOT ERRORLEVEL 1 (
	go clean ./myHL/
	git commit -m "commit" -a
	git push
	c:\curl\curl.exe -X POST --insecure --header "Content-Type: application/json" --header "Accept: application/json" -d @deploy.json "https://13972c4ab5b1492bb508463f3cac0e60-vp0.us.blockchain.ibm.com:5002/chaincode" > serviceKey.txt
)


pause

goto :build
