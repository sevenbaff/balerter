package mock

import (
	lua "github.com/yuin/gopher-lua"
	"strings"
)

func (m *ModuleMock) on(L *lua.LState) int {
	m.logger.Debug("[MOCK] on")

	if L.GetTop() == 0 {
		err := "mock.on should have first argument"
		m.logger.Error(err)
		m.errors = append(m.errors, err)
		return 0
	}

	methodNameL := L.Get(1)
	if methodNameL.Type() != lua.LTString {
		err := "mock.on first argument should be a string"
		m.logger.Error(err)
		m.errors = append(m.errors, err)
		return 0
	}

	methodName := strings.TrimSpace(methodNameL.String())
	if methodName == "" {
		err := "mock.on first argument should be not empty"
		m.logger.Error(err)
		m.errors = append(m.errors, err)
		return 0
	}

	var args []lua.LValue
	for i := 1; i < L.GetTop(); i++ {
		args = append(args, L.Get(i+1))
	}

	T := L.SetFuncs(L.NewTable(), map[string]lua.LGFunction{"response": m.saveResponse(methodName, args)})

	L.Push(T)

	return 1
}

func (m *ModuleMock) saveResponse(methodName string, callArgs []lua.LValue) lua.LGFunction {
	return func(L *lua.LState) int {
		retArgs := make([]lua.LValue, L.GetTop())
		for i := 0; i < L.GetTop(); i++ {
			retArgs[i] = L.Get(i + 1) // lua indexing starts with 1
		}

		err := m.registry.Register(AnyValue, methodName, callArgs, retArgs)
		if err != nil {
			m.errors = append(m.errors, "error register response: "+err.Error())
		}

		return 0
	}
}
