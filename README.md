# Go API client for opengraph

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v0.3.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import sw "./opengraph"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), sw.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), sw.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://ocis.ocis-traefik.latest.owncloud.works*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DrivesApi* | [**CreateDrive**](docs/DrivesApi.md#createdrive) | **Post** /drives | Create a new space of a specific type
*DrivesApi* | [**DeleteDrive**](docs/DrivesApi.md#deletedrive) | **Delete** /drives/{drive-id} | Delete a specific space
*DrivesApi* | [**GetDrive**](docs/DrivesApi.md#getdrive) | **Get** /drives/{drive-id} | Get drive by id
*DrivesApi* | [**UpdateDrive**](docs/DrivesApi.md#updatedrive) | **Patch** /drives/{drive-id} | Update the space
*DrivesRootApi* | [**GetRoot**](docs/DrivesRootApi.md#getroot) | **Get** /drives/{drive-id}/root | Get root from arbitrary space
*MeDriveApi* | [**GetHome**](docs/MeDriveApi.md#gethome) | **Get** /me/drive | Get personal space for user
*MeDriveRootApi* | [**HomeGetRoot**](docs/MeDriveRootApi.md#homegetroot) | **Get** /me/drive/root | Get root from personal space
*MeDriveRootChildrenApi* | [**HomeGetChildren**](docs/MeDriveRootChildrenApi.md#homegetchildren) | **Get** /me/drive/root/children | Get children from drive
*MeDrivesApi* | [**ListMyDrives**](docs/MeDrivesApi.md#listmydrives) | **Get** /me/drives | Get drives from me


## Documentation For Models

 - [CollectionOfDriveItems](docs/CollectionOfDriveItems.md)
 - [CollectionOfDrives](docs/CollectionOfDrives.md)
 - [Deleted](docs/Deleted.md)
 - [Drive](docs/Drive.md)
 - [DriveItem](docs/DriveItem.md)
 - [FileSystemInfo](docs/FileSystemInfo.md)
 - [Folder](docs/Folder.md)
 - [FolderView](docs/FolderView.md)
 - [Hashes](docs/Hashes.md)
 - [Identity](docs/Identity.md)
 - [IdentitySet](docs/IdentitySet.md)
 - [Image](docs/Image.md)
 - [ItemReference](docs/ItemReference.md)
 - [OdataError](docs/OdataError.md)
 - [OdataErrorDetail](docs/OdataErrorDetail.md)
 - [OdataErrorMain](docs/OdataErrorMain.md)
 - [OpenGraphFile](docs/OpenGraphFile.md)
 - [Quota](docs/Quota.md)
 - [User](docs/User.md)


## Documentation For Authorization

 Endpoints do not require authorization.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



