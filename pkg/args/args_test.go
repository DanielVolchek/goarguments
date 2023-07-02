package args

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"testing"
)

// func FormatError(key string, expected string, result string) string {
// 	return fmt.Sprintf("Expected %s %s, received %s %s", key, expected, key, result)
// }

func CompareExpectedAndResult(expectedResult *[]CmdArg, actualResult *[]CmdArg) error {

	expectedString := fmt.Sprintf("%#v", *expectedResult)
	resultString := fmt.Sprintf("%#v", *actualResult)

	if expectedString != resultString {
		return errors.New(fmt.Sprintf("expected to receive: \n%s\nreceived: %s", expectedString, resultString))
	}

	return nil
}

type ArgTest struct {
	osArgs   []string
	expected []CmdArg
}

func RunTestOnArgsHelper(test ArgTest) error {
	os.Args = test.osArgs

	result, err := LoadCmdArgs()

	if err != nil {
		return errors.Join(errors.New("expected err to be nil, instead received %s"), err)
	}

	err = CompareExpectedAndResult(&test.expected, &result)

	if err != nil {
		return err
	}

	return nil
}

func TestLoadCmdArgs(t *testing.T) {
	// test: check if the function parses normal commands without modifiers correctly

	t.Run("check_parse_no_modifiers 1", func(t *testing.T) {
		currentTest := ArgTest{
			osArgs: []string{"file", "--seed"},
			expected: []CmdArg{
				{arg: "--seed"}},
		}
		err := RunTestOnArgsHelper(currentTest)
		if err != nil {
			t.Errorf(err.Error())
		}
	})

	t.Run("check_parse_no_modifiers 2", func(t *testing.T) {
		currentTest := ArgTest{
			osArgs: []string{"file", "--seed", "--hello"},
			expected: []CmdArg{
				{arg: "--seed"}, {arg: "--hello"},
			}}
		err := RunTestOnArgsHelper(currentTest)
		if err != nil {
			t.Errorf(err.Error())
		}
	})

	t.Run("check_parse_no_modifiers 3", func(t *testing.T) {
		currentTest := ArgTest{
			osArgs: []string{"file", "--seed", "--hello", "--world"},
			expected: []CmdArg{
				{arg: "--seed"}, {arg: "--hello"}, {arg: "--world"},
			}}
		err := RunTestOnArgsHelper(currentTest)
		if err != nil {
			t.Errorf(err.Error())
		}
	})

	// test: check if the function parses normal commands with modifiers correctly

	t.Run("check_parse_modifiers 1", func(t *testing.T) {
		currentTest := ArgTest{
			osArgs: []string{"file", "--hello", "world"},
			expected: []CmdArg{
				{arg: "--hello", value: "world"},
			}}
		err := RunTestOnArgsHelper(currentTest)
		if err != nil {
			t.Errorf(err.Error())
		}
	})

	// test: check if the function parses normal commands with a mixture of modifiers and no modifiers correctly
	t.Run("check_parse_modifiers 2", func(t *testing.T) {
		currentTest := ArgTest{
			osArgs: []string{"file", "--hello", "world", "--goodbye", "--universe"},
			expected: []CmdArg{
				{arg: "--hello", value: "world"}, {arg: "--goodbye"}, {arg: "--universe"},
			}}
		err := RunTestOnArgsHelper(currentTest)
		if err != nil {
			t.Errorf(err.Error())
		}
	})

	// test: check if the function parses incorrectly formatted commands correctly (returning an error)

	t.Run("check formatted incorrectly", func(t *testing.T) {
		currentTest := ArgTest{
			osArgs:   []string{"file", "hello", "world", "--goodbye", "universe"},
			expected: nil,
		}
		err := RunTestOnArgsHelper(currentTest)
		if err == nil {
			t.Errorf(err.Error())
		}
	})

	// test: check if the function parses no args correctly
	t.Run("check formatted correctly empty", func(t *testing.T) {
		currentTest := ArgTest{
			osArgs:   []string{""},
			expected: nil,
		}
		err := RunTestOnArgsHelper(currentTest)
		if err == nil {
			t.Errorf("Expected to get an error but did not")
		}
	})
}

func TestEnvArgs(t *testing.T) {

	// tests to see if it gets the right env vars
	t.Run("test_env", func(t *testing.T) {
		t.Setenv("hello", "world")
		t.Setenv("goodbye", "universe")

		args := LoadEnvArgs("hello", "goodbye")
		expected := []string{"world", "universe"}
		if !reflect.DeepEqual(args, expected) {
			t.Errorf("Expected args to equal %#v, but was instead %#v", expected, args)
		}
	})

}
