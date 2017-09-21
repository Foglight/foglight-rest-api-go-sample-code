# Folders
- hostmonitor: A simple host resource monitor based on foglight rest api, support cpu/mem/disk/network utilization.
- rest: A simple go lib include some common code for data submission using foglight rest api
# Usage
- Get user token in Foglight
- Run hostmonitor on target host, support windows, linux, macos
# Build
- require go 1.8+ 
- cd hostmonitor
- go get -d
- go build
