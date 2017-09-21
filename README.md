# Folders
    - hostmonitor: A simple host resource monitor based on foglight rest api, support cpu/mem percent, network transfer rate, disk usage.
    - rest: A simple go lib include some common code for data submission using foglight rest api
# Usage
    - Import topology defined in file hostmonitor/topology-types.xml to FMS
    - Get user token in Foglight
    - Run hostmonitor on target host, support windows, linux, macos
# Build
    - cd hostmonitor
    - go get
    - go build
