package evaluator

import (
	"abs/object"
	"abs/util"
	"fmt"
	"os"
	"strconv"
)

func getFns() map[string]*object.Builtin {
	return map[string]*object.Builtin{
		"len": &object.Builtin{Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return util.NewError("wrong number of arguments. got=%d, want=1",
					len(args))
			}

			switch arg := args[0].(type) {
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			default:
				return util.NewError("argument to `len` not supported, got %s",
					args[0].Type())
			}
		},
		},
		"echo": &object.Builtin{
			Fn: func(args ...object.Object) object.Object {
				var arguments []interface{} = make([]interface{}, len(args)-1)
				for i, d := range args {
					if i > 0 {
						arguments[i-1] = d.Inspect()
					}
				}

				fmt.Printf(args[0].Inspect(), arguments...)
				fmt.Println("")

				return NULL
			},
		},
		"int": &object.Builtin{
			Fn: func(args ...object.Object) object.Object {
				if len(args) != 1 {
					return util.NewError("wrong number of arguments. got=%d, want=1", len(args))
				}

				switch arg := args[0].(type) {
				case *object.Integer:
					return &object.Integer{Value: int64(arg.Value)}
				case *object.String:
					i, err := strconv.Atoi(arg.Value)

					if err != nil {
						return util.NewError("int(...) can only be called on strings which represent integers, '%s' given", arg.Value)
					}

					return &object.Integer{Value: int64(i)}
				default:
					return util.NewError("argument to `int` not supported, got %s", args[0].Type())
				}
			},
		},
		"env": &object.Builtin{
			Fn: func(args ...object.Object) object.Object {
				if len(args) != 1 {
					return util.NewError("wrong number of arguments. got=%d, want=1",
						len(args))
				}

				switch arg := args[0].(type) {
				case *object.String:
					return &object.String{Value: os.Getenv(arg.Value)}
				default:
					return util.NewError("argument to `env` not supported, got %s",
						args[0].Type())
				}
			},
		},
	}

}