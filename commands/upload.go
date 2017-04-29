package commands

type UploadCommand struct {
	Vhd UploadVhdCommand `command:"vhd" alias:"v" description:"upload vhd file to the storage account container"`
}
