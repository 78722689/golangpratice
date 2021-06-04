package kafka

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type Arguments struct {
	cmd *flag.FlagSet

	Operation string
	Broker    string
	Topic     string
	Partition int
	Message   string
	Group     string
	Beginning bool
	Reblance  string
}

func (arg *Arguments) ParseKafkaArg() {
	help := false
	arg.cmd = flag.NewFlagSet("kafka", flag.ExitOnError)

	arg.cmd.StringVar(&arg.Operation, "operation", "consume", "Operate type. consumer, producer,create, list")
	arg.cmd.StringVar(&arg.Broker, "broker", "", "Kafka broker url")
	arg.cmd.StringVar(&arg.Topic, "topic", "", "Kafka topic name. Consumer group supports multiple topics split with comma.")
	arg.cmd.IntVar(&arg.Partition, "partition", -1, "Kafka partition number where to read/send")
	arg.cmd.StringVar(&arg.Message, "message", "", "Message to send to Kafka")
	arg.cmd.StringVar(&arg.Group, "cg", "", "Consumer group name")
	arg.cmd.BoolVar(&arg.Beginning, "begin", false, "Get all messages or the latest.")
	arg.cmd.StringVar(&arg.Reblance, "reblance", "roundrobin", "Consumer group reblance strategy.")

	arg.cmd.Parse(os.Args[2:])

	if (strings.Compare(arg.Operation, "produce") == 0 && arg.Message == "") ||
		(strings.Compare(arg.Operation, "consume") == 0 && arg.Topic == "") ||
		((strings.Compare(arg.Operation, "create") == 0 || strings.Compare(arg.Operation, "delete") == 0) && arg.Topic == "") ||
		help {
		arg.usage()
		os.Exit(1)
	}

}

func (arg *Arguments) usage() {
	fmt.Println("")
	fmt.Println("USAGE: sample.exe \n")
	fmt.Println("Kafka API remote call from client.\n")
	fmt.Println("ARGUMENTS:\n")

	arg.cmd.PrintDefaults()
}
