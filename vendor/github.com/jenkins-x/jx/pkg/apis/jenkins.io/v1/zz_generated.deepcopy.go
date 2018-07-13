// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	rbac_v1 "k8s.io/api/rbac/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommitSummary) DeepCopyInto(out *CommitSummary) {
	*out = *in
	if in.Author != nil {
		in, out := &in.Author, &out.Author
		if *in == nil {
			*out = nil
		} else {
			*out = new(UserDetails)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Committer != nil {
		in, out := &in.Committer, &out.Committer
		if *in == nil {
			*out = nil
		} else {
			*out = new(UserDetails)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.IssueIDs != nil {
		in, out := &in.IssueIDs, &out.IssueIDs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommitSummary.
func (in *CommitSummary) DeepCopy() *CommitSummary {
	if in == nil {
		return nil
	}
	out := new(CommitSummary)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CoreActivityStep) DeepCopyInto(out *CoreActivityStep) {
	*out = *in
	if in.StartedTimestamp != nil {
		in, out := &in.StartedTimestamp, &out.StartedTimestamp
		if *in == nil {
			*out = nil
		} else {
			*out = (*in).DeepCopy()
		}
	}
	if in.CompletedTimestamp != nil {
		in, out := &in.CompletedTimestamp, &out.CompletedTimestamp
		if *in == nil {
			*out = nil
		} else {
			*out = (*in).DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CoreActivityStep.
func (in *CoreActivityStep) DeepCopy() *CoreActivityStep {
	if in == nil {
		return nil
	}
	out := new(CoreActivityStep)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Environment) DeepCopyInto(out *Environment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Environment.
func (in *Environment) DeepCopy() *Environment {
	if in == nil {
		return nil
	}
	out := new(Environment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Environment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentFilter) DeepCopyInto(out *EnvironmentFilter) {
	*out = *in
	if in.Includes != nil {
		in, out := &in.Includes, &out.Includes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Excludes != nil {
		in, out := &in.Excludes, &out.Excludes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentFilter.
func (in *EnvironmentFilter) DeepCopy() *EnvironmentFilter {
	if in == nil {
		return nil
	}
	out := new(EnvironmentFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentList) DeepCopyInto(out *EnvironmentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Environment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentList.
func (in *EnvironmentList) DeepCopy() *EnvironmentList {
	if in == nil {
		return nil
	}
	out := new(EnvironmentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EnvironmentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentRepository) DeepCopyInto(out *EnvironmentRepository) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentRepository.
func (in *EnvironmentRepository) DeepCopy() *EnvironmentRepository {
	if in == nil {
		return nil
	}
	out := new(EnvironmentRepository)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentRoleBinding) DeepCopyInto(out *EnvironmentRoleBinding) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentRoleBinding.
func (in *EnvironmentRoleBinding) DeepCopy() *EnvironmentRoleBinding {
	if in == nil {
		return nil
	}
	out := new(EnvironmentRoleBinding)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EnvironmentRoleBinding) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentRoleBindingList) DeepCopyInto(out *EnvironmentRoleBindingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EnvironmentRoleBinding, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentRoleBindingList.
func (in *EnvironmentRoleBindingList) DeepCopy() *EnvironmentRoleBindingList {
	if in == nil {
		return nil
	}
	out := new(EnvironmentRoleBindingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EnvironmentRoleBindingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentRoleBindingSpec) DeepCopyInto(out *EnvironmentRoleBindingSpec) {
	*out = *in
	if in.Subjects != nil {
		in, out := &in.Subjects, &out.Subjects
		*out = make([]rbac_v1.Subject, len(*in))
		copy(*out, *in)
	}
	out.RoleRef = in.RoleRef
	if in.Environments != nil {
		in, out := &in.Environments, &out.Environments
		*out = make([]EnvironmentFilter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentRoleBindingSpec.
func (in *EnvironmentRoleBindingSpec) DeepCopy() *EnvironmentRoleBindingSpec {
	if in == nil {
		return nil
	}
	out := new(EnvironmentRoleBindingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentRoleBindingStatus) DeepCopyInto(out *EnvironmentRoleBindingStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentRoleBindingStatus.
func (in *EnvironmentRoleBindingStatus) DeepCopy() *EnvironmentRoleBindingStatus {
	if in == nil {
		return nil
	}
	out := new(EnvironmentRoleBindingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentSpec) DeepCopyInto(out *EnvironmentSpec) {
	*out = *in
	out.Source = in.Source
	in.TeamSettings.DeepCopyInto(&out.TeamSettings)
	out.PreviewGitSpec = in.PreviewGitSpec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentSpec.
func (in *EnvironmentSpec) DeepCopy() *EnvironmentSpec {
	if in == nil {
		return nil
	}
	out := new(EnvironmentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentStatus) DeepCopyInto(out *EnvironmentStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentStatus.
func (in *EnvironmentStatus) DeepCopy() *EnvironmentStatus {
	if in == nil {
		return nil
	}
	out := new(EnvironmentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitService) DeepCopyInto(out *GitService) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitService.
func (in *GitService) DeepCopy() *GitService {
	if in == nil {
		return nil
	}
	out := new(GitService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GitService) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitServiceList) DeepCopyInto(out *GitServiceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GitService, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitServiceList.
func (in *GitServiceList) DeepCopy() *GitServiceList {
	if in == nil {
		return nil
	}
	out := new(GitServiceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GitServiceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitServiceSpec) DeepCopyInto(out *GitServiceSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitServiceSpec.
func (in *GitServiceSpec) DeepCopy() *GitServiceSpec {
	if in == nil {
		return nil
	}
	out := new(GitServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitStatus) DeepCopyInto(out *GitStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitStatus.
func (in *GitStatus) DeepCopy() *GitStatus {
	if in == nil {
		return nil
	}
	out := new(GitStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IssueLabel) DeepCopyInto(out *IssueLabel) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IssueLabel.
func (in *IssueLabel) DeepCopy() *IssueLabel {
	if in == nil {
		return nil
	}
	out := new(IssueLabel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IssueSummary) DeepCopyInto(out *IssueSummary) {
	*out = *in
	if in.User != nil {
		in, out := &in.User, &out.User
		if *in == nil {
			*out = nil
		} else {
			*out = new(UserDetails)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Assignees != nil {
		in, out := &in.Assignees, &out.Assignees
		*out = make([]UserDetails, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ClosedBy != nil {
		in, out := &in.ClosedBy, &out.ClosedBy
		if *in == nil {
			*out = nil
		} else {
			*out = new(UserDetails)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.CreationTimestamp != nil {
		in, out := &in.CreationTimestamp, &out.CreationTimestamp
		if *in == nil {
			*out = nil
		} else {
			*out = (*in).DeepCopy()
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make([]IssueLabel, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IssueSummary.
func (in *IssueSummary) DeepCopy() *IssueSummary {
	if in == nil {
		return nil
	}
	out := new(IssueSummary)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineActivity) DeepCopyInto(out *PipelineActivity) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineActivity.
func (in *PipelineActivity) DeepCopy() *PipelineActivity {
	if in == nil {
		return nil
	}
	out := new(PipelineActivity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PipelineActivity) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineActivityList) DeepCopyInto(out *PipelineActivityList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PipelineActivity, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineActivityList.
func (in *PipelineActivityList) DeepCopy() *PipelineActivityList {
	if in == nil {
		return nil
	}
	out := new(PipelineActivityList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PipelineActivityList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineActivitySpec) DeepCopyInto(out *PipelineActivitySpec) {
	*out = *in
	if in.StartedTimestamp != nil {
		in, out := &in.StartedTimestamp, &out.StartedTimestamp
		if *in == nil {
			*out = nil
		} else {
			*out = (*in).DeepCopy()
		}
	}
	if in.CompletedTimestamp != nil {
		in, out := &in.CompletedTimestamp, &out.CompletedTimestamp
		if *in == nil {
			*out = nil
		} else {
			*out = (*in).DeepCopy()
		}
	}
	if in.Steps != nil {
		in, out := &in.Steps, &out.Steps
		*out = make([]PipelineActivityStep, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineActivitySpec.
func (in *PipelineActivitySpec) DeepCopy() *PipelineActivitySpec {
	if in == nil {
		return nil
	}
	out := new(PipelineActivitySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineActivityStatus) DeepCopyInto(out *PipelineActivityStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineActivityStatus.
func (in *PipelineActivityStatus) DeepCopy() *PipelineActivityStatus {
	if in == nil {
		return nil
	}
	out := new(PipelineActivityStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineActivityStep) DeepCopyInto(out *PipelineActivityStep) {
	*out = *in
	if in.Stage != nil {
		in, out := &in.Stage, &out.Stage
		if *in == nil {
			*out = nil
		} else {
			*out = new(StageActivityStep)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Promote != nil {
		in, out := &in.Promote, &out.Promote
		if *in == nil {
			*out = nil
		} else {
			*out = new(PromoteActivityStep)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Preview != nil {
		in, out := &in.Preview, &out.Preview
		if *in == nil {
			*out = nil
		} else {
			*out = new(PreviewActivityStep)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineActivityStep.
func (in *PipelineActivityStep) DeepCopy() *PipelineActivityStep {
	if in == nil {
		return nil
	}
	out := new(PipelineActivityStep)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PreviewActivityStep) DeepCopyInto(out *PreviewActivityStep) {
	*out = *in
	in.CoreActivityStep.DeepCopyInto(&out.CoreActivityStep)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PreviewActivityStep.
func (in *PreviewActivityStep) DeepCopy() *PreviewActivityStep {
	if in == nil {
		return nil
	}
	out := new(PreviewActivityStep)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PreviewGitSpec) DeepCopyInto(out *PreviewGitSpec) {
	*out = *in
	out.User = in.User
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PreviewGitSpec.
func (in *PreviewGitSpec) DeepCopy() *PreviewGitSpec {
	if in == nil {
		return nil
	}
	out := new(PreviewGitSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PromoteActivityStep) DeepCopyInto(out *PromoteActivityStep) {
	*out = *in
	in.CoreActivityStep.DeepCopyInto(&out.CoreActivityStep)
	if in.PullRequest != nil {
		in, out := &in.PullRequest, &out.PullRequest
		if *in == nil {
			*out = nil
		} else {
			*out = new(PromotePullRequestStep)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Update != nil {
		in, out := &in.Update, &out.Update
		if *in == nil {
			*out = nil
		} else {
			*out = new(PromoteUpdateStep)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PromoteActivityStep.
func (in *PromoteActivityStep) DeepCopy() *PromoteActivityStep {
	if in == nil {
		return nil
	}
	out := new(PromoteActivityStep)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PromotePullRequestStep) DeepCopyInto(out *PromotePullRequestStep) {
	*out = *in
	in.CoreActivityStep.DeepCopyInto(&out.CoreActivityStep)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PromotePullRequestStep.
func (in *PromotePullRequestStep) DeepCopy() *PromotePullRequestStep {
	if in == nil {
		return nil
	}
	out := new(PromotePullRequestStep)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PromoteUpdateStep) DeepCopyInto(out *PromoteUpdateStep) {
	*out = *in
	in.CoreActivityStep.DeepCopyInto(&out.CoreActivityStep)
	if in.Statuses != nil {
		in, out := &in.Statuses, &out.Statuses
		*out = make([]GitStatus, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PromoteUpdateStep.
func (in *PromoteUpdateStep) DeepCopy() *PromoteUpdateStep {
	if in == nil {
		return nil
	}
	out := new(PromoteUpdateStep)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QuickStartLocation) DeepCopyInto(out *QuickStartLocation) {
	*out = *in
	if in.Includes != nil {
		in, out := &in.Includes, &out.Includes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Excludes != nil {
		in, out := &in.Excludes, &out.Excludes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QuickStartLocation.
func (in *QuickStartLocation) DeepCopy() *QuickStartLocation {
	if in == nil {
		return nil
	}
	out := new(QuickStartLocation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Release) DeepCopyInto(out *Release) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Release.
func (in *Release) DeepCopy() *Release {
	if in == nil {
		return nil
	}
	out := new(Release)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Release) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleaseList) DeepCopyInto(out *ReleaseList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Release, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleaseList.
func (in *ReleaseList) DeepCopy() *ReleaseList {
	if in == nil {
		return nil
	}
	out := new(ReleaseList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReleaseList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleaseSpec) DeepCopyInto(out *ReleaseSpec) {
	*out = *in
	if in.Commits != nil {
		in, out := &in.Commits, &out.Commits
		*out = make([]CommitSummary, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Issues != nil {
		in, out := &in.Issues, &out.Issues
		*out = make([]IssueSummary, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PullRequests != nil {
		in, out := &in.PullRequests, &out.PullRequests
		*out = make([]IssueSummary, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleaseSpec.
func (in *ReleaseSpec) DeepCopy() *ReleaseSpec {
	if in == nil {
		return nil
	}
	out := new(ReleaseSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleaseStatus) DeepCopyInto(out *ReleaseStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleaseStatus.
func (in *ReleaseStatus) DeepCopy() *ReleaseStatus {
	if in == nil {
		return nil
	}
	out := new(ReleaseStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StageActivityStep) DeepCopyInto(out *StageActivityStep) {
	*out = *in
	in.CoreActivityStep.DeepCopyInto(&out.CoreActivityStep)
	if in.Steps != nil {
		in, out := &in.Steps, &out.Steps
		*out = make([]CoreActivityStep, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StageActivityStep.
func (in *StageActivityStep) DeepCopy() *StageActivityStep {
	if in == nil {
		return nil
	}
	out := new(StageActivityStep)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TeamSettings) DeepCopyInto(out *TeamSettings) {
	*out = *in
	if in.QuickstartLocations != nil {
		in, out := &in.QuickstartLocations, &out.QuickstartLocations
		*out = make([]QuickStartLocation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TeamSettings.
func (in *TeamSettings) DeepCopy() *TeamSettings {
	if in == nil {
		return nil
	}
	out := new(TeamSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *User) DeepCopyInto(out *User) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.User.DeepCopyInto(&out.User)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new User.
func (in *User) DeepCopy() *User {
	if in == nil {
		return nil
	}
	out := new(User)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *User) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserDetails) DeepCopyInto(out *UserDetails) {
	*out = *in
	if in.CreationTimestamp != nil {
		in, out := &in.CreationTimestamp, &out.CreationTimestamp
		if *in == nil {
			*out = nil
		} else {
			*out = (*in).DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserDetails.
func (in *UserDetails) DeepCopy() *UserDetails {
	if in == nil {
		return nil
	}
	out := new(UserDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserList) DeepCopyInto(out *UserList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]User, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserList.
func (in *UserList) DeepCopy() *UserList {
	if in == nil {
		return nil
	}
	out := new(UserList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UserList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserSpec) DeepCopyInto(out *UserSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserSpec.
func (in *UserSpec) DeepCopy() *UserSpec {
	if in == nil {
		return nil
	}
	out := new(UserSpec)
	in.DeepCopyInto(out)
	return out
}
