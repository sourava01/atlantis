// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/runatlantis/atlantis/server/events/vcs (interfaces: Client)

package mocks

import (
	pegomock "github.com/petergtz/pegomock/v4"
	models "github.com/runatlantis/atlantis/server/events/models"
	"reflect"
	"time"
)

type MockClient struct {
	fail func(message string, callerSkip ...int)
}

func NewMockClient(options ...pegomock.Option) *MockClient {
	mock := &MockClient{}
	for _, option := range options {
		option.Apply(mock)
	}
	return mock
}

func (mock *MockClient) SetFailHandler(fh pegomock.FailHandler) { mock.fail = fh }
func (mock *MockClient) FailHandler() pegomock.FailHandler      { return mock.fail }

func (mock *MockClient) CreateComment(repo models.Repo, pullNum int, comment string, command string) error {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClient().")
	}
	params := []pegomock.Param{repo, pullNum, comment, command}
	result := pegomock.GetGenericMockFrom(mock).Invoke("CreateComment", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockClient) DiscardReviews(repo models.Repo, pull models.PullRequest) error {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClient().")
	}
	params := []pegomock.Param{repo, pull}
	result := pegomock.GetGenericMockFrom(mock).Invoke("DiscardReviews", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockClient) GetCloneURL(VCSHostType models.VCSHostType, repo string) (string, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClient().")
	}
	params := []pegomock.Param{VCSHostType, repo}
	result := pegomock.GetGenericMockFrom(mock).Invoke("GetCloneURL", params, []reflect.Type{reflect.TypeOf((*string)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 string
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(string)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockClient) GetFileContent(pull models.PullRequest, fileName string) (bool, []byte, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClient().")
	}
	params := []pegomock.Param{pull, fileName}
	result := pegomock.GetGenericMockFrom(mock).Invoke("GetFileContent", params, []reflect.Type{reflect.TypeOf((*bool)(nil)).Elem(), reflect.TypeOf((*[]byte)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 bool
	var ret1 []byte
	var ret2 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(bool)
		}
		if result[1] != nil {
			ret1 = result[1].([]byte)
		}
		if result[2] != nil {
			ret2 = result[2].(error)
		}
	}
	return ret0, ret1, ret2
}

func (mock *MockClient) GetModifiedFiles(repo models.Repo, pull models.PullRequest) ([]string, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClient().")
	}
	params := []pegomock.Param{repo, pull}
	result := pegomock.GetGenericMockFrom(mock).Invoke("GetModifiedFiles", params, []reflect.Type{reflect.TypeOf((*[]string)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 []string
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].([]string)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockClient) GetPullLabels(repo models.Repo, pull models.PullRequest) ([]string, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClient().")
	}
	params := []pegomock.Param{repo, pull}
	result := pegomock.GetGenericMockFrom(mock).Invoke("GetPullLabels", params, []reflect.Type{reflect.TypeOf((*[]string)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 []string
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].([]string)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockClient) GetTeamNamesForUser(repo models.Repo, user models.User) ([]string, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClient().")
	}
	params := []pegomock.Param{repo, user}
	result := pegomock.GetGenericMockFrom(mock).Invoke("GetTeamNamesForUser", params, []reflect.Type{reflect.TypeOf((*[]string)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 []string
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].([]string)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockClient) HidePrevCommandComments(repo models.Repo, pullNum int, command string, dir string) error {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClient().")
	}
	params := []pegomock.Param{repo, pullNum, command, dir}
	result := pegomock.GetGenericMockFrom(mock).Invoke("HidePrevCommandComments", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockClient) MarkdownPullLink(pull models.PullRequest) (string, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClient().")
	}
	params := []pegomock.Param{pull}
	result := pegomock.GetGenericMockFrom(mock).Invoke("MarkdownPullLink", params, []reflect.Type{reflect.TypeOf((*string)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 string
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(string)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockClient) MergePull(pull models.PullRequest, pullOptions models.PullRequestOptions) error {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClient().")
	}
	params := []pegomock.Param{pull, pullOptions}
	result := pegomock.GetGenericMockFrom(mock).Invoke("MergePull", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockClient) PullIsApproved(repo models.Repo, pull models.PullRequest) (models.ApprovalStatus, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClient().")
	}
	params := []pegomock.Param{repo, pull}
	result := pegomock.GetGenericMockFrom(mock).Invoke("PullIsApproved", params, []reflect.Type{reflect.TypeOf((*models.ApprovalStatus)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 models.ApprovalStatus
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(models.ApprovalStatus)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockClient) PullIsMergeable(repo models.Repo, pull models.PullRequest, vcsstatusname string) (bool, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClient().")
	}
	params := []pegomock.Param{repo, pull, vcsstatusname}
	result := pegomock.GetGenericMockFrom(mock).Invoke("PullIsMergeable", params, []reflect.Type{reflect.TypeOf((*bool)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 bool
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(bool)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockClient) ReactToComment(repo models.Repo, pullNum int, commentID int64, reaction string) error {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClient().")
	}
	params := []pegomock.Param{repo, pullNum, commentID, reaction}
	result := pegomock.GetGenericMockFrom(mock).Invoke("ReactToComment", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockClient) SupportsSingleFileDownload(repo models.Repo) bool {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClient().")
	}
	params := []pegomock.Param{repo}
	result := pegomock.GetGenericMockFrom(mock).Invoke("SupportsSingleFileDownload", params, []reflect.Type{reflect.TypeOf((*bool)(nil)).Elem()})
	var ret0 bool
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(bool)
		}
	}
	return ret0
}

func (mock *MockClient) UpdateStatus(repo models.Repo, pull models.PullRequest, state models.CommitStatus, src string, description string, url string) error {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClient().")
	}
	params := []pegomock.Param{repo, pull, state, src, description, url}
	result := pegomock.GetGenericMockFrom(mock).Invoke("UpdateStatus", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockClient) VerifyWasCalledOnce() *VerifierMockClient {
	return &VerifierMockClient{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockClient) VerifyWasCalled(invocationCountMatcher pegomock.InvocationCountMatcher) *VerifierMockClient {
	return &VerifierMockClient{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockClient) VerifyWasCalledInOrder(invocationCountMatcher pegomock.InvocationCountMatcher, inOrderContext *pegomock.InOrderContext) *VerifierMockClient {
	return &VerifierMockClient{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockClient) VerifyWasCalledEventually(invocationCountMatcher pegomock.InvocationCountMatcher, timeout time.Duration) *VerifierMockClient {
	return &VerifierMockClient{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierMockClient struct {
	mock                   *MockClient
	invocationCountMatcher pegomock.InvocationCountMatcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierMockClient) CreateComment(repo models.Repo, pullNum int, comment string, command string) *MockClient_CreateComment_OngoingVerification {
	params := []pegomock.Param{repo, pullNum, comment, command}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "CreateComment", params, verifier.timeout)
	return &MockClient_CreateComment_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockClient_CreateComment_OngoingVerification struct {
	mock              *MockClient
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockClient_CreateComment_OngoingVerification) GetCapturedArguments() (models.Repo, int, string, string) {
	repo, pullNum, comment, command := c.GetAllCapturedArguments()
	return repo[len(repo)-1], pullNum[len(pullNum)-1], comment[len(comment)-1], command[len(command)-1]
}

func (c *MockClient_CreateComment_OngoingVerification) GetAllCapturedArguments() (_param0 []models.Repo, _param1 []int, _param2 []string, _param3 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.Repo, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(models.Repo)
		}
		_param1 = make([]int, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(int)
		}
		_param2 = make([]string, len(c.methodInvocations))
		for u, param := range params[2] {
			_param2[u] = param.(string)
		}
		_param3 = make([]string, len(c.methodInvocations))
		for u, param := range params[3] {
			_param3[u] = param.(string)
		}
	}
	return
}

func (verifier *VerifierMockClient) DiscardReviews(repo models.Repo, pull models.PullRequest) *MockClient_DiscardReviews_OngoingVerification {
	params := []pegomock.Param{repo, pull}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "DiscardReviews", params, verifier.timeout)
	return &MockClient_DiscardReviews_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockClient_DiscardReviews_OngoingVerification struct {
	mock              *MockClient
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockClient_DiscardReviews_OngoingVerification) GetCapturedArguments() (models.Repo, models.PullRequest) {
	repo, pull := c.GetAllCapturedArguments()
	return repo[len(repo)-1], pull[len(pull)-1]
}

func (c *MockClient_DiscardReviews_OngoingVerification) GetAllCapturedArguments() (_param0 []models.Repo, _param1 []models.PullRequest) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.Repo, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(models.Repo)
		}
		_param1 = make([]models.PullRequest, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(models.PullRequest)
		}
	}
	return
}

func (verifier *VerifierMockClient) GetCloneURL(VCSHostType models.VCSHostType, repo string) *MockClient_GetCloneURL_OngoingVerification {
	params := []pegomock.Param{VCSHostType, repo}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "GetCloneURL", params, verifier.timeout)
	return &MockClient_GetCloneURL_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockClient_GetCloneURL_OngoingVerification struct {
	mock              *MockClient
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockClient_GetCloneURL_OngoingVerification) GetCapturedArguments() (models.VCSHostType, string) {
	VCSHostType, repo := c.GetAllCapturedArguments()
	return VCSHostType[len(VCSHostType)-1], repo[len(repo)-1]
}

func (c *MockClient_GetCloneURL_OngoingVerification) GetAllCapturedArguments() (_param0 []models.VCSHostType, _param1 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.VCSHostType, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(models.VCSHostType)
		}
		_param1 = make([]string, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(string)
		}
	}
	return
}

func (verifier *VerifierMockClient) GetFileContent(pull models.PullRequest, fileName string) *MockClient_GetFileContent_OngoingVerification {
	params := []pegomock.Param{pull, fileName}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "GetFileContent", params, verifier.timeout)
	return &MockClient_GetFileContent_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockClient_GetFileContent_OngoingVerification struct {
	mock              *MockClient
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockClient_GetFileContent_OngoingVerification) GetCapturedArguments() (models.PullRequest, string) {
	pull, fileName := c.GetAllCapturedArguments()
	return pull[len(pull)-1], fileName[len(fileName)-1]
}

func (c *MockClient_GetFileContent_OngoingVerification) GetAllCapturedArguments() (_param0 []models.PullRequest, _param1 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.PullRequest, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(models.PullRequest)
		}
		_param1 = make([]string, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(string)
		}
	}
	return
}

func (verifier *VerifierMockClient) GetModifiedFiles(repo models.Repo, pull models.PullRequest) *MockClient_GetModifiedFiles_OngoingVerification {
	params := []pegomock.Param{repo, pull}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "GetModifiedFiles", params, verifier.timeout)
	return &MockClient_GetModifiedFiles_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockClient_GetModifiedFiles_OngoingVerification struct {
	mock              *MockClient
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockClient_GetModifiedFiles_OngoingVerification) GetCapturedArguments() (models.Repo, models.PullRequest) {
	repo, pull := c.GetAllCapturedArguments()
	return repo[len(repo)-1], pull[len(pull)-1]
}

func (c *MockClient_GetModifiedFiles_OngoingVerification) GetAllCapturedArguments() (_param0 []models.Repo, _param1 []models.PullRequest) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.Repo, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(models.Repo)
		}
		_param1 = make([]models.PullRequest, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(models.PullRequest)
		}
	}
	return
}

func (verifier *VerifierMockClient) GetPullLabels(repo models.Repo, pull models.PullRequest) *MockClient_GetPullLabels_OngoingVerification {
	params := []pegomock.Param{repo, pull}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "GetPullLabels", params, verifier.timeout)
	return &MockClient_GetPullLabels_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockClient_GetPullLabels_OngoingVerification struct {
	mock              *MockClient
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockClient_GetPullLabels_OngoingVerification) GetCapturedArguments() (models.Repo, models.PullRequest) {
	repo, pull := c.GetAllCapturedArguments()
	return repo[len(repo)-1], pull[len(pull)-1]
}

func (c *MockClient_GetPullLabels_OngoingVerification) GetAllCapturedArguments() (_param0 []models.Repo, _param1 []models.PullRequest) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.Repo, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(models.Repo)
		}
		_param1 = make([]models.PullRequest, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(models.PullRequest)
		}
	}
	return
}

func (verifier *VerifierMockClient) GetTeamNamesForUser(repo models.Repo, user models.User) *MockClient_GetTeamNamesForUser_OngoingVerification {
	params := []pegomock.Param{repo, user}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "GetTeamNamesForUser", params, verifier.timeout)
	return &MockClient_GetTeamNamesForUser_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockClient_GetTeamNamesForUser_OngoingVerification struct {
	mock              *MockClient
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockClient_GetTeamNamesForUser_OngoingVerification) GetCapturedArguments() (models.Repo, models.User) {
	repo, user := c.GetAllCapturedArguments()
	return repo[len(repo)-1], user[len(user)-1]
}

func (c *MockClient_GetTeamNamesForUser_OngoingVerification) GetAllCapturedArguments() (_param0 []models.Repo, _param1 []models.User) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.Repo, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(models.Repo)
		}
		_param1 = make([]models.User, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(models.User)
		}
	}
	return
}

func (verifier *VerifierMockClient) HidePrevCommandComments(repo models.Repo, pullNum int, command string, dir string) *MockClient_HidePrevCommandComments_OngoingVerification {
	params := []pegomock.Param{repo, pullNum, command, dir}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "HidePrevCommandComments", params, verifier.timeout)
	return &MockClient_HidePrevCommandComments_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockClient_HidePrevCommandComments_OngoingVerification struct {
	mock              *MockClient
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockClient_HidePrevCommandComments_OngoingVerification) GetCapturedArguments() (models.Repo, int, string, string) {
	repo, pullNum, command, dir := c.GetAllCapturedArguments()
	return repo[len(repo)-1], pullNum[len(pullNum)-1], command[len(command)-1], dir[len(dir)-1]
}

func (c *MockClient_HidePrevCommandComments_OngoingVerification) GetAllCapturedArguments() (_param0 []models.Repo, _param1 []int, _param2 []string, _param3 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.Repo, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(models.Repo)
		}
		_param1 = make([]int, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(int)
		}
		_param2 = make([]string, len(c.methodInvocations))
		for u, param := range params[2] {
			_param2[u] = param.(string)
		}
		_param3 = make([]string, len(c.methodInvocations))
		for u, param := range params[3] {
			_param3[u] = param.(string)
		}
	}
	return
}

func (verifier *VerifierMockClient) MarkdownPullLink(pull models.PullRequest) *MockClient_MarkdownPullLink_OngoingVerification {
	params := []pegomock.Param{pull}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "MarkdownPullLink", params, verifier.timeout)
	return &MockClient_MarkdownPullLink_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockClient_MarkdownPullLink_OngoingVerification struct {
	mock              *MockClient
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockClient_MarkdownPullLink_OngoingVerification) GetCapturedArguments() models.PullRequest {
	pull := c.GetAllCapturedArguments()
	return pull[len(pull)-1]
}

func (c *MockClient_MarkdownPullLink_OngoingVerification) GetAllCapturedArguments() (_param0 []models.PullRequest) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.PullRequest, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(models.PullRequest)
		}
	}
	return
}

func (verifier *VerifierMockClient) MergePull(pull models.PullRequest, pullOptions models.PullRequestOptions) *MockClient_MergePull_OngoingVerification {
	params := []pegomock.Param{pull, pullOptions}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "MergePull", params, verifier.timeout)
	return &MockClient_MergePull_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockClient_MergePull_OngoingVerification struct {
	mock              *MockClient
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockClient_MergePull_OngoingVerification) GetCapturedArguments() (models.PullRequest, models.PullRequestOptions) {
	pull, pullOptions := c.GetAllCapturedArguments()
	return pull[len(pull)-1], pullOptions[len(pullOptions)-1]
}

func (c *MockClient_MergePull_OngoingVerification) GetAllCapturedArguments() (_param0 []models.PullRequest, _param1 []models.PullRequestOptions) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.PullRequest, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(models.PullRequest)
		}
		_param1 = make([]models.PullRequestOptions, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(models.PullRequestOptions)
		}
	}
	return
}

func (verifier *VerifierMockClient) PullIsApproved(repo models.Repo, pull models.PullRequest) *MockClient_PullIsApproved_OngoingVerification {
	params := []pegomock.Param{repo, pull}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "PullIsApproved", params, verifier.timeout)
	return &MockClient_PullIsApproved_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockClient_PullIsApproved_OngoingVerification struct {
	mock              *MockClient
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockClient_PullIsApproved_OngoingVerification) GetCapturedArguments() (models.Repo, models.PullRequest) {
	repo, pull := c.GetAllCapturedArguments()
	return repo[len(repo)-1], pull[len(pull)-1]
}

func (c *MockClient_PullIsApproved_OngoingVerification) GetAllCapturedArguments() (_param0 []models.Repo, _param1 []models.PullRequest) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.Repo, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(models.Repo)
		}
		_param1 = make([]models.PullRequest, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(models.PullRequest)
		}
	}
	return
}

func (verifier *VerifierMockClient) PullIsMergeable(repo models.Repo, pull models.PullRequest, vcsstatusname string) *MockClient_PullIsMergeable_OngoingVerification {
	params := []pegomock.Param{repo, pull, vcsstatusname}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "PullIsMergeable", params, verifier.timeout)
	return &MockClient_PullIsMergeable_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockClient_PullIsMergeable_OngoingVerification struct {
	mock              *MockClient
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockClient_PullIsMergeable_OngoingVerification) GetCapturedArguments() (models.Repo, models.PullRequest, string) {
	repo, pull, vcsstatusname := c.GetAllCapturedArguments()
	return repo[len(repo)-1], pull[len(pull)-1], vcsstatusname[len(vcsstatusname)-1]
}

func (c *MockClient_PullIsMergeable_OngoingVerification) GetAllCapturedArguments() (_param0 []models.Repo, _param1 []models.PullRequest, _param2 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.Repo, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(models.Repo)
		}
		_param1 = make([]models.PullRequest, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(models.PullRequest)
		}
		_param2 = make([]string, len(c.methodInvocations))
		for u, param := range params[2] {
			_param2[u] = param.(string)
		}
	}
	return
}

func (verifier *VerifierMockClient) ReactToComment(repo models.Repo, pullNum int, commentID int64, reaction string) *MockClient_ReactToComment_OngoingVerification {
	params := []pegomock.Param{repo, pullNum, commentID, reaction}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "ReactToComment", params, verifier.timeout)
	return &MockClient_ReactToComment_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockClient_ReactToComment_OngoingVerification struct {
	mock              *MockClient
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockClient_ReactToComment_OngoingVerification) GetCapturedArguments() (models.Repo, int, int64, string) {
	repo, pullNum, commentID, reaction := c.GetAllCapturedArguments()
	return repo[len(repo)-1], pullNum[len(pullNum)-1], commentID[len(commentID)-1], reaction[len(reaction)-1]
}

func (c *MockClient_ReactToComment_OngoingVerification) GetAllCapturedArguments() (_param0 []models.Repo, _param1 []int, _param2 []int64, _param3 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.Repo, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(models.Repo)
		}
		_param1 = make([]int, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(int)
		}
		_param2 = make([]int64, len(c.methodInvocations))
		for u, param := range params[2] {
			_param2[u] = param.(int64)
		}
		_param3 = make([]string, len(c.methodInvocations))
		for u, param := range params[3] {
			_param3[u] = param.(string)
		}
	}
	return
}

func (verifier *VerifierMockClient) SupportsSingleFileDownload(repo models.Repo) *MockClient_SupportsSingleFileDownload_OngoingVerification {
	params := []pegomock.Param{repo}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "SupportsSingleFileDownload", params, verifier.timeout)
	return &MockClient_SupportsSingleFileDownload_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockClient_SupportsSingleFileDownload_OngoingVerification struct {
	mock              *MockClient
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockClient_SupportsSingleFileDownload_OngoingVerification) GetCapturedArguments() models.Repo {
	repo := c.GetAllCapturedArguments()
	return repo[len(repo)-1]
}

func (c *MockClient_SupportsSingleFileDownload_OngoingVerification) GetAllCapturedArguments() (_param0 []models.Repo) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.Repo, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(models.Repo)
		}
	}
	return
}

func (verifier *VerifierMockClient) UpdateStatus(repo models.Repo, pull models.PullRequest, state models.CommitStatus, src string, description string, url string) *MockClient_UpdateStatus_OngoingVerification {
	params := []pegomock.Param{repo, pull, state, src, description, url}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "UpdateStatus", params, verifier.timeout)
	return &MockClient_UpdateStatus_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockClient_UpdateStatus_OngoingVerification struct {
	mock              *MockClient
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockClient_UpdateStatus_OngoingVerification) GetCapturedArguments() (models.Repo, models.PullRequest, models.CommitStatus, string, string, string) {
	repo, pull, state, src, description, url := c.GetAllCapturedArguments()
	return repo[len(repo)-1], pull[len(pull)-1], state[len(state)-1], src[len(src)-1], description[len(description)-1], url[len(url)-1]
}

func (c *MockClient_UpdateStatus_OngoingVerification) GetAllCapturedArguments() (_param0 []models.Repo, _param1 []models.PullRequest, _param2 []models.CommitStatus, _param3 []string, _param4 []string, _param5 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.Repo, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(models.Repo)
		}
		_param1 = make([]models.PullRequest, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(models.PullRequest)
		}
		_param2 = make([]models.CommitStatus, len(c.methodInvocations))
		for u, param := range params[2] {
			_param2[u] = param.(models.CommitStatus)
		}
		_param3 = make([]string, len(c.methodInvocations))
		for u, param := range params[3] {
			_param3[u] = param.(string)
		}
		_param4 = make([]string, len(c.methodInvocations))
		for u, param := range params[4] {
			_param4[u] = param.(string)
		}
		_param5 = make([]string, len(c.methodInvocations))
		for u, param := range params[5] {
			_param5[u] = param.(string)
		}
	}
	return
}
