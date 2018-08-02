// Auto-generated by avdl-compiler v1.3.24 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/keybase1/prove_ui.avdl

package keybase1

import (
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
	context "golang.org/x/net/context"
)

type PromptOverwriteType int

const (
	PromptOverwriteType_SOCIAL PromptOverwriteType = 0
	PromptOverwriteType_SITE   PromptOverwriteType = 1
)

func (o PromptOverwriteType) DeepCopy() PromptOverwriteType { return o }

var PromptOverwriteTypeMap = map[string]PromptOverwriteType{
	"SOCIAL": 0,
	"SITE":   1,
}

var PromptOverwriteTypeRevMap = map[PromptOverwriteType]string{
	0: "SOCIAL",
	1: "SITE",
}

func (e PromptOverwriteType) String() string {
	if v, ok := PromptOverwriteTypeRevMap[e]; ok {
		return v
	}
	return ""
}

type PromptOverwriteArg struct {
	SessionID int                 `codec:"sessionID" json:"sessionID"`
	Account   string              `codec:"account" json:"account"`
	Typ       PromptOverwriteType `codec:"typ" json:"typ"`
}

type PromptUsernameArg struct {
	SessionID int     `codec:"sessionID" json:"sessionID"`
	Prompt    string  `codec:"prompt" json:"prompt"`
	PrevError *Status `codec:"prevError,omitempty" json:"prevError,omitempty"`
}

type OutputPrechecksArg struct {
	SessionID int  `codec:"sessionID" json:"sessionID"`
	Text      Text `codec:"text" json:"text"`
}

type PreProofWarningArg struct {
	SessionID int  `codec:"sessionID" json:"sessionID"`
	Text      Text `codec:"text" json:"text"`
}

type OutputInstructionsArg struct {
	SessionID    int    `codec:"sessionID" json:"sessionID"`
	Instructions Text   `codec:"instructions" json:"instructions"`
	Proof        string `codec:"proof" json:"proof"`
}

type OkToCheckArg struct {
	SessionID int    `codec:"sessionID" json:"sessionID"`
	Name      string `codec:"name" json:"name"`
	Attempt   int    `codec:"attempt" json:"attempt"`
}

type DisplayRecheckWarningArg struct {
	SessionID int  `codec:"sessionID" json:"sessionID"`
	Text      Text `codec:"text" json:"text"`
}

type ProveUiInterface interface {
	PromptOverwrite(context.Context, PromptOverwriteArg) (bool, error)
	PromptUsername(context.Context, PromptUsernameArg) (string, error)
	OutputPrechecks(context.Context, OutputPrechecksArg) error
	PreProofWarning(context.Context, PreProofWarningArg) (bool, error)
	OutputInstructions(context.Context, OutputInstructionsArg) error
	OkToCheck(context.Context, OkToCheckArg) (bool, error)
	DisplayRecheckWarning(context.Context, DisplayRecheckWarningArg) error
}

func ProveUiProtocol(i ProveUiInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "keybase.1.proveUi",
		Methods: map[string]rpc.ServeHandlerDescription{
			"promptOverwrite": {
				MakeArg: func() interface{} {
					ret := make([]PromptOverwriteArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]PromptOverwriteArg)
					if !ok {
						err = rpc.NewTypeError((*[]PromptOverwriteArg)(nil), args)
						return
					}
					ret, err = i.PromptOverwrite(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"promptUsername": {
				MakeArg: func() interface{} {
					ret := make([]PromptUsernameArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]PromptUsernameArg)
					if !ok {
						err = rpc.NewTypeError((*[]PromptUsernameArg)(nil), args)
						return
					}
					ret, err = i.PromptUsername(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"outputPrechecks": {
				MakeArg: func() interface{} {
					ret := make([]OutputPrechecksArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]OutputPrechecksArg)
					if !ok {
						err = rpc.NewTypeError((*[]OutputPrechecksArg)(nil), args)
						return
					}
					err = i.OutputPrechecks(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"preProofWarning": {
				MakeArg: func() interface{} {
					ret := make([]PreProofWarningArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]PreProofWarningArg)
					if !ok {
						err = rpc.NewTypeError((*[]PreProofWarningArg)(nil), args)
						return
					}
					ret, err = i.PreProofWarning(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"outputInstructions": {
				MakeArg: func() interface{} {
					ret := make([]OutputInstructionsArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]OutputInstructionsArg)
					if !ok {
						err = rpc.NewTypeError((*[]OutputInstructionsArg)(nil), args)
						return
					}
					err = i.OutputInstructions(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"okToCheck": {
				MakeArg: func() interface{} {
					ret := make([]OkToCheckArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]OkToCheckArg)
					if !ok {
						err = rpc.NewTypeError((*[]OkToCheckArg)(nil), args)
						return
					}
					ret, err = i.OkToCheck(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"displayRecheckWarning": {
				MakeArg: func() interface{} {
					ret := make([]DisplayRecheckWarningArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]DisplayRecheckWarningArg)
					if !ok {
						err = rpc.NewTypeError((*[]DisplayRecheckWarningArg)(nil), args)
						return
					}
					err = i.DisplayRecheckWarning(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
		},
	}
}

type ProveUiClient struct {
	Cli rpc.GenericClient
}

func (c ProveUiClient) PromptOverwrite(ctx context.Context, __arg PromptOverwriteArg) (res bool, err error) {
	err = c.Cli.Call(ctx, "keybase.1.proveUi.promptOverwrite", []interface{}{__arg}, &res)
	return
}

func (c ProveUiClient) PromptUsername(ctx context.Context, __arg PromptUsernameArg) (res string, err error) {
	err = c.Cli.Call(ctx, "keybase.1.proveUi.promptUsername", []interface{}{__arg}, &res)
	return
}

func (c ProveUiClient) OutputPrechecks(ctx context.Context, __arg OutputPrechecksArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.proveUi.outputPrechecks", []interface{}{__arg}, nil)
	return
}

func (c ProveUiClient) PreProofWarning(ctx context.Context, __arg PreProofWarningArg) (res bool, err error) {
	err = c.Cli.Call(ctx, "keybase.1.proveUi.preProofWarning", []interface{}{__arg}, &res)
	return
}

func (c ProveUiClient) OutputInstructions(ctx context.Context, __arg OutputInstructionsArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.proveUi.outputInstructions", []interface{}{__arg}, nil)
	return
}

func (c ProveUiClient) OkToCheck(ctx context.Context, __arg OkToCheckArg) (res bool, err error) {
	err = c.Cli.Call(ctx, "keybase.1.proveUi.okToCheck", []interface{}{__arg}, &res)
	return
}

func (c ProveUiClient) DisplayRecheckWarning(ctx context.Context, __arg DisplayRecheckWarningArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.proveUi.displayRecheckWarning", []interface{}{__arg}, nil)
	return
}
