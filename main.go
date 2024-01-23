package main

import (
	"log"
	"os"
	"time"

	"github.com/rabilrbl/jiotv_go/v3/cmd"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:      "JioTV Go",
		Usage:     "Stream JioTV on any device",
		HelpName:  "jiotv_go",
		Version:  "v3.0.1",
		Copyright: "© JioTV Go by Mohammed Rabil (https://github.com/rabilrbl/jiotv_go)",
		Compiled:  time.Now(),
		Suggest:   true,
		Commands: []*cli.Command{
			{
				Name:        "serve",
				Aliases:     []string{"run", "start"},
				Usage:       "Start JioTV Go server",
				Description: "The serve command starts JioTV Go server, and listens on the host and port. The default host is localhost and port is 5001.",
				Action: func(c *cli.Context) error {
					host := c.String("host")
					// overwrite host if --public flag is passed
					if c.Bool("public") {
						log.Println("INFO: You are exposing your server to outside your local network (public)!")
						log.Println("INFO: Overwriting host to 0.0.0.0 for public access")
						host = "0.0.0.0"
					}
					port := c.String("port")
					prefork := c.Bool("prefork")
					configPath := c.String("config")
					return cmd.JioTVServer(host, port, configPath, prefork)
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "config",
						Aliases: []string{"c"},
						Value:   "",
						Usage:   "Path to config file",
					},
					&cli.StringFlag{
						Name:    "host",
						Aliases: []string{"H"},
						Value:   "localhost",
						Usage:   "Host to listen on",
					},
					&cli.StringFlag{
						Name:    "port",
						Aliases: []string{"p"},
						Value:   "5001",
						Usage:   "Port to listen on",
					},
					&cli.BoolFlag{
						Name:    "public",
						Aliases: []string{"P"},
						Usage:   "Open server to public. This will expose your server outside your local network. Equivalent to passing --host 0.0.0.0",
					},
					&cli.BoolFlag{
						Name:  "prefork",
						Usage: "Enable prefork. This will enable preforking the server to multiple processes. This is useful for production deployment.",
					},
				},
			},
			{
				Name:        "update",
				Aliases:     []string{"upgrade", "u"},
				Usage:       "Update JioTV Go to latest version",
				Description: "The update command updates JioTV Go by identifying the operating system and architecture, downloading the latest release from GitHub, and replacing the current binary with the latest one.",
				Action: func(c *cli.Context) error {
					return cmd.Update(c.App.Version, c.String("version"))
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "version",
						Aliases: []string{"v"},
						Value:   "",
						Usage:   "Update to a custom specific version that is not latest",
					},
				},
			},
			{
				Name:        "epg",
				Aliases:     []string{"e"},
				Usage:       "Manage EPG",
				Description: "The epg command manages EPG. It can be used to generate EPG, regenerate EPG, and delete EPG.",
				Subcommands: []*cli.Command{
					{
						Name:        "generate",
						Aliases:     []string{"gen", "g"},
						Usage:       "Generate EPG",
						Description: "The generate command generates EPG by downloading the latest EPG from JioTV, and saving it to epg.xml.gz. It will delete the existing EPG file if it exists. Once the EPG file is generated, it will automatically updated by the server. If you want to disable, do epg delete command.",
						Action: func(c *cli.Context) error {
							return cmd.GenEPG()
						},
					},
					{
						Name:        "Delete",
						Aliases:     []string{"del", "d"},
						Usage:       "Delete EPG",
						Description: "The delete command deletes the existing EPG file if it exists. This will disable EPG on the server.",
						Action: func(c *cli.Context) error {
							return cmd.DeleteEPG()
						},
					},
				},
			},
			{
				Name:        "login",
				Aliases:     []string{"l"},
				Usage:       "Manage login",
				Description: "The login command manages login. It can be used to login, logout.",
				Subcommands: []*cli.Command{
					{
						Name:        "otp",
						Aliases:     []string{"o"},
						Usage:       "Login using OTP",
						Description: "The otp command logs you in using OTP. It will send OTP to your mobile number, and you have to enter the OTP to login.",
						Action: func(c *cli.Context) error {
							return cmd.LoginOTP()
						},
					},
					{
						Name:        "password",
						Aliases:     []string{"p"},
						Usage:       "Login using password",
						Description: "The password command logs you in using password. It will ask for your username and password, and login using that.",
						Action: func(c *cli.Context) error {
							return cmd.LoginPassword()
						},
					},
					{
						Name:        "reset",
						Aliases:     []string{"lo", "logout"},
						Usage:       "Logout",
						Description: "The logout command logs you out. It will delete the login file.",
						Action: func(c *cli.Context) error {
							return cmd.Logout()
						},
					},
				},
			},
			{
				Name:        "autostart",
				Usage:       "Manage auto start for bash shell",
				Description: "The autostart command manages auto start for bash shell. It can be used to enable or disable auto start. We only support BASH Terminal and recommend on Android Termux.",
				Action: func(c *cli.Context) error {
					return cmd.AutoStart(c.String("args"))
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "args",
						Aliases: []string{"a"},
						Value:   "",
						Usage:   "String Value Arguments passed to serve/run/start command while auto starting",
					},
				},
			},
		},
		CommandNotFound: func(c *cli.Context, command string) {
			log.Printf("Command '%s' not found.\n", command)
			// Print help for invalid commands
			cli.ShowAppHelpAndExit(c, 3)
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}