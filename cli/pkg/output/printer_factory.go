package output

import (
	"encoding/json"
	"fmt"
	"os"

	stackservice "github.com/chanzuckerberg/happy/cli/pkg/stack_mgr"
	"github.com/chanzuckerberg/happy/cli/pkg/util"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

type Printer interface {
	PrintStacks(stackInfos []stackservice.StackInfo) error
	Fatal(err error)
}

type TextPrinter struct{}
type JSONPrinter struct{}
type YAMLPrinter struct{}

func NewPrinter(outputFormat string) Printer {
	switch outputFormat {
	case "json":
		return &JSONPrinter{}
	case "yaml":
		return &YAMLPrinter{}
	default:
		return &TextPrinter{}
	}
}

type StackConsoleInfo struct {
	Name        string `header:"Name"`
	Owner       string `header:"Owner"`
	Tag         string `header:"Tags"`
	Status      string `header:"Status"`
	FrontendUrl string `header:"URLs"`
	LastUpdated string `header:"LastUpdated"`
}

func Stack2Console(stack stackservice.StackInfo) StackConsoleInfo {
	return StackConsoleInfo{
		Name:        stack.Name,
		Owner:       stack.Owner,
		Tag:         stack.Tag,
		Status:      stack.Status,
		FrontendUrl: stack.Outputs["frontend_url"],
		LastUpdated: stack.LastUpdated,
	}
}

func (p *TextPrinter) PrintStacks(stackInfos []stackservice.StackInfo) error {
	printer := util.NewTablePrinter()

	stacks := make([]StackConsoleInfo, 0)
	for _, stackInfo := range stackInfos {
		stacks = append(stacks, Stack2Console(stackInfo))
	}
	printer.Print(stacks)

	return nil
}

func (p *TextPrinter) Fatal(err error) {
	logrus.Fatal(err)
}

func (p *JSONPrinter) PrintStacks(stackInfos []stackservice.StackInfo) error {
	b, err := json.Marshal(stackInfos)
	if err != nil {
		return err
	}
	PrintOutput(string(b))
	return nil
}

func (p *JSONPrinter) Fatal(err error) {
	PrintError(err)
}

func (p *YAMLPrinter) PrintStacks(stackInfos []stackservice.StackInfo) error {
	b, err := yaml.Marshal(stackInfos)
	if err != nil {
		return err
	}
	PrintOutput(string(b))
	return nil
}

func (p *YAMLPrinter) Fatal(err error) {
	PrintError(err)
}

func PrintError(err error) {
	os.Stderr.WriteString(fmt.Sprintf("Error: %s\n", err.Error()))
}

func PrintOutput(output string) {
	os.Stdout.WriteString(fmt.Sprintf("%s\n", output))
}
