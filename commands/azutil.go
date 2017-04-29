package commands

type AzutilCommand struct {
	Create CreateCommand `command:"create" alias:"c" description:"Create the target resource"`
	Upload UploadCommand `command:"upload" alias:"u" description:"Upload the target resource"`
}

var Azutil AzutilCommand
