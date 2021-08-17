package ebs

// StubArtifact is an artifact implementation that is empty but valid
type StubArtifact struct {
	// BuilderId is the unique ID for the builder that created this artifact
	BuilderIdValue string
}

func (a *StubArtifact) BuilderId() string {
	return a.BuilderIdValue
}

func (*StubArtifact) Files() []string {
	// We have no files
	return nil
}


func (a *StubArtifact) Id() string {
	return "amazon-ebs.artifact.stub"
}

func (a *StubArtifact) String() string {
	return "Stub artifact"
}

func (a *StubArtifact) State(name string) interface{} {
	return nil
}

func (a *StubArtifact) Destroy() error {
	return nil
}
