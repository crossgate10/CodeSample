@echo off
chcp 950

REM echo %~DP0

mklink "run-gRPCServer.exe" ".\gRPCServer\bin\Debug\gRPCServer.exe"
mklink "run-gRPCClient.exe" ".\gRPCClient\bin\Debug\gRPCClient.exe"