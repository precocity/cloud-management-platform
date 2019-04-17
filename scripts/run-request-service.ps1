
$script_dir = Get-Location
$u_root_dir = $script_dir.ToString() + "\.."
$root_dir = Resolve-Path $u_root_dir

$bin_dir = Join-Path -Path $root_dir -ChildPath "\bin"
$build_dir = Join-Path -Path $root_dir -ChildPath "\build"

if(Test-Path $bin_dir) {
    Remove-Item -Path $bin_dir -Recurse
}

New-Item -Path $bin_dir -ItemType "directory"

Push-Location $bin_dir

$env:CGO_ENABLED = 0 
$env:GOOS = "linux"

go build -a -installsuffix cgo $root_dir/cmd/resource-request-service/main.go

Rename-Item -Path $bin_dir\main -NewName request-service

docker-compose -f $build_dir/docker-compose.yml rm -vsf
docker-compose -f $build_dir/docker-compose.yml down -v --remove-ophans
docker-compose -f $build_dir/docker-compose.yml build request.service requests.db

#starting up dockerized development database
docker-compose -f $build_dir/docker-compose.yml up -d requests.db

#starting up go application
docker-compose -f $build_dir/docker-compose.yml up -d request.service

#show logs
docker-compose -f $build_dir/docker-compose.yml logs -f request.service requests.db

Pop-Location