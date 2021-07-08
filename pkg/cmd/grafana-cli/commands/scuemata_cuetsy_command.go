package commands

import (
	"fmt"
	"os"
	"path/filepath"

	"cuelang.org/go/cue"
	"github.com/grafana/grafana/pkg/cmd/grafana-cli/utils"
	"github.com/grafana/grafana/pkg/schema/load"
	"github.com/sdboyer/cuetsy/encoder"
)

type attrTSTarget string

const (
	tgtType      attrTSTarget = "type"
	tgtInterface attrTSTarget = "interface"
	tgtEnum      attrTSTarget = "enum"
)

func (cmd Command) generateDashboardTypeScripts(c utils.CommandLine) error {
	dest := c.String("dest")

	if err := generateTypeScriptFromCUE(dest, paths); err != nil {
		return err
	}

	return nil
}

func generateTypeScriptFromCUE(dest string, p load.BaseLoadPaths) error {
	panelSchemaMap, err := load.ReadPanelModels(p)
	if err != nil {
		return err
	}

	for panelName, panelSchema := range panelSchemaMap {
		// got := fmt.Sprintf("%+v", panelSchema.CUE())
		// fmt.Println("<<<<<<<<<<", got)
		obj, err := prepare(panelSchema.CUE())

		b, err := encoder.Generate(panelSchema.CUE(), encoder.Config{})
		if err != nil {
			return err
		}
		writeTypeScriptFiles(filepath.Join(dest, panelName+".ts"), string(b))
	}
	return nil
}

func prepare(iCUE cue.Value) (cue.Value, error) {
	iter, err := iCUE.Fields(cue.Definitions(true))
	if err != nil {
		return iCUE, err
	}
	for iter.Next() {
		/*
			It is very difficult to add new Attribute to Value here,
			so we probably want just to use the underline function of cuetsy
		*/
		a := iter.Value().Attribute("cuetsy")
		tt, found, err := a.Lookup(0, "targetType")
		targetType := attrTSTarget(tt)

		if a.Err() != nil || !found || err != nil {
			targetType = tgtInterface
		}
		switch targetType {

		}
		if err != nil {
			return iCUE, err
		}
	}
	return iCUE, err
}

func writeTypeScriptFiles(dest string, content string) error {
	fd, err := os.Create(dest)
	if err != nil {
		return err
	}
	fmt.Fprint(fd, content)
	return nil
}
