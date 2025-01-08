package tests

import (
	file "automatic-doodle/tests_parts/file"
	global "automatic-doodle/tests_parts/global"
	test "automatic-doodle/tests_parts/global"
	job "automatic-doodle/tests_parts/job"
	userTest "automatic-doodle/tests_parts/user"
	"testing"
)

func TestRegister(t *testing.T) {
	userTest.SendUserRegistrationRequest()
	// Will only run true once
}
func TestBrowseJob(t *testing.T) {
	response, _ := userTest.LoginAndGetTokens(test.UrlLogin, userTest.LoginReq)

	println(response.Token)

	job.BrowseJob(response.Token, global.UrlJobCreate, job.JobQuery{Identifier: "Full", PageNumber: 0})

}

func TestSendJobAdv(t *testing.T) {
	response, _ := userTest.LoginAndGetTokens(test.UrlLogin, userTest.LoginReq)

	resp, err := job.SendJobAdvRequest(response.Token)

	if err != nil {
		panic(err)
	}

	println("%s", resp.Description)

	fileResp, err := file.CreateUploadUrl("mertens.jpg", "POST_FILE", response.Token)

	if err != nil {
		panic(err)
	}

	file.UploadToLocalStackS3("localhost:4566", "POST_FILE", "mertens.jpg", fileResp.Key)
}
