package evaluator

import (
	"fmt"
	"interpreter/object"
)

var builtins = map[string]*object.Builtin{
	"len": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments got %d, but wanted %d", len(args), 1)
			}

			switch arg := args[0].(type) {

			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}

			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}

			default:
				return newError("argument to `len` not supported, got %s", args[0].Type())
			}
		},
	},
	"first": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments got %d, but wanted %d", len(args), 1)
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument input is not an array object, got %s", args[0].Type())
			}
			arr := args[0].(*object.Array)
			if len(arr.Elements) > 0 {
				return arr.Elements[0]
			}

			return NULL
		},
	},
	"last": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments got %d, but wanted %d", len(args), 1)
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument input is not an array object, got %s", args[0].Type())
			}
			arr := args[0].(*object.Array)
			length := len(arr.Elements) - 1
			if len(arr.Elements) > 0 {
				return arr.Elements[length]
			}

			return NULL
		},
	},
	"rest": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments got %d, but wanted %d", len(args), 1)
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument input is not an array object, got %s", args[0].Type())
			}
			arr := args[0].(*object.Array)
			length := len(arr.Elements) - 1
			if len(arr.Elements) > 0 {
				newElements := make([]object.Object, length-1, length-1)
				copy(newElements, arr.Elements[1:length])
				return &object.Array{Elements: newElements}
			}

			return NULL
		},
	},
	"put": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			for _, arg := range args {
				fmt.Println(arg.Inspect())
			}

			return NULL
		},
	},
}
