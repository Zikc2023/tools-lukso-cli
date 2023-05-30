package errors

import "errors"

var (
	ErrNeedRoot             = errors.New("⚠️  You need root privileges to perform this action ")
	ErrFlagMissing          = errors.New("⚠️  Couldn't find given flag ")
	ErrMoreNetworksSelected = errors.New("⚠️  You can only specify 1 network ")
	ErrNotEnoughArguments   = errors.New("⚠️  Not enough arguments provided ")
	ErrProcessNotFound      = errors.New("⚠️  Process not found ")
	ErrFlagPathInvalid      = errors.New("⚠️  Invalid flag path ")
	ErrAlreadyRunning       = errors.New("⚠️  Process is already running ")
	ErrValidatorNotImported = errors.New("Validator has not been initialized - use 'lukso validator import' to initialize your validator ")
	ErrGenesisNotFound      = errors.New("❌  Genesis JSON not found")
)

const (
	NoSuchFlag              = "no such flag" // no emoji here - this error should match the CLI lib error - we don't throw it to user anyway
	FolderNotInitialized    = "⚠️  Folder not initialized - please make sure that you are working in an initialized directory. You can initialize the directory with the 'lukso init' command."
	SelectedClientsNotFound = "⚠️  No selected client found in LUKSO configuration file. Please make sure that you have installed your clients. You can use the install command to install clients."
	WrongPassword           = "wrong password for wallet"
)
