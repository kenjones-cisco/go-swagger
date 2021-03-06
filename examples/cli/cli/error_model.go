// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/go-swagger/examples/cli/models"
	"github.com/spf13/cobra"
)

// register flags to command
func registerErrorFlags(cmdPrefix string, cmd *cobra.Command) error {

	codeCmdStr := fmt.Sprintf("%v.code", cmdPrefix)

	var codeCmdStrDefault int64
	_ = cmd.PersistentFlags().Int64(codeCmdStr, codeCmdStrDefault, "")

	messageCmdStr := fmt.Sprintf("%v.message", cmdPrefix)

	var messageCmdStrDefault string
	_ = cmd.PersistentFlags().String(messageCmdStr, messageCmdStrDefault, "Required. ")

	if err := cmd.MarkPersistentFlagRequired(messageCmdStr); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveErrorFlags(m *models.Error, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	codeCmdStr := fmt.Sprintf("%v.code", cmdPrefix)
	if cmd.Flags().Changed(codeCmdStr) {
		codeValue, err := cmd.Flags().GetInt64(codeCmdStr)
		if err != nil {
			return err, false
		}
		m.Code = codeValue
		retAdded = true
	}
	messageCmdStr := fmt.Sprintf("%v.message", cmdPrefix)
	if cmd.Flags().Changed(messageCmdStr) {
		messageValue, err := cmd.Flags().GetString(messageCmdStr)
		if err != nil {
			return err, false
		}
		m.Message = &messageValue
		retAdded = true
	}

	return nil, retAdded
}
