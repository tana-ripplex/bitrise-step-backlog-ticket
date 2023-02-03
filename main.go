package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {

	// [パラメーターを読み込む]
	// ※APIKey（認証情報）
	// ※プロジェクトID
	// 課題の件名
	// 課題の詳細
	// ※課題の種別ID
	// ※課題の優先度
	// お知らせするユーザーID
	// 課題のマイルストーン
	// backlogのID(xxxx.backlog.jp)

	// 課題の追加

	var endpoint = "https://rxdev.backlog.jp/find/RXGSP?allOver=false&assignerId=153170&offset=0&order=false&projectId=65406&simpleSearch=true&sort=UPDATED&statusId=1&statusId=2&statusId=3"

	// Uri
	var uri = "/api/v2/issues"

	fmt.Println("This is the value specified for the input 'example_step_input':", os.Getenv("example_step_input"))

	//
	// --- Step Outputs: Export Environment Variables for other Steps:
	// You can export Environment Variables for other Steps with
	//  envman, which is automatically installed by `bitrise setup`.
	// A very simple example:
	cmdLog, err := exec.Command("bitrise", "envman", "add", "--key", "EXAMPLE_STEP_OUTPUT", "--value", "the value you want to share").CombinedOutput()
	if err != nil {
		fmt.Printf("Failed to expose output with envman, error: %#v | output: %s", err, cmdLog)
		os.Exit(1)
	}
	// You can find more usage examples on envman's GitHub page
	//  at: https://github.com/bitrise-io/envman

	//
	// --- Exit codes:
	// The exit code of your Step is very important. If you return
	//  with a 0 exit code `bitrise` will register your Step as "successful".
	// Any non zero exit code will be registered as "failed" by `bitrise`.
	os.Exit(0)
}
