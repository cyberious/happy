package workspace_repo

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/chanzuckerberg/happy/cli/pkg/diagnostics"
	"github.com/chanzuckerberg/happy/cli/pkg/options"
	"github.com/chanzuckerberg/happy/cli/pkg/util"
	"github.com/docker/go-units"
	"github.com/hashicorp/go-tfe"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const alertAfter = 300 * time.Second

// implements the Workspace interface
type TFEWorkspace struct {
	tfc          *tfe.Client
	workspace    *tfe.Workspace
	outputs      map[string]string
	vars         map[string]map[string]*tfe.Variable
	currentRun   *tfe.Run
	currentRunID string
}

// For testing purposes only
func (s *TFEWorkspace) SetClient(tfc *tfe.Client) {
	s.tfc = tfc
}

// For testing purposes only
func (s *TFEWorkspace) SetWorkspace(workspace *tfe.Workspace) {
	s.workspace = workspace
}

func (s *TFEWorkspace) GetWorkspaceID() string {
	return s.workspace.ID
}

func (s *TFEWorkspace) GetWorkspaceName() string {
	return s.workspace.Name
}

func (s *TFEWorkspace) GetWorkspaceOrganizationName() string {
	if s.workspace.Organization == nil {
		return ""
	}
	return s.workspace.Organization.Name
}

func (s *TFEWorkspace) GetCurrentRunID() string {
	if s.currentRunID == "" {
		currentRun := s.workspace.CurrentRun
		if currentRun != nil {
			s.currentRunID = currentRun.ID
		}
	}
	return s.currentRunID
}

func (s *TFEWorkspace) getCurrentRun() (*tfe.Run, error) {
	if s.currentRun != nil {
		return s.currentRun, nil
	}

	if s.GetCurrentRunID() == "" {
		return nil, errors.Errorf("fail to get current Run for %s: Run ID is empty", s.WorkspaceName())
	}

	currentRun, err := s.tfc.Runs.Read(context.Background(), s.GetCurrentRunID())
	if err != nil {
		return nil, errors.Wrap(err, "could not get tfe run")
	}
	s.currentRun = currentRun
	return s.currentRun, nil
}

func (s *TFEWorkspace) DiscardRun(ctx context.Context, runID string) error {
	if len(runID) == 0 {
		return errors.New("no run to discard")
	}
	return s.tfc.Runs.Discard(ctx, runID, tfe.RunDiscardOptions{
		Comment: tfe.String("cancelled by happy"),
	})
}

func (s *TFEWorkspace) GetLatestConfigVersionID() (string, error) {
	currentRun, err := s.getCurrentRun()
	if err != nil {
		return "", errors.Wrap(err, "failed to get the lastest ConfigVersion ID")
	}

	return currentRun.ConfigurationVersion.ID, nil
}

func (s *TFEWorkspace) Run(isDestroy bool, dryRun util.DryRunType) error {
	logrus.Infof("running workspace %s ...", s.workspace.Name)
	lastConfigVersionId, err := s.GetLatestConfigVersionID()
	if err != nil {
		return err
	}
	err = s.RunConfigVersion(lastConfigVersionId, isDestroy, dryRun)
	if err != nil {
		return err
	}

	return nil
}

func (s *TFEWorkspace) HasState(ctx context.Context) (bool, error) {
	options := tfe.StateVersionListOptions{
		ListOptions: tfe.ListOptions{
			PageNumber: 0,
			PageSize:   10,
		},
		Organization: s.GetWorkspaceOrganizationName(),
		Workspace:    s.WorkspaceName(),
	}
	list, err := s.tfc.StateVersions.List(ctx, &options)
	if err != nil {
		if errors.Is(err, tfe.ErrResourceNotFound) {
			return false, nil
		}
		return true, err
	}
	return len(list.Items) > 0, nil
}

func (s *TFEWorkspace) getVars() (map[string]map[string]*tfe.Variable, error) {
	if s.vars == nil {
		workspaceVars, err := s.tfc.Variables.List(context.Background(), s.GetWorkspaceId(), &tfe.VariableListOptions{})
		if err != nil {
			return nil, errors.Errorf("failed to get workspace vars: %v", err)
		}

		s.vars = map[string]map[string]*tfe.Variable{}
		for _, varItem := range workspaceVars.Items {
			category := string(varItem.Category)
			if _, ok := s.vars[category]; !ok {
				innerMap := map[string]*tfe.Variable{}
				s.vars[category] = innerMap
			}
			s.vars[category][varItem.Key] = varItem
		}
	}

	return s.vars, nil
}

func (s *TFEWorkspace) WorkspaceName() string {
	return s.workspace.Name
}

func (s *TFEWorkspace) SetVars(key string, value string, description string, sensitive bool) error {
	category := "terraform" // Hard-coded, not allowing setting environment vars directly
	isHCL := false

	variableMap, ok := s.vars[category]
	variable, variableOK := variableMap[key]
	update := ok && variableOK
	if update {
		logrus.WithField("existing vars", s.vars).
			WithField("tfe_workspace_id", s.GetWorkspaceID()).
			Debugf("tfe attempting to update variable %s:%s", key, value)
		options := tfe.VariableUpdateOptions{
			Type:        "vars",
			Key:         &key,
			Value:       &value,
			Description: &description,
			HCL:         &isHCL,
			Sensitive:   &sensitive,
		}
		_, err := s.tfc.Variables.Update(context.Background(), s.GetWorkspaceID(), variable.ID, options)
		return errors.Wrapf(err, "could not update TFE variable %s:%s", key, value)
	}

	logrus.WithField("existing vars", s.vars).
		WithField("tfe_workspace_id", s.GetWorkspaceID()).
		Debugf("tfe attempting to create variable %s:%s", key, value)

	options := tfe.VariableCreateOptions{
		Type:        "vars",
		Key:         &key,
		Value:       &value,
		Description: &description,
		Category:    tfe.Category(tfe.CategoryType(category)),
		HCL:         &isHCL,
		Sensitive:   &sensitive,
	}
	if util.IsLocalstackMode() {
		return nil
	}
	_, err := s.tfc.Variables.Create(context.Background(), s.GetWorkspaceID(), options)
	return errors.Wrapf(err, "could not create TFE variable %s:%s", key, value)
}

func (s *TFEWorkspace) RunConfigVersion(configVersionId string, isDestroy bool, dryRun util.DryRunType) error {
	// TODO: say who queued this or give more contextual info
	logrus.Debugf("version ID: %s, idDestroy: %t", configVersionId, isDestroy)

	option := tfe.RunCreateOptions{
		Type:      "runs",
		IsDestroy: &isDestroy,
		Message:   tfe.String("Queued from happy cli"),
		ConfigurationVersion: &tfe.ConfigurationVersion{
			ID:          configVersionId,
			Speculative: bool(dryRun),
		},
		Workspace: &tfe.Workspace{
			ID: s.GetWorkspaceID(),
		},
		TargetAddrs: []string{},
	}
	if bool(dryRun) && isDestroy {
		option.AutoApply = tfe.Bool(false)
	}
	run, err := s.tfc.Runs.Create(context.Background(), option)
	if err != nil {
		return err
	}
	// the run just created is the current run
	s.currentRunID = run.ID
	s.currentRun = nil
	s.outputs = nil
	return nil
}

func (s *TFEWorkspace) Wait(ctx context.Context, dryRun util.DryRunType) error {
	return s.WaitWithOptions(ctx, options.WaitOptions{}, dryRun)
}

func (s *TFEWorkspace) WaitWithOptions(ctx context.Context, waitOptions options.WaitOptions, dryRun util.DryRunType) error {
	RunDoneStatuses := map[tfe.RunStatus]bool{
		tfe.RunApplied:            true,
		tfe.RunDiscarded:          true,
		tfe.RunErrored:            true,
		tfe.RunCanceled:           true,
		tfe.RunPolicySoftFailed:   true,
		tfe.RunPlannedAndFinished: true,
	}

	TfeSuccessStatuses := map[tfe.RunStatus]struct{}{
		tfe.RunApplied:            {},
		tfe.RunPlannedAndFinished: {},
	}

	if dryRun {
		RunDoneStatuses = map[tfe.RunStatus]bool{
			tfe.RunDiscarded:          true,
			tfe.RunErrored:            true,
			tfe.RunCanceled:           true,
			tfe.RunPolicyChecked:      true,
			tfe.RunPlannedAndFinished: true,
		}

		TfeSuccessStatuses = map[tfe.RunStatus]struct{}{
			tfe.RunPlanned:            {},
			tfe.RunPolicyChecked:      {},
			tfe.RunPlannedAndFinished: {},
		}
	}

	diagnostics.AddTfeRunInfoOrg(ctx, s.GetWorkspaceOrganizationName())
	diagnostics.AddTfeRunInfoWorkspace(ctx, s.GetWorkspaceName())
	diagnostics.AddTfeRunInfoRunId(ctx, s.GetCurrentRunID())
	diagnostics.PrintTfeRunLink(ctx)

	startTimestamp := time.Now()
	printedAlert := false

	var sentinelStatus tfe.RunStatus = ""
	lastStatus := sentinelStatus

	done := false

	logCtx, cancel := context.WithCancel(ctx)
	defer cancel()
	for done = false; !done; _, done = RunDoneStatuses[lastStatus] {
		if lastStatus != sentinelStatus {
			time.Sleep(5 * time.Second)
		}
		run, err := s.tfc.Runs.Read(ctx, s.GetCurrentRunID())
		if err != nil {
			return err
		}
		status := run.Status

		if waitOptions.Orchestrator != nil && !printedAlert && len(waitOptions.StackName) > 0 && time.Since(startTimestamp) > alertAfter {
			// TODO(el): A more helpful message
			logrus.Warn("This apply is taking an unusually long time. Are your containers crashing?")
			err = waitOptions.Orchestrator.GetEvents(ctx, waitOptions.StackName, waitOptions.Services)
			if err != nil {
				return err
			}
			printedAlert = true
		}

		if status != lastStatus {
			elapsed := time.Since(startTimestamp)
			logrus.Infof("[%s] -> [%s]: %s elapsed", lastStatus, status, units.HumanDuration(elapsed))
			lastStatus = status

			if status == tfe.RunPlanning {
				if run.Plan != nil && len(run.Plan.ID) > 0 {
					logs, err := s.tfc.Plans.Logs(logCtx, run.Plan.ID)
					if err != nil {
						logrus.Errorf("cannot retrieve logs: %s", err.Error())
					} else {
						go s.streamLogs(logCtx, logs)
					}
				}
			}

			if status == tfe.RunApplying {
				if run.Apply != nil && len(run.Apply.ID) > 0 {
					logs, err := s.tfc.Applies.Logs(logCtx, run.Apply.ID)
					if err != nil {
						logrus.Errorf("cannot retrieve logs: %s", err.Error())
					} else {
						go s.streamLogs(logCtx, logs)
					}
				}
			}
		}
	}

	_, success := TfeSuccessStatuses[lastStatus]
	if !success {
		return errors.Errorf("error applying, ended in status %s", lastStatus)
	}

	return nil
}

func (s *TFEWorkspace) streamLogs(ctx context.Context, logs io.Reader) {
	// NOTE: in certain contexts
	// we don't want to show these unless specifically requested
	logfunc := logrus.Info
	if util.IsCI(ctx) {
		logfunc = logrus.Debug
	}

	logfunc("...streaming logs...")
	scanner := bufio.NewScanner(logs)
	for scanner.Scan() {
		select {
		case <-ctx.Done():
			logfunc("...log stream cancelled...")
			return
		default:
			logfunc(string(scanner.Text()))
		}
	}
	if err := scanner.Err(); err != nil {
		if !errors.Is(err, context.Canceled) && !errors.Is(err, io.EOF) {
			logrus.Errorf("...log stream error: %s...", err.Error())
			return
		}
	}
	logrus.Info("...log stream ended...")
}

func (s *TFEWorkspace) ResetCache() {
	s.vars = nil
	s.outputs = nil
	s.currentRun = nil
}

func (s *TFEWorkspace) GetTags() (map[string]string, error) {
	tags := map[string]string{}

	vars, err := s.getVars()
	if err != nil {
		return nil, err
	}

	terraformVars, ok := vars["terraform"]
	if !ok {
		return tags, nil
	}

	happyMetaVar, ok := terraformVars["happymeta_"]
	if !ok {
		return tags, nil
	}

	if happyMetaVar.Sensitive {
		return nil, errors.Errorf("invalid meta var for stack %s, must not be sensitive", s.workspace.Name)
	}

	// Timestamp tags come back as numeric values, and cannot be deserialized into map[string]string; code below
	// converts float64 to string, all other non-string value types will be converted.
	allTags := map[string]interface{}{}
	err = json.Unmarshal([]byte(happyMetaVar.Value), &allTags)
	for tag, value := range allTags {
		tags[tag] = util.TagValueToString(value)
	}
	return tags, errors.Wrap(err, "could not parse json")
}

func (s *TFEWorkspace) GetWorkspaceId() string {
	return s.workspace.ID
}

// For testing purposes only
func (s *TFEWorkspace) SetOutputs(outputs map[string]string) {
	s.outputs = outputs
}

func (s *TFEWorkspace) GetOutputs() (map[string]string, error) {
	if s.outputs != nil {
		return s.outputs, nil
	}

	s.outputs = map[string]string{}
	stateVersion, err := s.tfc.StateVersions.ReadCurrentWithOptions(context.Background(), s.GetWorkspaceId(), &tfe.StateVersionCurrentOptions{Include: []tfe.StateVersionIncludeOpt{"outputs"}})
	if err != nil {
		return nil, errors.Errorf("failed to get state for workspace %s", s.GetWorkspaceID())
	}

	var svOutputIDs []string
	for _, svOutput := range stateVersion.Outputs {
		svOutputIDs = append(svOutputIDs, svOutput.ID)
	}

	for _, svOutputID := range svOutputIDs {
		svOutput, err := s.tfc.StateVersionOutputs.Read(context.Background(), svOutputID)
		if err != nil {
			return nil, errors.Wrap(err, "could not read state version outputs")
		}

		if !svOutput.Sensitive {
			bytes, err := json.MarshalIndent(svOutput.Value, "", "\t")
			if err != nil {
				s.outputs[svOutput.Name] = fmt.Sprintf("%v", svOutput.Value)
			} else {
				s.outputs[svOutput.Name] = string(bytes)
			}
		}
	}

	return s.outputs, nil
}

func (s *TFEWorkspace) GetCurrentRunStatus() string {
	if s.currentRun == nil {
		currentRun, err := s.tfc.Runs.Read(context.Background(), s.workspace.CurrentRun.ID)
		if err != nil {
			return ""
		}
		s.currentRun = currentRun
	}
	return string(s.currentRun.Status)
}

// create a new ConfigurationVersion in a TFE workspace, upload the targz file to
// the new ConfigurationVersion, and finally return its ID.
func (s *TFEWorkspace) UploadVersion(targzFilePath string, dryRun util.DryRunType) (string, error) {
	autoQueueRun := false
	options := tfe.ConfigurationVersionCreateOptions{
		Type:          "configuration-versions",
		AutoQueueRuns: &autoQueueRun,
		Speculative:   tfe.Bool(bool(dryRun)),
	}
	configVersion, err := s.tfc.ConfigurationVersions.Create(context.Background(), s.GetWorkspaceID(), options)
	if err != nil {
		return "", err
	}
	if err := s.tfc.ConfigurationVersions.Upload(context.Background(), configVersion.UploadURL, targzFilePath); err != nil {
		return "", err
	}
	return configVersion.ID, nil
}
