package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"syscall"

	"github.com/alexflint/go-arg"
	"github.com/brightpuddle/goaci"
	"github.com/brightpuddle/goaci/backup"
	"golang.org/x/crypto/ssh/terminal"
)

var version string

// Args are command line arguments.
type Args struct {
	Target   string   `arg:"positional,required" help:"Hostname or backup file" `
	Mode     string   `args:"-m" help:"Force mode [http|backup]"`
	Class    string   `arg:"-c" help:"Comma separated classnames to query"`
	Dn       string   `arg:"-d" help:"DN of the MO"`
	Filter   string   `arg:"-f" help:"Property filter to accept/reject MOs"`
	Options  []string `arg:"-x" help:"Extra query options"`
	User     string   `arg:"-u" help:"Username for APIC"`
	Password string   `arg:"-p" help:"Password for APIC"`
}

// Description described the app.
func (Args) Description() string {
	return "reQuery ACI offline and remote query tool."
}

func (Args) Version() string {
	return fmt.Sprintf("reQuery version %s", version)
}

// input get command line input from the user.
func input(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s ", prompt)
	input, _ := reader.ReadString('\n')
	return strings.Trim(input, "\r\n")
}

// printResults prints the result object.
func printResult(res goaci.Res, out io.Writer) {

	printObj := func(class, body goaci.Res) bool {
		fmt.Fprintln(out, "#", class.String()+".attributes")
		fmt.Fprintln(out, body.Get("attributes|@pretty"))
		return true
	}
	switch {
	case res.IsArray():
		fmt.Fprintf(out, "Total count: %d\n\n", len(res.Array()))
		for _, obj := range res.Array() {
			obj.ForEach(printObj)
		}
	case res.IsObject():
		fmt.Fprintf(out, "Total count: 1\n\n")
		res.ForEach(printObj)
	}
}

// httpQuery performs an HTTP query using goaci.Client.
func httpQuery(args Args) (res goaci.Res, err error) {
	if args.User == "" {
		args.User = input("Username:")
	}
	if args.Password == "" {
		fmt.Print("Password: ")
		pwd, _ := terminal.ReadPassword(int(syscall.Stdin))
		args.Password = string(pwd)
	}

	client, err := goaci.NewClient(args.Target, args.User, args.Password)
	if err != nil {
		return
	}

	if err = client.Login(); err != nil {
		return
	}

	// Process query parameters
	var query []func(*goaci.Req)
	if args.Filter != "" {
		query = append(query, goaci.Query("query-target-filter", args.Filter))
	}

	for _, option := range args.Options {
		parts := strings.Split(option, "=")
		if len(parts) != 2 {
			continue
		}
		query = append(query, goaci.Query(parts[0], parts[1]))
	}

	// Make request
	switch {
	case args.Dn != "":
		res, err = client.GetDn(args.Dn, query...)
	case args.Class != "":
		res, err = client.GetClass(args.Class, query...)
	default:
		err = errors.New("Class or DN is required.")
	}
	return
}

// backupQuery performs a backup file query using backup.Client.
func backupQuery(args Args) (res goaci.Res, err error) {
	client, err := backup.NewClient(args.Target)
	if err != nil {
		return
	}

	switch {
	case args.Dn != "":
		res, err = client.GetDn(args.Dn)
	case args.Class != "":
		res, err = client.GetClass(args.Class)
	default:
		err = errors.New("Class or DN is required.")
	}
	return
}

func main() {
	args := Args{}
	arg.MustParse(&args)

	// Determine mode
	var backupMode bool
	switch {
	case args.Mode == "http":
	case args.Mode == "backup" || strings.HasSuffix(args.Target, ".tar.gz"):
		backupMode = true
	}

	// Query
	var (
		res goaci.Res
		err error
	)
	if backupMode {
		res, err = backupQuery(args)
	} else {
		res, err = httpQuery(args)
	}
	if err != nil {
		log.Fatal(err)
	}
	printResult(res, os.Stdout)
}
