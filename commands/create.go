package commands

type CreateCommand struct {
	Container CreateContainerCommand `command:"container" alias:"c" description:"Create the target container"`
}
