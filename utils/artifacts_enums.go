package utils

const (
	JobArtifactUserSpace   = "user"
	JobArtifactSystemSpace = "system"
)

func GetJobArtifactUserSpace() *string {
	space := JobArtifactUserSpace
	return &space
}

func GetJobArtifactSystemSpace() *string {
	space := JobArtifactSystemSpace
	return &space
}
