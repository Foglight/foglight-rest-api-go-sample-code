Folders
    hostmonitor: A simple host resource monitor based on foglight rest api, support cpu/mem percent, network transfer rate, disk usage.
    rest: A simple go lib include some common code for data submission using foglight rest api
Usage
    1. Import topology defined in file hostmonitor/topology-types.xml to FMS
    2. Get user token in Foglight
    3. Run hostmonitor on target host, support windows, linux, macos
Build
    cd hostmonitor
    go get
    go build