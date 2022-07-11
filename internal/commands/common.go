// Package commands defines structures and variables used to
// parse command line input
package commands

// CommonOpts stores the options that are common to all image types
type CommonOpts struct {
	Debug            bool     `short:"d" long:"debug" description:"Enable debugging output"`
	Size             string   `short:"i" long:"image-size" description:"The suggested size of the generated disk image file. If this size is smaller than the minimum calculated size of the image a warning will be issued and --image-size will be ignored. The value is the size in bytes, with allowable suffixes \"M\" for MiB and \"G\" for GiB. Use an extended syntax to define the suggested size for the disk images generated by a multi-volume gadget.yaml spec" value-name:"SIZE"`
	ImageFileList    string   `long:"image-file-list" description:"Print to this file, a list of the file system paths to all the disk images created by the command, if any." value-name:"FILENAME"`
	CloudInit        string   `long:"cloud-init" description:"cloud-config data to be copied to the image" value-name:"USER-DATA-FILE"`
	HooksDirectories []string `long:"hooks-directory" description:"Path or comma-separated list of paths of directories in which scripts for build-time hooks will be located." value-name:"DIRECTORY"`
	DiskInfo         string   `long:"disk-info" description:"File to be used as .disk/info on the image's rootfs. This file can contain useful information about the target image, like image identification data, system name, build timestamp etc." value-name:"DISK-INFO-CONTENTS"`
	OutputDir        string   `short:"O" long:"output-dir" description:"The directory in which to put generated disk image files. The disk image files themselves will be named <volume>.img inside this directory, where <volume> is the volume name taken from the gadget.yaml file." value-name:"DIRECTORY"`
	Version          bool     `long:"version" description:"Print the version number of ubuntu-image and exit"`
	Channel          string   `short:"c" long:"channel" description:"The default snap channel to use" value-name:"CHANNEL"`
	SectorSize       string   `long:"sector-size" description:"Sector size to use when creating the disk image. Only 512 and 4k sector sizes are supported." choice:"512" choice:"4096" value-name:"SECTOR-SIZE" default:"512"`
}

// StateMachineOpts stores the options that are related to the state machine
type StateMachineOpts struct {
	WorkDir string `short:"w" long:"workdir" description:"The working directory in which to download and unpack all the source files for the image. This directory can exist or not, and it is not removed after this program exits. If not given, a temporary working directory is used instead, which *is* deleted after this program exits. Use -w if you want to be able to resume a partial state machine run." value-name:"DIRECTORY" group:"State Machine Options" default:""`
	Until   string `short:"u" long:"until" description:"Run the state machine until the given STEP, non-inclusively. STEP must be the name of the step." value-name:"STEP" default:""`
	Thru    string `short:"t" long:"thru" description:"Run the state machine through the given STEP, inclusively. STEP must be the name of the step." value-name:"STEP" default:""`
	Resume  bool   `short:"r" long:"resume" description:"Continue the state machine from the previously saved state. It is an error if there is no previous state."`
}

// UbuntuImageCommand is needed for the parser to store positional arguments and flags
type UbuntuImageCommand struct {
	Snap struct {
		SnapArgsPassed SnapArgs `positional-args:"true" required:"false"`
		SnapOptsPassed SnapOpts
	} `command:"snap"`
	Classic struct {
		ClassicArgsPassed ClassicArgs `positional-args:"true" required:"false"`
		ClassicOptsPassed ClassicOpts
	} `command:"classic"`
}

type commonOptions struct {
	CommonOptsPassed       CommonOpts
	StateMachineOptsPassed StateMachineOpts
}
