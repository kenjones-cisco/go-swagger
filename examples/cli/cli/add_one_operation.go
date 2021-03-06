// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/go-swagger/go-swagger/examples/cli/client/todos"
	"github.com/go-swagger/go-swagger/examples/cli/models"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationTodosAddOneCmd returns a cmd to handle operation addOne
func makeOperationTodosAddOneCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "addOne",
		Short: ``,
		RunE:  runOperationTodosAddOne,
	}

	if err := registerOperationTodosAddOneParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationTodosAddOne uses cmd flags to call endpoint api
func runOperationTodosAddOne(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := todos.NewAddOneParams()
	if cmd.Flags().Changed("body") {
		// Read body string from cmd and unmarshal
		bodyValueStr, err := cmd.Flags().GetString("body")
		if err != nil {
			return err
		}
		bodyValue := models.Item{}
		if err := bodyValue.UnmarshalBinary([]byte(bodyValueStr)); err != nil {
			return fmt.Errorf("cannot unmarshal body string in models.Item: %v", err)
		}
		params.Body = &bodyValue
	}
	bodyValueModel := params.Body
	if swag.IsZero(bodyValueModel) {
		bodyValueModel = &models.Item{}
	}
	err, added := retrieveItemFlags(bodyValueModel, "item", cmd)
	if err != nil {
		return err
	}
	if added {
		params.Body = bodyValueModel
	}
	bodyValueDebugBytes, err := json.Marshal(params.Body)
	if err != nil {
		return err
	}
	logDebugf("Body payload: %v", string(bodyValueDebugBytes))
	// make request and then print result
	resp, respErr := appCli.Todos.AddOne(params, nil)
	if err := printOperationTodosAddOneResult(resp, respErr); err != nil {
		return err
	}
	return nil
}

// printOperationTodosAddOneResult prints output to stdout
func printOperationTodosAddOneResult(resp *todos.AddOneCreated, respErr error) error {
	if respErr != nil {

		var iResp interface{} = respErr
		defaultResp, ok := iResp.(*todos.AddOneDefault)
		if !ok {
			return respErr
		}
		if defaultResp.Payload != nil {
			msgStr, err := json.Marshal(defaultResp.Payload)
			if err != nil {
				return err
			}
			fmt.Println(string(msgStr))
			return nil
		}

		return respErr
	}

	if resp.Payload != nil {
		msgStr, err := json.Marshal(resp.Payload)
		if err != nil {
			return err
		}
		fmt.Println(string(msgStr))
	}

	return nil
}

// registerOperationTodosAddOneParamFlags registers all flags needed to fill params
func registerOperationTodosAddOneParamFlags(cmd *cobra.Command) error {

	exampleBodyStr, err := json.Marshal(&models.Item{})
	if err != nil {
		return err
	}
	_ = cmd.PersistentFlags().String("body", "", fmt.Sprintf("Optional json string for body of form %v.", string(exampleBodyStr)))

	// add flags for body
	registerItemFlags("item", cmd)
	return nil
}
