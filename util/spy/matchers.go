package spy

import (
	"fmt"

	"github.com/onsi/gomega/types"
)

//
// Called matcher
//

func Called() types.GomegaMatcher {
	return &called{}
}

type called struct{}

func (cw *called) Match(actual any) (bool, error) {
	spy, ok := actual.(Spy)
	if !ok {
		return false, fmt.Errorf("BeCalled can only match Spy objects")
	}

	if !spy.Called() {
		return false, nil
	}
	return true, nil
}

func (cw *called) FailureMessage(actual any) string {
	return fmt.Sprintf("Expected call")
}

func (cw *called) NegatedFailureMessage(actual any) string {
	return fmt.Sprintf("Expected no calls")
}

//
// Called-With-Args matcher
//

func CalledWith(expectedArgs ...any) types.GomegaMatcher {
	return &calledWith{expectedArgs}
}

type calledWith struct {
	expected []any
}

func (cw *calledWith) Match(actual any) (bool, error) {
	spy, ok := actual.(Spy)
	if !ok {
		return false, fmt.Errorf("CalledWith can only match Spy objects")
	}

	if !spy.Called() {
		return false, nil
	}

	if !spy.CalledWith(cw.expected...) {
		return false, nil
	}

	return true, nil
}

func (cw *calledWith) FailureMessage(actual any) string {
	return fmt.Sprintf("Expected call with args:\n\t%#v", cw.expected)
}

func (cw *calledWith) NegatedFailureMessage(actual any) string {
	return fmt.Sprintf("Expected no calls with args:\n\t%#v", cw.expected)
}
